package model

// User model User
type User struct {
	ID       string `json:"id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Surname  string `json:"surname,omitempty" bson:"surname"`
	Name     string `json:"name,omitempty" bson:"name"`
	Password string `json:"password,omitempty" bson:"password"`
}

// Authentification Response when logging in with success
type Authentification struct {
	Token string `json:"token"`
}
