package controllers

import (
	"fmt"
	"forum/src/middlewares"
	"forum/src/models"
	"forum/src/services"
	"forum/src/utilitaire"
	"html/template"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type AccueilController struct {
	template          *template.Template
	store             *sessions.CookieStore
	topicsService     *services.TopicsServices
	categoriesService *services.CategoriesServices
	messagesService   *services.MessagesServices
}

type AccueilData struct {
	Authenticated  bool
	User           models.Users
	Topics         []models.Topics_Join_Users
	Categories     []models.Categories
	RecentMessages []models.Topics_Join_Messages
	Pagination     PaginationData
	Breadcrumbs    []models.Breadcrumb
}

type PaginationData struct {
	CurrentPage int
	TotalPages  int
	HasPrev     bool
	HasNext     bool
	PrevPage    int
	NextPage    int
}

func AccueilControllerInit(template *template.Template, topicsService *services.TopicsServices, categoriesService *services.CategoriesServices, messagesService *services.MessagesServices, store *sessions.CookieStore) *AccueilController {
	return &AccueilController{
		template:          template,
		store:             store,
		topicsService:     topicsService,
		categoriesService: categoriesService,
		messagesService:   messagesService,
	}
}

func (c *AccueilController) AccueilRouter(r *mux.Router) {
	r.HandleFunc("/", c.DisplayAccueil).Methods("GET")
}

func (c *AccueilController) DisplayAccueil(w http.ResponseWriter, r *http.Request) {
	var data AccueilData

	data.Authenticated = middlewares.SessionCheck(r, c.store)
	session, _ := c.store.Get(r, "session")
	if data.Authenticated {
		data.User.User_id = session.Values["user_id"].(int)
	}
	data.Breadcrumbs = []models.Breadcrumb{
		{Name: "Accueil", URL: ""},
	}

	// Récupération des catégories
	categories, err := c.categoriesService.GetAll()
	if err != nil {
		http.Redirect(w, r, "/error?code=500&message=categories_not_found", http.StatusSeeOther)
		return
	}
	data.Categories = categories

	// Récupération des messages récents
	recentMessages, err := c.messagesService.DisplayRecent(5)
	if err != nil {
		fmt.Println(" Impossible d'afficher les messages")
	}
	data.RecentMessages = recentMessages

	// Gestion de la pagination pour les topics
	pageStr := r.FormValue("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	topicsPerPage := 10
	totalTopics, err := c.topicsService.GetTotalTopicsCount()
	if err != nil {
		http.Redirect(w, r, "/error?code=500&message=topics_count_failed", http.StatusSeeOther)
		return
	}

	totalPages := int(math.Ceil(float64(totalTopics) / float64(topicsPerPage)))
	offset := (page - 1) * topicsPerPage

	paginatedTopics, err := c.topicsService.GetPaginatedTopics(offset, topicsPerPage)
	if err != nil {
		http.Redirect(w, r, "/error?code=500&message=topics_fetch_failed", http.StatusSeeOther)
		return
	}

	// Formatter la date pour chaque topic
	for i := range paginatedTopics {
		formatted, _ := utilitaire.ConvertTime(paginatedTopics[i].Topics.Created_at, paginatedTopics[i].Topics.Updated_at, w, r)
		paginatedTopics[i].CreatedAtFormatted = formatted
	}
	data.Topics = paginatedTopics

	// Construire les données de pagination
	data.Pagination = PaginationData{
		CurrentPage: page,
		TotalPages:  totalPages,
		HasPrev:     page > 1,
		PrevPage:    page - 1,
		HasNext:     page < totalPages,
		NextPage:    page + 1,
	}

	c.template.ExecuteTemplate(w, "accueil", data)
}
