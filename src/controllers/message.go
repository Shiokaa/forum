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

// Structure avec injection de service et template
type MessageController struct {
	service  *services.MessagesServices
	template *template.Template
	store    *sessions.CookieStore
}

type MessageData struct {
	Item               models.Topics_Join_Messages
	Replies            []models.Replies_Join_User
	Error              bool
	Authenticated      bool
	CreatedAtFormatted string
	Breadcrumbs        []models.Breadcrumb
}

// Fonction pour initialiser le controller et les injections
func MessageControllerInit(template *template.Template, service *services.MessagesServices, store *sessions.CookieStore) *MessageController {
	return &MessageController{template: template, service: service, store: store}
}

// Routeur pour mettre en place les routes de message
func (c *MessageController) MessageRouter(r *mux.Router) {
	r.HandleFunc("/message", c.DisplayMessage).Methods("GET")
}

// Fonction permettant d'afficher la page formulaire d'inscription avec une gestion d'erreur
func (c *MessageController) DisplayMessage(w http.ResponseWriter, r *http.Request) {
	var data MessageData

	// Determine si l'utilisateur est connecté ou non
	data.Authenticated = middlewares.SessionCheck(r, c.store)

	session, _ := c.store.Get(r, "session")

	if data.Authenticated {
		data.Item.Users.User_id = session.Values["user_id"].(int)
	}
	// Récupération de l'ID depuis les paramètres
	idString := r.FormValue("id")
	idInt, errConv := strconv.Atoi(idString)
	if errConv != nil {
		http.Redirect(w, r, "/error?code=404&message=invalid_id", http.StatusSeeOther)
		return
	}

	item, errMessages := c.service.ReadMessagesId(idInt)
	if errMessages != nil {
		http.Redirect(w, r, "/error?code=404&message=item_not_found", http.StatusSeeOther)
		return
	}

	data.Breadcrumbs = []models.Breadcrumb{
		{Name: "Accueil", URL: "/"},
		{Name: item.Topics.Title, URL: "/topic?id=" + strconv.Itoa(item.Topics.Topic_id)},
		{Name: "Message", URL: ""},
	}

	items, errReplies := c.service.ReadRepliesId(idInt)
	if errReplies != nil {
		http.Redirect(w, r, "/error?code=404&message=item_not_found", http.StatusSeeOther)
		return
	}

	created_at, _ := utilitaire.ConvertTime(item.Messages.Created_at, item.Messages.Updated_at, w, r)
	data.CreatedAtFormatted = created_at

	for i := range items {
		formatted, _ := utilitaire.ConvertTime(items[i].Replies.Created_at, items[i].Replies.Updated_at, w, r)
		items[i].CreatedAtFormatted = formatted
	}

	data.Replies = items
	data.Item = item

	c.template.ExecuteTemplate(w, "message", data)
}
