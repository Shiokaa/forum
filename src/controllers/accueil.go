package controllers

import (
	"forum/src/middlewares"
	"forum/src/models"
	"forum/src/services"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

// Structure avec injection de service et template
type AccueilController struct {
	service  *services.TopicsServices
	template *template.Template
	store    *sessions.CookieStore
}

// Structure créant une liste de topics et users
type AccueilData struct {
	TopicsWithUsers []models.Topics_Join_Users
	Error           bool
	Authenticated   bool
	User            models.Users
}

// Fonction pour initialiser le controller et les injections
func AccueilControllerInit(template *template.Template, service *services.TopicsServices, store *sessions.CookieStore) *AccueilController {
	return &AccueilController{template: template, service: service, store: store}
}

// Routeur pour mettre en place les routes d'accueil
func (c *AccueilController) AccueilRouter(r *mux.Router) {
	r.HandleFunc("/accueil", c.DisplayAccueil).Methods("GET")
}

// Fonctiob permettant d'afficher les données sur l'accueil
func (c *AccueilController) DisplayAccueil(w http.ResponseWriter, r *http.Request) {
	// Récupération des variables
	var data AccueilData

	// Determine si l'utilisateur est connecté ou non
	data.Authenticated = middlewares.SessionCheck(r, c.store)

	session, _ := c.store.Get(r, "session")

	if data.Authenticated {
		data.User.User_id = session.Values["user_id"].(int)
	}

	// Vérification de la précense des données
	code := r.FormValue("code")
	if code != "" {
		data.Error = true
		data.TopicsWithUsers = []models.Topics_Join_Users{}
		c.template.ExecuteTemplate(w, "/accueil", data)
		return
	}

	// Ici on récupére les données de la base de donnée
	items, errData := c.service.Display()
	if errData != nil {
		http.Redirect(w, r, "/accueil?code=invalid_data", http.StatusSeeOther)
	}

	// Récupération des données pour les envoyés dans l'html
	data.TopicsWithUsers = items

	c.template.ExecuteTemplate(w, "accueil", data)
}
