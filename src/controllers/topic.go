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
	Item          models.Topics_Join_Users_Forums
	Messages      []models.Topics_Join_Messages
	Error         bool
	Authenticated bool
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

	// Gérer les codes d'erreur passés en paramètre
	code := r.FormValue("code")
	if code == "invalid_id" || code == "item_not_found" || code == "messages_not_found" {
		data.Error = true
		data.Item = models.Topics_Join_Users_Forums{}
		data.Messages = []models.Topics_Join_Messages{}
		c.template.ExecuteTemplate(w, "topic", data)
		return
	}

	// Récupération de l'ID depuis les paramètres
	idString := r.FormValue("id")
	idInt, errConv := strconv.Atoi(idString)
	if errConv != nil {
		http.Redirect(w, r, "/topic?code=invalid_id", http.StatusSeeOther)
		return
	}

	// Lecture du topic par ID
	item, errReadId := c.service.ReadId(idInt)
	if errReadId != nil {
		http.Redirect(w, r, "/topic?code=item_not_found", http.StatusSeeOther)
		return
	}

	// Lecture des messages par ID
	messages, errReadMessages := c.service.ReadMessages(idInt)
	if errReadMessages != nil {
		http.Redirect(w, r, "/topic?code=messages_not_found", http.StatusSeeOther)
		return
	}

	// Affichage du topic
	data.Item = item
	data.Messages = messages

	c.template.ExecuteTemplate(w, "topic", data)
}
