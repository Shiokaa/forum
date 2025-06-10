package controllers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type ErrorController struct {
	template *template.Template
}

type ErrorData struct {
	Code    string
	Message string
}

func ErrorControllerInit(template *template.Template) *ErrorController {
	return &ErrorController{template: template}
}

func (c *ErrorController) ErrorRouter(r *mux.Router) {
	r.HandleFunc("/error", c.DisplayError).Methods("GET")
}

func (c *ErrorController) DisplayError(w http.ResponseWriter, r *http.Request) {
	// Récupération des variables
	code := r.FormValue("code")
	message := r.FormValue("message")
	data := ErrorData{
		Code:    code,
		Message: message,
	}

	c.template.ExecuteTemplate(w, "error", data)
}
