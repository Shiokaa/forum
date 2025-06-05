package models

// Modele d'un join entre user et topic
type Topics_Join_Users struct {
	Topics Topics
	Users  Users
}

// Modele d'un join entre user, topic et forum
type Topics_Join_Users_Forums struct {
	Topics Topics
	Users  Users
	Forums Forums
}

type Topics_Join_Messages struct {
	Topics   Topics
	Messages Messages
	Users    Users
}
