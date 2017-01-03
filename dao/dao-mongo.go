package dao

import (
	"github.com/MounirOnGithub/go-rest-service/model"
	"gopkg.in/mgo.v2"
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

const (
	database string = "go-rest-service"
	collection string = "user"
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
	user.ID = uuid.NewV4().String()
	session := dm.Session.Copy()

	c := session.DB(database).C(collection)
	err := c.Insert(*user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// DeleteUser delete a user from db
func (dm *Mdb) DeleteUser(userID string) error {
	session := dm.Session.Copy()
	c := session.DB(database).C(collection)
	err := c.RemoveId(userID)
	if err != nil {
		return err
	}
	return nil
}

// GetUserByID returns a user from db
func (dm *Mdb) GetUserByID(userID string) (*model.User, error) {
	session := dm.Session.Copy()

	c := session.DB(database).C(collection)
	u := &model.User{}
	err := c.FindId(userID).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// GetUserByUserName returns a mocked user
func (dm *Mdb) GetUserByUserName(userName string) (*model.User, error) {
	session := dm.Session.Copy()

	c := session.DB(database).C(collection)
	u := &model.User{}
	err := c.Find(bson.M{"username": userName}).One(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// UpdateUser modify a user from db
func (dm *Mdb) UpdateUser(user *model.User) (*model.User, error) {
	session := dm.Session.Copy()
	fmt.Println(*user)

	c := session.DB(database).C(collection)
	c.UpdateId(user.ID, user)
	return user, nil
}
