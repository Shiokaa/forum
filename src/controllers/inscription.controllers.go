package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type InscriptionController struct {
	template *template.Template
}

func InscriptionControllerInit(template *template.Template) *InscriptionController {
	return &InscriptionController{template}
}

func (c *InscriptionController) InsciptionRouter(r *mux.Router) {
	r.HandleFunc("/inscription", c.DisplayInscription).Methods("GET")
	r.HandleFunc("/inscription/traitement", c.InscriptionTraitement).Methods("POST")
}

func (c *InscriptionController) DisplayInscription(w http.ResponseWriter, r *http.Request) {
	err := r.FormValue("code")
	if err != "" {
		c.template.ExecuteTemplate(w, "inscription", err)
		return
	}

	c.template.ExecuteTemplate(w, "inscription", nil)
}

func (c *InscriptionController) InscriptionTraitement(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	if username == "" || email == "" || password == "" {
		http.Redirect(w, r, "/inscription?code=400", http.StatusMovedPermanently)
	}

	/* 	newUser := models.User{
		Name:     username,
		Email:    email,
		Password: password,
	} */

	fmt.Printf("Username : %s   Email : %s    Password : %s", username, email, password)
}
