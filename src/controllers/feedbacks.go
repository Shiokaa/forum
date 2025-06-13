package controllers

import (
	"forum/src/middlewares"
	"forum/src/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type FeedbackController struct {
	store   *sessions.CookieStore
	service *services.FeedbacksServices
}

func FeedbackControllerInit(s *sessions.CookieStore, f *services.FeedbacksServices) *FeedbackController {
	return &FeedbackController{store: s, service: f}
}

func (c *FeedbackController) FeedbackRouter(r *mux.Router) {
	r.Handle("/feedback/submit", middlewares.RequireAuth(c.store, http.HandlerFunc(c.ProcessFeedback))).Methods("POST")
}

func (c *FeedbackController) ProcessFeedback(w http.ResponseWriter, r *http.Request) {
	session, _ := c.store.Get(r, "session")
	userID := session.Values["user_id"].(int)

	messageID, _ := strconv.Atoi(r.FormValue("message_id"))
	voteType := r.FormValue("vote_type") // "like" ou "dislike"
	topicID := r.FormValue("topic_id")   // Pour la redirection

	if messageID == 0 || (voteType != "like" && voteType != "dislike") || topicID == "" {
		http.Redirect(w, r, "/error?code=400&message=bad_request", http.StatusSeeOther)
		return
	}

	err := c.service.HandleFeedback(userID, messageID, voteType)
	if err != nil {
		http.Redirect(w, r, "/error?code=500&message=feedback_failed", http.StatusSeeOther)
		return
	}

	// Redirection vers la page du topic
	http.Redirect(w, r, "/topic?id="+topicID, http.StatusSeeOther)
}
