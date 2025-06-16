package controllers

import (
	"errors"
	"fmt"
	"forum/src/middlewares"
	"forum/src/models"
	"forum/src/services"
	"html/template"
	"net/http"
	"net/url"
	"unicode"

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
	Erreur  string
	Message string
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
	message := r.FormValue("message")

	if code != "" {
		data.Erreur = code
		data.Message = message
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

	errValidation := validatePassword(password)
	if errValidation != nil {
		// On redirige avec un message d'erreur spécifique pour le mot de passe
		errorMessage := url.QueryEscape(errValidation.Error())
		redirectURL := fmt.Sprintf("/inscription?code=invalid_password&message=%s", errorMessage)
		http.Redirect(w, r, redirectURL, http.StatusSeeOther)
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

func validatePassword(password string) error {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	if len(password) >= 12 {
		hasMinLen = true
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char): // C'est une bonne pratique d'exiger aussi une minuscule
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasMinLen {
		return errors.New(" Le mot de passe doit faire au moins 12 caractères")
	}
	if !hasUpper {
		return errors.New(" Le mot de passe doit contenir au moins une majuscule")
	}
	if !hasLower {
		return errors.New(" Le mot de passe doit contenir au moins une minuscule")
	}
	if !hasNumber {
		return errors.New(" Le mot de passe doit contenir au moins un chiffre")
	}
	if !hasSpecial {
		return errors.New(" Le mot de passe doit contenir au moins un caractère spécial")
	}

	return nil // Le mot de passe est valide
}
