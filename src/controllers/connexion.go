package controllers

import (
	"forum/src/middlewares"
	"forum/src/services"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

// Structure avec injection de service et template
type ConnexionControllers struct {
	service  *services.UsersServices
	template *template.Template
	store    *sessions.CookieStore
}

// Fonction pour initialiser le controller et les injections
func ConnexionControllerInit(template *template.Template, service *services.UsersServices, store *sessions.CookieStore) *ConnexionControllers {
	return &ConnexionControllers{template: template, service: service, store: store}
}

// Routeur pour mettre en place les routes de connexion
func (c *ConnexionControllers) ConnexionRouter(r *mux.Router) {
	r.Handle("/connexion", middlewares.RequireGuest(c.store, http.HandlerFunc(c.DisplayConnexion))).Methods("GET")
	r.Handle("/connexion/traitement", middlewares.RequireGuest(c.store, http.HandlerFunc(c.ConnexionTraitement))).Methods("POST")
	r.Handle("/deconnexion", middlewares.RequireAuth(c.store, http.HandlerFunc(c.DisconnectUser))).Methods("GET")

}

// Fonction permettant d'afficher la page formulaire d'Connexion avec une gestion d'erreur
func (c *ConnexionControllers) DisplayConnexion(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")

	if code != "" {
		c.template.ExecuteTemplate(w, "inscription", code)
		return
	}

	c.template.ExecuteTemplate(w, "connexion", nil)
}

func (c *ConnexionControllers) ConnexionTraitement(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "" || password == "" {
		http.Redirect(w, r, "/connexion?code=invalid_data", http.StatusSeeOther)
		return
	}

	session, _ := c.store.Get(r, "session")

	_, err := c.service.Connect(email, password)
	if err != nil {
		http.Redirect(w, r, "/connexion?code=invalid_data", http.StatusSeeOther)
		return
	}

	session.Values["authenticated"] = true
	session.Save(r, w)

	http.Redirect(w, r, "/accueil", http.StatusMovedPermanently)
}

func (c *ConnexionControllers) DisconnectUser(w http.ResponseWriter, r *http.Request) {
	session, _ := c.store.Get(r, "session")

	session.Values["authenticated"] = false
	session.Options.MaxAge = -1
	session.Save(r, w)

	http.Redirect(w, r, "/connexion", http.StatusSeeOther)
}
