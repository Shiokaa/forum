package controllers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

// Structure avec injection de service et template
type ConnexionControllers struct {
	template *template.Template
}

// Fonction pour initialiser le controller et les injections
func ConnexionControllerInit(template *template.Template) *ConnexionControllers {
	return &ConnexionControllers{template: template}
}

// Routeur pour mettre en place les routes de connexion
func (c *ConnexionControllers) ConnexionRouter(r *mux.Router) {
	r.HandleFunc("/connexion", c.DisplayConnexion).Methods("GET")
}

// Fonction permettant d'afficher la page formulaire d'Connexion avec une gestion d'erreur
func (c *ConnexionControllers) DisplayConnexion(w http.ResponseWriter, r *http.Request) {
	c.template.ExecuteTemplate(w, "connexion", nil)
}
