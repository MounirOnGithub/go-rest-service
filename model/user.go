package model

const (
	// RoleUser CRU
	RoleUser string = "User"
	// RoleAdmin CRUD
	RoleAdmin string = "Admin"
)

// User model User
type User struct {
	ID       string   `json:"id" bson:"_id" valid:"-"`
	Username string   `json:"username" bson:"username" valid:"required"`
	Password string   `json:"password" bson:"password" valid:"required"`
	Roles    []string `json:"roles" bson:"roles" valid:"required"`
	Enabled  bool     `json:"enabled" bson:"enabled" valid:"required"`
	Surname  string   `json:"surname" bson:"surname" valid:"-"`
	Name     string   `json:"name" bson:"name" valid:"-"`
}

// Authentification Response when logging in with success
type Authentification struct {
	Token string `json:"token"`
}
