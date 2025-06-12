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
	Topics   Topics
	Messages Messages
	Users    Users
}

// Modele d'un join entre une réponse et un user
type Replies_Join_User struct {
	Replies            Replies
	Users              Users
	CreatedAtFormatted string
}
