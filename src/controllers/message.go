package controllers

import (
	"forum/src/models"
	"forum/src/services"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Structure avec injection de service et template
type MessageController struct {
	service  *services.MessagesServices
	template *template.Template
}

type MessageData struct {
	Item  models.Topics_Join_Messages
	Error bool
}

// Fonction pour initialiser le controller et les injections
func MessageControllerInit(template *template.Template, service *services.MessagesServices) *MessageController {
	return &MessageController{template: template, service: service}
}

// Routeur pour mettre en place les routes de message
func (c *MessageController) MessageRouter(r *mux.Router) {
	r.HandleFunc("/message", c.DisplayMessage).Methods("GET")
}

// Fonction permettant d'afficher la page formulaire d'inscription avec une gestion d'erreur
func (c *MessageController) DisplayMessage(w http.ResponseWriter, r *http.Request) {
	var data MessageData

	// Gérer les codes d'erreur passés en paramètre
	code := r.FormValue("code")
	if code == "invalid_id" || code == "item_not_found" {
		data.Error = true
		data.Item = models.Topics_Join_Messages{}
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

	item, err := c.service.MessageRepositories.GetMessageById(idInt)
	if err != nil {
		http.Redirect(w, r, "/topic?code=item_not_found", http.StatusSeeOther)
		return
	}

	data.Item = item

	c.template.ExecuteTemplate(w, "message", data)
}
