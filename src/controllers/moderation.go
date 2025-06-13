package controllers

import (
	"forum/src/middlewares"
	"forum/src/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type ModerationController struct {
	store           *sessions.CookieStore
	messagesService *services.MessagesServices
}

func ModerationControllerInit(store *sessions.CookieStore, msgService *services.MessagesServices) *ModerationController {
	return &ModerationController{
		store:           store,
		messagesService: msgService,
	}
}

func (c *ModerationController) ModerationRouter(r *mux.Router) {
	r.Handle("/message/delete", middlewares.RequireAuth(c.store, http.HandlerFunc(c.DeleteMessage))).Methods("POST")
}

func (c *ModerationController) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	// Récupérer l'ID du message depuis le formulaire
	messageIDStr := r.FormValue("message_id")
	messageID, _ := strconv.Atoi(messageIDStr)

	// Récupérer les informations de l'utilisateur depuis la session
	session, _ := c.store.Get(r, "session")
	currentUserID := session.Values["user_id"].(int)
	currentUserRoleID := session.Values["role_id"].(int)

	// Récupérer les détails du message pour vérifier le propriétaire et obtenir le topic_id
	message, err := c.messagesService.ReadMessagesId(messageID)
	if err != nil {
		http.Redirect(w, r, "/error?code=404&message=message_not_found", http.StatusSeeOther)
		return
	}

	// VÉRIFICATION DES DROITS : l'utilisateur est-il le propriétaire OU un admin (rôle ID 1) ?
	if currentUserID != message.Messages.User_id && currentUserRoleID != 1 {
		http.Redirect(w, r, "/error?code=403&message=forbidden", http.StatusSeeOther)
		return
	}

	// Suppression du message
	err = c.messagesService.DeleteMessage(messageID)
	if err != nil {
		http.Redirect(w, r, "/error?code=500&message=delete_failed", http.StatusSeeOther)
		return
	}

	// Redirection vers la page du topic d'origine
	http.Redirect(w, r, "/topic?id="+strconv.Itoa(message.Messages.Topic_id), http.StatusSeeOther)
}
