package model

// User model User
type User struct {
	ID       string `json:"id" bson:"_id" valid:"-"`
	Username string `json:"username" bson:"username" valid:"required"`
	Password string `json:"password" bson:"password" valid:"required"`
	Surname  string `json:"surname" bson:"surname" valid:"-"`
	Name     string `json:"name" bson:"name" valid:"-"`
}

// Authentification Response when logging in with success
type Authentification struct {
	Token string `json:"token"`
}
