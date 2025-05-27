package controllers

import (
	"fmt"
	"forum/src/services"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type AccueilController struct {
	service   *services.TopicsServices
	templates *template.Template
}

type DisplayData struct {
	Username string
	Title    string
}

func AccueilControllerInit(template *template.Template, service *services.TopicsServices) *AccueilController {
	return &AccueilController{templates: template, service: service}
}

func (c *AccueilController) AccueilRouter(r *mux.Router) {
	r.HandleFunc("/accueil", c.DisplayAccueil).Methods("GET")
}

func (c *AccueilController) DisplayAccueil(w http.ResponseWriter, r *http.Request) {
	var data DisplayData

	topicTitle, username, err := c.service.Display()
	if err != nil {
		http.Redirect(w, r, "404", http.StatusMovedPermanently)
	}

	fmt.Println(topicTitle, username)

	c.templates.ExecuteTemplate(w, "accueil", data)
}
