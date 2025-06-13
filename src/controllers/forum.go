package controllers

import (
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

type ForumController struct {
	template      *template.Template
	store         *sessions.CookieStore
	forumsService *services.ForumsServices
	topicsService *services.TopicsServices
}

type ForumPageData struct {
	Authenticated bool
	User          models.Users
	ForumInfo     models.ForumWithCategory
	Topics        []models.Topics_Join_Users
	Pagination    PaginationData
	Breadcrumbs   []models.Breadcrumb
}

func ForumControllerInit(template *template.Template, store *sessions.CookieStore, forumService *services.ForumsServices, topicService *services.TopicsServices) *ForumController {
	return &ForumController{
		template:      template,
		store:         store,
		forumsService: forumService,
		topicsService: topicService,
	}
}

func (c *ForumController) ForumRouter(r *mux.Router) {
	r.HandleFunc("/forum", c.DisplaySingleForum).Methods("GET")
}

func (c *ForumController) DisplaySingleForum(w http.ResponseWriter, r *http.Request) {
	var data ForumPageData
	data.Authenticated = middlewares.SessionCheck(r, c.store)
	if data.Authenticated {
		session, _ := c.store.Get(r, "session")
		data.User.User_id = session.Values["user_id"].(int)
	}

	// Récupérer l'ID du forum depuis l'URL
	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Redirect(w, r, "/error?code=400&message=invalid_id", http.StatusSeeOther)
		return
	}

	// Récupérer les informations du forum et de sa catégorie
	forumInfo, err := c.forumsService.GetByIDWithCategory(id)
	if err != nil {
		http.Redirect(w, r, "/error?code=404&message=forum_not_found", http.StatusSeeOther)
		return
	}
	data.ForumInfo = forumInfo

	// Pagination des topics
	pageStr := r.FormValue("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	topicsPerPage := 15
	totalTopics, _ := c.topicsService.GetTotalTopicsCountByForumID(id)
	totalPages := int(math.Ceil(float64(totalTopics) / float64(topicsPerPage)))
	offset := (page - 1) * topicsPerPage

	topics, _ := c.topicsService.GetByForumIDPaginated(id, offset, topicsPerPage)
	for i := range topics {
		formatted, _ := utilitaire.ConvertTime(topics[i].Topics.Created_at, topics[i].Topics.Updated_at, w, r)
		topics[i].CreatedAtFormatted = formatted
	}
	data.Topics = topics
	data.Pagination = PaginationData{
		CurrentPage: page, TotalPages: totalPages,
		HasPrev: page > 1, PrevPage: page - 1,
		HasNext: page < totalPages, NextPage: page + 1,
	}

	// Fil d'Ariane
	data.Breadcrumbs = []models.Breadcrumb{
		{Name: "Accueil", URL: "/"},
		{Name: "Catégories", URL: "/categories"},
		{Name: forumInfo.Category.Name, URL: "/categorie?id=" + strconv.Itoa(forumInfo.Category.Category_id)},
		{Name: forumInfo.Forum.Name, URL: ""},
	}

	c.template.ExecuteTemplate(w, "forum", data)
}
