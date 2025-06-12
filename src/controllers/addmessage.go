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

// Structure avec injection de service et template
type AddMessageController struct {
	service  *services.MessagesServices
	template *template.Template
	store    *sessions.CookieStore
}

type AddMessageData struct {
	Topic_id      string
	Authenticated bool
	User_id       int
}

// Fonction pour initialiser le controller et les injections
func AddMessageControllerInit(template *template.Template, service *services.MessagesServices, store *sessions.CookieStore) *AddMessageController {
	return &AddMessageController{template: template, service: service, store: store}
}

// Routeur pour mettre en place les routes
func (c *AddMessageController) AddMessageRouter(r *mux.Router) {
	r.Handle("/addmessage", middlewares.RequireAuth(c.store, http.HandlerFunc(c.DisplayAddMessage))).Methods("GET")
	r.Handle("/addmessage/traitement", middlewares.RequireAuth(c.store, http.HandlerFunc(c.AddMessageTraitement))).Methods("POST")
}

// Fonction permettant d'afficher le formulaire
func (c *AddMessageController) DisplayAddMessage(w http.ResponseWriter, r *http.Request) {
	var data AddMessageData
	session, _ := c.store.Get(r, "session")

	data.Authenticated = true
	data.User_id = session.Values["user_id"].(int)
	data.Topic_id = r.FormValue("topic_id")

	c.template.ExecuteTemplate(w, "addmessage", data)
}

// Fonction de traitement pour le message
func (c *AddMessageController) AddMessageTraitement(w http.ResponseWriter, r *http.Request) {
	content := r.FormValue("content")
	topicIdString := r.FormValue("topic_id")

	topicId, errConv := strconv.Atoi(topicIdString)
	if errConv != nil {
		http.Redirect(w, r, "/error?code=400&message=invalid_topic_id", http.StatusSeeOther)
		return
	}

	session, _ := c.store.Get(r, "session")
	userId := session.Values["user_id"].(int)

	message := models.Messages{
		Topic_id: topicId,
		User_id:  userId,
		Content:  content,
	}

	_, err := c.service.CreateMessage(message)
	if err != nil {
		http.Redirect(w, r, "/error?code=500&message=create_message_failed", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/topic?id="+topicIdString, http.StatusMovedPermanently)
}
