package controllers

import (
	"forum/src/middlewares"
	"forum/src/models"
	"forum/src/services"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type CreateTopicController struct {
	template        *template.Template
	store           *sessions.CookieStore
	topicsService   *services.TopicsServices
	forumsService   *services.ForumsServices
	messagesService *services.MessagesServices
}

type CreateTopicPageData struct {
	Authenticated bool
	User          models.Users
	Forums        []models.Forums
	SelectedForum int
	Breadcrumbs   []models.Breadcrumb
}

func CreateTopicControllerInit(t *template.Template, s *sessions.CookieStore, ts *services.TopicsServices, fs *services.ForumsServices, ms *services.MessagesServices) *CreateTopicController {
	return &CreateTopicController{
		template:        t,
		store:           s,
		topicsService:   ts,
		forumsService:   fs,
		messagesService: ms,
	}
}

func (c *CreateTopicController) CreateTopicRouter(r *mux.Router) {
	r.Handle("/topic/creer", middlewares.RequireAuth(c.store, http.HandlerFunc(c.DisplayCreateTopicForm))).Methods("GET")
	r.Handle("/topic/creer/traitement", middlewares.RequireAuth(c.store, http.HandlerFunc(c.ProcessCreateTopicForm))).Methods("POST")
}

// DisplayCreateTopicForm affiche le formulaire de création.
func (c *CreateTopicController) DisplayCreateTopicForm(w http.ResponseWriter, r *http.Request) {
	var data CreateTopicPageData
	session, _ := c.store.Get(r, "session")
	data.Authenticated = true
	data.User.User_id = session.Values["user_id"].(int)

	// Récupérer tous les forums pour le menu déroulant
	forums, _ := c.forumsService.GetAll()
	data.Forums = forums

	// Pré-sélectionner le forum si un ID est passé dans l'URL
	forumIDStr := r.FormValue("forum_id")
	if forumID, err := strconv.Atoi(forumIDStr); err == nil {
		data.SelectedForum = forumID
	}

	data.Breadcrumbs = []models.Breadcrumb{
		{Name: "Accueil", URL: "/"},
		{Name: "Créer un Topic", URL: ""},
	}
	c.template.ExecuteTemplate(w, "create_topic", data)
}

// ProcessCreateTopicForm traite la soumission du formulaire.
func (c *CreateTopicController) ProcessCreateTopicForm(w http.ResponseWriter, r *http.Request) {
	session, _ := c.store.Get(r, "session")
	userID := session.Values["user_id"].(int)

	r.ParseForm()
	title := r.FormValue("title")
	content := r.FormValue("content")
	forumID, _ := strconv.Atoi(r.FormValue("forum_id"))

	newTopic := models.Topics{
		Title:    title,
		Forum_id: forumID,
		User_Id:  userID,
	}
	topicID, err := c.topicsService.CreateTopic(newTopic)
	if err != nil {
		http.Redirect(w, r, "/error?code=500&message=topic_creation_failed", http.StatusSeeOther)
		return
	}

	initialMessage := models.Messages{
		Topic_id: topicID,
		User_id:  userID,
		Content:  content,
	}
	_, err = c.messagesService.CreateMessage(initialMessage)
	if err != nil {
		http.Redirect(w, r, "/error?code=500&message=message_creation_failed", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/topic?id="+strconv.Itoa(topicID), http.StatusSeeOther)
}
