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
type ProfilControllers struct {
	service  *services.UsersServices
	template *template.Template
	store    *sessions.CookieStore
}

type ProfilData struct {
	User               models.Users
	Authenticated      bool
	CreatedAtFormatted string
	UpdatedAtFormatted string
	Breadcrumbs        []models.Breadcrumb
}

// Fonction pour initialiser le controller et les injections
func ProfilControllerInit(template *template.Template, service *services.UsersServices, store *sessions.CookieStore) *ProfilControllers {
	return &ProfilControllers{template: template, service: service, store: store}
}

// Routeur pour mettre en place les routes d'accueil
func (c *ProfilControllers) ProfilRouter(r *mux.Router) {
	r.HandleFunc("/profil", c.DisplayProfil).Methods("GET")
}

func (c *ProfilControllers) DisplayProfil(w http.ResponseWriter, r *http.Request) {
	var data ProfilData

	data.Authenticated = middlewares.SessionCheck(r, c.store)

	idString := r.FormValue("id")
	idInt, errConv := strconv.Atoi(idString)
	if errConv != nil {
		http.Redirect(w, r, "/erreur?code=invalid_id", http.StatusSeeOther)
		return
	}

	user, _ := c.service.ReadId(idInt)
	data.User = user

	data.Breadcrumbs = []models.Breadcrumb{
		{Name: "Accueil", URL: "/"},
		{Name: "Profil de " + user.Name, URL: ""},
	}

	created_at, updated_at := utilitaire.ConvertTime(data.User.Created_at, data.User.Updated_at, w, r)

	data.CreatedAtFormatted = created_at
	data.UpdatedAtFormatted = updated_at

	c.template.ExecuteTemplate(w, "profil", data)
}
