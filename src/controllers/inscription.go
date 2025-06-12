package controllers

import (
	"forum/src/middlewares"
	"forum/src/models"
	"forum/src/services"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

// Structure avec injection de service et template
type InscriptionController struct {
	service  *services.UsersServices
	template *template.Template
	store    *sessions.CookieStore
}

type InscriptionData struct {
	Erreur string
}

// Fonction pour initialiser le controller et les injections
func InscriptionControllerInit(template *template.Template, service *services.UsersServices, store *sessions.CookieStore) *InscriptionController {
	return &InscriptionController{template: template, service: service, store: store}
}

// Routeur pour mettre en place les routes d'inscription
func (c *InscriptionController) InsciptionRouter(r *mux.Router) {
	r.Handle("/inscription", middlewares.RequireGuest(c.store, http.HandlerFunc(c.DisplayInscription))).Methods("GET")
	r.Handle("/inscription/traitement", middlewares.RequireGuest(c.store, http.HandlerFunc(c.InscriptionTraitement))).Methods("POST")
}

// Fonction permettant d'afficher la page formulaire d'inscription avec une gestion d'erreur
func (c *InscriptionController) DisplayInscription(w http.ResponseWriter, r *http.Request) {
	var data InscriptionData

	code := r.FormValue("code")
	if code == "invalid_data" {
		data.Erreur = "invalid_data"
		c.template.ExecuteTemplate(w, "inscription", data)
		return
	} else if code == "data_exist" {
		data.Erreur = "data_exist"
		c.template.ExecuteTemplate(w, "inscription", data)
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
		http.Redirect(w, r, "/inscription?code=invalid_data", http.StatusSeeOther)
		return
	}

	hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(password), 14)
	if errHash != nil {
		http.Redirect(w, r, "/inscription?code=invalid_data", http.StatusSeeOther)
		return
	}

	// Création d'un objet user pour enregistrer les données reçu
	newUser := models.Users{
		Role_id:  3,
		Name:     username,
		Email:    email,
		Password: string(hashedPassword),
	}

	// Création de l'utilisateur via le service de création d'utilisateur
	_, userErr := c.service.Create(newUser)
	if userErr != nil {
		http.Redirect(w, r, "/inscription?code=data_exist", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/connexion", http.StatusMovedPermanently)
}
