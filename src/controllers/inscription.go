package controllers

import (
	"forum/src/models"
	"forum/src/services"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

// Structure avec injection de service et template
type InscriptionController struct {
	service  *services.UsersServices
	template *template.Template
}

// Fonction pour initialiser le controller et les injections
func InscriptionControllerInit(template *template.Template, service *services.UsersServices) *InscriptionController {
	return &InscriptionController{template: template, service: service}
}

// Routeur pour mettre en place les routes d'inscription
func (c *InscriptionController) InsciptionRouter(r *mux.Router) {
	r.HandleFunc("/inscription", c.DisplayInscription).Methods("GET")
	r.HandleFunc("/inscription/traitement", c.InscriptionTraitement).Methods("POST")
}

// Fonction permettant d'afficher la page formulaire d'inscription avec une gestion d'erreur
func (c *InscriptionController) DisplayInscription(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	if code != "" {
		c.template.ExecuteTemplate(w, "inscription", code)
		return
	}

	c.template.ExecuteTemplate(w, "inscription", nil)
}

// Fonction de traitement pour gérer les données envoyées par l'utilisateur dans la page d'inscription
func (c *InscriptionController) InscriptionTraitement(w http.ResponseWriter, r *http.Request) {
	// Récupération des données
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Vérification de la présence des données
	if username == "" || email == "" || password == "" {
		http.Redirect(w, r, "/inscription?invalid_data", http.StatusSeeOther)
		return
	}

	// Création d'un objet user pour enregistrer les données reçu
	newUser := models.Users{
		Role_id:  3,
		Name:     username,
		Email:    email,
		Password: password,
	}

	// Création de l'utilisateur via le service de création d'utilisateur
	_, userErr := c.service.Create(newUser)
	if userErr != nil {
		http.Error(w, userErr.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/accueil", http.StatusMovedPermanently)
}
