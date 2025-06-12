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
type TopicController struct {
	service  *services.TopicsServices
	template *template.Template
	store    *sessions.CookieStore
}

// Fonction pour initialiser le controller et les injections
func TopicControllerInit(template *template.Template, service *services.TopicsServices, store *sessions.CookieStore) *TopicController {
	return &TopicController{service: service, template: template, store: store}
}

// Structure créant un item qui est le topic ainsi que l'utilisateur ayant écrit le topic, récupère aussi les messages du topic
type TopicData struct {
	Item               models.Topics_Join_Users_Forums
	Messages           []models.Topics_Join_Messages
	Error              bool
	Authenticated      bool
	CreatedAtFormatted string
	UpdatedAtFormatted string
	Breadcrumbs        []models.Breadcrumb
}

// Routeur pour mettre en place les routes d'inscription
func (c *TopicController) TopicRouteur(r *mux.Router) {
	r.HandleFunc("/topic", c.DisplayTopic).Methods("GET")
}

// Fonction permettant d'afficher les topics et de gérer les données
func (c *TopicController) DisplayTopic(w http.ResponseWriter, r *http.Request) {
	var data TopicData

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

	// Lecture du topic par ID
	item, errReadId := c.service.ReadId(idInt)
	if errReadId != nil {
		http.Redirect(w, r, "/error?code=404&message=item_not_found", http.StatusSeeOther)
		return
	}

	// Lecture des messages par ID
	messages, errReadMessages := c.service.ReadMessages(idInt)
	if errReadMessages != nil {
		http.Redirect(w, r, "/error?code=404&message=messages_not_found", http.StatusSeeOther)
		return
	}

	for i := range messages {
		formatted, _ := utilitaire.ConvertTime(messages[i].Messages.Created_at, messages[i].Messages.Updated_at, w, r)
		messages[i].CreatedAtFormatted = formatted
	}

	data.Breadcrumbs = []models.Breadcrumb{
		{Name: "Accueil", URL: "/"},
		{Name: item.Topics.Title, URL: ""},
	}

	// Affichage du topic
	data.Item = item
	data.Messages = messages

	// Conversion de la date et affichage
	created_at, updated_at := utilitaire.ConvertTime(item.Topics.Created_at, item.Topics.Updated_at, w, r)
	data.CreatedAtFormatted = created_at
	data.UpdatedAtFormatted = updated_at

	c.template.ExecuteTemplate(w, "topic", data)
}
