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
