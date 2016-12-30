package dao

import (
	"github.com/MounirOnGithub/go-rest-service/model"
	"github.com/Sirupsen/logrus"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
)

// Mdb mocking database
type Mdb struct {
	Session *mgo.Session
}

// NewDao returns a Mdb
func NewDao(session *mgo.Session) (Dao, error) {
	dm := &Mdb{
		Session: session,
	}

	return dm, nil
}

// AddUser create a new user in db
func (dm *Mdb) AddUser(user *model.User) (*model.User, error) {
	user.ID = "42"
	session := dm.Session.Copy()

	c := session.DB("go-rest-service").C("user")
	err := c.Insert(*user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// DeleteUser delete a user from db
func (dm *Mdb) DeleteUser(userID string) error {
	logrus.WithField("user ID", userID).Debug("DeleteUser")
	return nil
}

// GetUserByID returns a user from db
func (dm *Mdb) GetUserByID(userID string) (*model.User, error) {
	logrus.WithField("user ID", userID).Debug("GetUserByID")
	return nil, nil
}

// GetUserByUserName returns a mocked user
func (dm *Mdb) GetUserByUserName(userName string) (*model.User, error) {
	logrus.WithField("username", userName).Debug("GetUserByUserName")
	return nil, errors.New("test")
}

// UpdateUser modify a user from db
func (dm *Mdb) UpdateUser(user *model.User) (*model.User, error) {
	return nil, nil
}
