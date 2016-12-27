package model

// User model User
type User struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Surname  string `json:"surname,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}

// Authentification Response when logging in with success
type Authentification struct {
	Token string `json:"token"`
}
