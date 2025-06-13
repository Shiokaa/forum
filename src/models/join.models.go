package models

// Modele d'un join entre user et topic
type Topics_Join_Users struct {
	Topics             Topics
	Users              Users
	CreatedAtFormatted string
}

// Modele d'un join entre user, topic et forum
type Topics_Join_Users_Forums struct {
	Topics Topics
	Users  Users
	Forums Forums
}

// Modele d'un join entre topics, messages et users
type Topics_Join_Messages struct {
	Topics             Topics
	Messages           Messages
	Users              Users
	CreatedAtFormatted string
}

// Modele d'un join entre une réponse et un user
type Replies_Join_User struct {
	Replies            Replies
	Users              Users
	CreatedAtFormatted string
}

// Modele d'un join entre une réponse, un user et un message
type Replies_Joins_User_Message struct {
	Replies  Replies
	Users    Users
	Messages Messages
}

// Structure pour le fil d'Ariane (Breadcrumbs)
type Breadcrumb struct {
	Name string
	URL  string
}

// Modele pour regrouper une catégorie et ses forums
type CategoryWithForums struct {
	Categories Categories
	Forums     []Forums
}

// Modele pour regrouper un forum et sa catégorie
type ForumWithCategory struct {
	Forum    Forums
	Category Categories
}

// SearchResult représente un résultat générique pour la recherche globale.
type SearchResult struct {
	Type        string
	Title       string
	Description string
	URL         string
}

// FeedbackInfo regroupe les données d'un vote pour l'affichage.
type FeedbackInfo struct {
	LikeCount    int
	DislikeCount int
}
