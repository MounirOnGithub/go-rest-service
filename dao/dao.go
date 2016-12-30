package dao

import (
	"github.com/MounirOnGithub/go-rest-service/model"
	"gopkg.in/mgo.v2"
)

// Dao interface needs all operations for database
type Dao interface {
	// AddUser add a user to database
	AddUser(user *model.User) (*model.User, error)
	// GetUserByID fetch a user from database
	GetUserByID(userID string) (*model.User, error)
	// GetUserByUserName fetch a user from database with its username
	GetUserByUserName(userName string) (*model.User, error)
	// UpdateUser update a user data from database
	UpdateUser(user *model.User) (*model.User, error)
	// DeleteUser delete a user from database
	DeleteUser(userID string) error
}

// GetSession create the session to connect to our MongoDB
func GetSession() (*mgo.Session, error) {
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		return nil, err
	}
	return s, nil
}