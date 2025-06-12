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
type RepliesController struct {
	service  *services.MessagesServices
	template *template.Template
	store    *sessions.CookieStore
}

type RepliesData struct {
	Item          models.Replies_Joins_User_Message
	Message_id    string
	Authenticated bool
	Breadcrumbs   []models.Breadcrumb
}

// Fonction pour initialiser le controller et les injections
func RepliesControllerInit(template *template.Template, service *services.MessagesServices, store *sessions.CookieStore) *RepliesController {
	return &RepliesController{template: template, service: service, store: store}
}

// Routeur pour mettre en place les routes de reponse
func (c *RepliesController) RepliesRouter(r *mux.Router) {
	r.Handle("/reponse", middlewares.RequireAuth(c.store, http.HandlerFunc(c.DisplayReplies))).Methods("GET")
	r.Handle("/reponse/traitement", middlewares.RequireAuth(c.store, http.HandlerFunc(c.ReplyTraitement))).Methods("POST")
}

// Fonction permettant d'afficher la page formulaire d'inscription avec une gestion d'erreur
func (c *RepliesController) DisplayReplies(w http.ResponseWriter, r *http.Request) {
	var data RepliesData

	session, _ := c.store.Get(r, "session")

	data.Authenticated = true
	data.Item.Users.User_id = session.Values["user_id"].(int)

	message_id := r.FormValue("id")
	messageIdInt, _ := strconv.Atoi(message_id)

	messageDetails, err := c.service.ReadMessagesId(messageIdInt)
	if err != nil {
		http.Redirect(w, r, "/error?code=404&message=item_not_found", http.StatusSeeOther)
		return
	}

	data.Breadcrumbs = []models.Breadcrumb{
		{Name: "Accueil", URL: "/"},
		{Name: messageDetails.Topics.Title, URL: "/topic?id=" + strconv.Itoa(messageDetails.Topics.Topic_id)},
		{Name: "Message", URL: "/message?id=" + message_id},
		{Name: "Répondre", URL: ""},
	}

	data.Message_id = message_id

	c.template.ExecuteTemplate(w, "reponse", data)
}

func (c *RepliesController) ReplyTraitement(w http.ResponseWriter, r *http.Request) {
	content := r.FormValue("content")

	idString := r.FormValue("id")
	idInt, errConv := strconv.Atoi(idString)
	if errConv != nil {
		http.Redirect(w, r, "/erreur?code=invalid_id", http.StatusSeeOther)
		return
	}

	session, _ := c.store.Get(r, "session")

	user := models.Users{
		User_id: session.Values["user_id"].(int),
	}

	reply := models.Replies{
		Content: content,
	}

	message := models.Messages{
		Message_id: idInt,
	}

	item := models.Replies_Joins_User_Message{
		Users:    user,
		Replies:  reply,
		Messages: message,
	}

	_, err := c.service.CreatedReply(item)
	if err != nil {
		http.Redirect(w, r, "/erreur?code=invalid_data", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/message?id="+idString, http.StatusMovedPermanently)
}
