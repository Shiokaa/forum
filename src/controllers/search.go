// Fichier : src/controllers/search.go
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

type SearchController struct {
	template      *template.Template
	store         *sessions.CookieStore
	searchService *services.SearchServices
}

type SearchPageData struct {
	Authenticated bool
	User          models.Users
	Query         string
	Results       []models.SearchResult
	Breadcrumbs   []models.Breadcrumb
}

func SearchControllerInit(template *template.Template, store *sessions.CookieStore, searchService *services.SearchServices) *SearchController {
	return &SearchController{
		template:      template,
		store:         store,
		searchService: searchService,
	}
}

func (c *SearchController) SearchRouter(r *mux.Router) {
	r.HandleFunc("/recherche", c.DisplaySearchResults).Methods("GET")
}

func (c *SearchController) DisplaySearchResults(w http.ResponseWriter, r *http.Request) {
	var data SearchPageData
	data.Authenticated = middlewares.SessionCheck(r, c.store)
	if data.Authenticated {
		session, _ := c.store.Get(r, "session")
		data.User.User_id = session.Values["user_id"].(int)
	}

	// Récupérer le terme de la recherche depuis l'URL (?q=...)
	query := r.FormValue("q")
	data.Query = query

	// Lancer la recherche si la requête n'est pas vide
	if query != "" {
		results, err := c.searchService.PerformGlobalSearch(query)
		if err == nil {
			data.Results = results
		}
	}

	data.Breadcrumbs = []models.Breadcrumb{
		{Name: "Accueil", URL: "/"},
		{Name: "Recherche", URL: ""},
	}

	c.template.ExecuteTemplate(w, "search_results", data)
}
