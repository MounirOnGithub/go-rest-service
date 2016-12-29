package dao

import (
	"github.com/MounirOnGithub/go-rest-service/model"
	"github.com/Sirupsen/logrus"
)

// Mdb mocking database
type Mdb struct {
	User *model.User
}

// NewDao returns a Mdb
func NewDao() (Dao, error) {
	dm := &Mdb{
		User: user,
	}

	return dm, nil
}

// AddUser create a new user in db
func (dm *Mdb) AddUser(user *model.User) (*model.User, error) {
	logrus.WithField("user", user.ID).Debug("AddUser")
	return dm.User, nil
}

// DeleteUser delete a user from db
func (dm *Mdb) DeleteUser(userID string) error {
	logrus.WithField("user ID", userID).Debug("DeleteUser")
	return nil
}

// GetUserByID returns a user from db
func (dm *Mdb) GetUserByID(userID string) (*model.User, error) {
	logrus.WithField("user ID", userID).Debug("GetUserByID")
	return dm.User, nil
}

// GetUserByUserName returns a mocked user
func (dm *Mdb) GetUserByUserName(userName string) (*model.User, error) {
	logrus.WithField("username", userName).Debug("GetUserByUserName")
	return dm.User, nil
}

// UpdateUser modify a user from db
func (dm *Mdb) UpdateUser(user *model.User) (*model.User, error) {
	logrus.WithField("user ID", dm.User.ID).Debug("UpdateUser")
	return dm.User, nil
}
