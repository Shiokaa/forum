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
	topicsService   *services.TopicsServices
}

func ModerationControllerInit(store *sessions.CookieStore, msgService *services.MessagesServices, topicService *services.TopicsServices) *ModerationController {
	return &ModerationController{
		store:           store,
		messagesService: msgService,
		topicsService:   topicService,
	}
}

func (c *ModerationController) ModerationRouter(r *mux.Router) {
	r.Handle("/message/delete", middlewares.RequireAuth(c.store, http.HandlerFunc(c.DeleteMessage))).Methods("POST")
	r.Handle("/reply/delete", middlewares.RequireAuth(c.store, http.HandlerFunc(c.DeleteReply))).Methods("POST")
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

// DeleteReply gère la suppression d'une réponse à un message.
func (c *ModerationController) DeleteReply(w http.ResponseWriter, r *http.Request) {
	replyIDStr := r.FormValue("reply_id")
	replyID, _ := strconv.Atoi(replyIDStr)

	session, _ := c.store.Get(r, "session")
	currentUserID, okUserID := session.Values["user_id"].(int)
	currentUserRoleID, okRoleID := session.Values["role_id"].(int)

	if !okUserID || !okRoleID {
		http.Redirect(w, r, "/error?code=401&message=invalid_session", http.StatusSeeOther)
		return
	}

	// Récupérer les détails de la réponse pour vérifier les droits et la redirection
	reply, err := c.messagesService.ReadReplyByID(replyID)
	if err != nil {
		http.Redirect(w, r, "/error?code=404&message=reply_not_found", http.StatusSeeOther)
		return
	}

	// VÉRIFICATION DES DROITS : l'utilisateur est-il le propriétaire OU un admin ?
	if currentUserID != reply.User_id && currentUserRoleID != 1 {
		http.Redirect(w, r, "/error?code=403&message=forbidden", http.StatusSeeOther)
		return
	}

	// Suppression de la réponse
	err = c.messagesService.DeleteReply(replyID)
	if err != nil {
		http.Redirect(w, r, "/error?code=500&message=delete_failed", http.StatusSeeOther)
		return
	}

	// Redirection vers la page du message d'origine
	http.Redirect(w, r, "/message?id="+strconv.Itoa(reply.Reply_to_id), http.StatusSeeOther)
}

// DeleteTopic gère la suppression d'un topic.
func (c *ModerationController) DeleteTopic(w http.ResponseWriter, r *http.Request) {
	topicIDStr := r.FormValue("topic_id")
	topicID, _ := strconv.Atoi(topicIDStr)

	session, _ := c.store.Get(r, "session")
	currentUserID, okUserID := session.Values["user_id"].(int)
	currentUserRoleID, okRoleID := session.Values["role_id"].(int)

	if !okUserID || !okRoleID {
		http.Redirect(w, r, "/error?code=401&message=invalid_session", http.StatusSeeOther)
		return
	}

	// Récupérer les détails du topic pour vérifier les droits et la redirection
	topic, err := c.topicsService.ReadId(topicID)
	if err != nil {
		http.Redirect(w, r, "/error?code=404&message=topic_not_found", http.StatusSeeOther)
		return
	}

	// VÉRIFICATION DES DROITS : l'utilisateur est-il le propriétaire OU un admin ?
	if currentUserID != topic.Topics.User_Id && currentUserRoleID != 1 {
		http.Redirect(w, r, "/error?code=403&message=forbidden", http.StatusSeeOther)
		return
	}

	// Suppression du topic
	err = c.topicsService.DeleteTopic(topicID)
	if err != nil {
		http.Redirect(w, r, "/error?code=500&message=delete_failed", http.StatusSeeOther)
		return
	}

	// Redirection vers la page du forum d'origine
	http.Redirect(w, r, "/forum?id="+strconv.Itoa(topic.Topics.Forum_id), http.StatusSeeOther)
}
