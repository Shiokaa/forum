package controllers

import (
	"forum/src/middlewares"
	"forum/src/models"
	"forum/src/services"
	"forum/src/utilitaire"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type CategoryController struct {
	template          *template.Template
	store             *sessions.CookieStore
	categoriesService *services.CategoriesServices
	forumsService     *services.ForumsServices
	topicsService     *services.TopicsServices
}

type CategoryPageData struct {
	Authenticated bool
	User          models.Users
	Categories    []models.CategoryWithForums
	Breadcrumbs   []models.Breadcrumb
}

type SingleCategoryData struct {
	Authenticated bool
	User          models.Users
	Category      models.Categories
	Forums        []models.Forums
	Topics        []models.Topics_Join_Users
	Breadcrumbs   []models.Breadcrumb
}

func CategoryControllerInit(template *template.Template, store *sessions.CookieStore, catService *services.CategoriesServices, forumService *services.ForumsServices, topicService *services.TopicsServices) *CategoryController {
	return &CategoryController{
		template:          template,
		store:             store,
		categoriesService: catService,
		forumsService:     forumService,
		topicsService:     topicService,
	}
}

func (c *CategoryController) CategoryRouter(r *mux.Router) {
	r.HandleFunc("/categories", c.DisplayCategories).Methods("GET")
	r.HandleFunc("/categorie", c.DisplaySingleCategory).Methods("GET")
}

func (c *CategoryController) DisplayCategories(w http.ResponseWriter, r *http.Request) {
	var data CategoryPageData
	data.Authenticated = middlewares.SessionCheck(r, c.store)
	if data.Authenticated {
		session, _ := c.store.Get(r, "session")
		data.User.User_id = session.Values["user_id"].(int)
		data.User.Role_id = session.Values["role_id"].(int)
	}

	categoriesWithForums, err := c.categoriesService.GetCategoriesWithForums(c.forumsService)
	if err != nil {
		http.Redirect(w, r, "/error?code=500&message=data_fetch_error", http.StatusSeeOther)
		return
	}
	data.Categories = categoriesWithForums

	data.Breadcrumbs = []models.Breadcrumb{
		{Name: "Accueil", URL: "/"},
		{Name: "Catégories", URL: ""},
	}

	c.template.ExecuteTemplate(w, "categories", data)
}

func (c *CategoryController) DisplaySingleCategory(w http.ResponseWriter, r *http.Request) {
	var data SingleCategoryData
	data.Authenticated = middlewares.SessionCheck(r, c.store)
	if data.Authenticated {
		session, _ := c.store.Get(r, "session")
		data.User.User_id = session.Values["user_id"].(int)
		data.User.Role_id = session.Values["role_id"].(int)
	}

	// Récupérer l'ID de la catégorie depuis l'URL
	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Redirect(w, r, "/error?code=400&message=invalid_id", http.StatusSeeOther)
		return
	}

	// Récupérer les données
	category, err := c.categoriesService.GetByID(id)
	if err != nil {
		http.Redirect(w, r, "/error?code=404&message=category_not_found", http.StatusSeeOther)
		return
	}
	data.Category = category

	forums, _ := c.forumsService.GetByCategoryID(id)
	data.Forums = forums

	topics, _ := c.topicsService.GetByCategoryID(id, 10) // On récupère les 10 derniers topics
	// Formatter la date pour chaque topic
	for i := range topics {
		formatted, _ := utilitaire.ConvertTime(topics[i].Topics.Created_at, topics[i].Topics.Updated_at, w, r)
		topics[i].CreatedAtFormatted = formatted
	}
	data.Topics = topics

	// Construire le fil d'Ariane
	data.Breadcrumbs = []models.Breadcrumb{
		{Name: "Accueil", URL: "/"},
		{Name: "Catégories", URL: "/categories"},
		{Name: category.Name, URL: ""},
	}

	c.template.ExecuteTemplate(w, "single_category", data)
}
