package controllers

import (
	"forum/src/middlewares"
	"forum/src/models"
	"forum/src/services"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type AdminController struct {
	template     *template.Template
	store        *sessions.CookieStore
	usersService *services.UsersServices
}

type AdminDashboardData struct {
	Authenticated bool
	User          models.Users   // L'admin actuellement connecté
	AllUsers      []models.Users // La liste de tous les utilisateurs
	Breadcrumbs   []models.Breadcrumb
}

func AdminControllerInit(t *template.Template, s *sessions.CookieStore, u *services.UsersServices) *AdminController {
	return &AdminController{
		template:     t,
		store:        s,
		usersService: u,
	}
}

func (c *AdminController) AdminRouter(r *mux.Router) {
	// On protège toute la section /admin avec notre nouveau middleware
	adminSubrouter := r.PathPrefix("/admin").Subrouter()
	adminSubrouter.Use(func(next http.Handler) http.Handler {
		return middlewares.RequireAdmin(c.store, next)
	})

	adminSubrouter.HandleFunc("", c.DisplayDashboard).Methods("GET")
	adminSubrouter.HandleFunc("/user/delete", c.DeleteUser).Methods("POST")
}

// DisplayDashboard affiche la page principale du dashboard.
func (c *AdminController) DisplayDashboard(w http.ResponseWriter, r *http.Request) {
	var data AdminDashboardData
	session, _ := c.store.Get(r, "session")

	data.Authenticated = true // Garanti par le middleware
	data.User.User_id = session.Values["user_id"].(int)
	data.User.Role_id = session.Values["role_id"].(int)

	// Récupérer la liste de tous les utilisateurs
	allUsers, _ := c.usersService.GetAll()
	data.AllUsers = allUsers

	data.Breadcrumbs = []models.Breadcrumb{
		{Name: "Accueil", URL: "/"},
		{Name: "Dashboard Admin", URL: ""},
	}

	c.template.ExecuteTemplate(w, "admin_dashboard", data)
}

// DeleteUser gère la suppression d'un compte par un admin.
func (c *AdminController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	session, _ := c.store.Get(r, "session")
	adminUserID := session.Values["user_id"].(int)

	userIDToDeleteStr := r.FormValue("user_id")
	userIDToDelete, _ := strconv.Atoi(userIDToDeleteStr)

	// VÉRIFICATION CRITIQUE : Un admin ne peut pas supprimer son propre compte.
	if adminUserID == userIDToDelete {
		http.Redirect(w, r, "/error?code=403&message=self_delete_forbidden", http.StatusSeeOther)
		return
	}

	// Appel du service pour supprimer l'utilisateur
	err := c.usersService.DeleteUser(userIDToDelete)
	if err != nil {
		http.Redirect(w, r, "/error?code=500&message=user_delete_failed", http.StatusSeeOther)
		return
	}

	// Redirection vers le dashboard admin
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
