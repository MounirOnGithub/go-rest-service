package dao

import "github.com/MounirOnGithub/go-rest-service/model"

// Dao interface needs all operations for database
type Dao interface {
	// AddUser add a user to database
	AddUser(user *model.User) (*model.User, error)
	// GetUserByID fetch a user from database
	GetUserByID(userID int) (*model.User, error)
	// UpdateUser update a user data from database
	UpdateUser(user *model.User) (*model.User, error)
	// DeleteUser delete a user from database
	DeleteUser(userID int) error
}
