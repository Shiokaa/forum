package middlewares

import (
	"net/http"

	"github.com/gorilla/sessions"
)

// RequireAuth empêche l'accès aux utilisateurs non connectés
func RequireAuth(store *sessions.CookieStore, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")
		auth, ok := session.Values["authenticated"].(bool)
		if !ok || !auth {
			http.Redirect(w, r, "/connexion", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// RequireGuest empêche l'accès aux utilisateurs connectés
func RequireGuest(store *sessions.CookieStore, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")
		auth, ok := session.Values["authenticated"].(bool)
		if ok && auth {
			http.Redirect(w, r, "/accueil", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Vérifie si l'utilisateur possède une session
func SessionCheck(r *http.Request, store *sessions.CookieStore) bool {
	session, _ := store.Get(r, "session")
	auth, ok := session.Values["authenticated"].(bool)
	return ok && auth
}

// RequireAdmin empêche l'accès aux utilisateurs qui ne sont pas administrateurs.
func RequireAdmin(store *sessions.CookieStore, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")

		// Vérifier si l'utilisateur est authentifié
		auth, ok := session.Values["authenticated"].(bool)
		if !ok || !auth {
			http.Redirect(w, r, "/connexion", http.StatusSeeOther)
			return
		}

		// Vérifier si l'utilisateur a le rôle d'admin (ID = 1)
		roleID, ok := session.Values["role_id"].(int)
		if !ok || roleID != 1 {
			// Redirige vers l'accueil si pas admin
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		// Si tout est bon, on continue vers la page demandée
		next.ServeHTTP(w, r)
	})
}
