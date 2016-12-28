package dao

import (
	"github.com/MounirOnGithub/go-rest-service/model"
	"github.com/Sirupsen/logrus"
)

// Mock mocking database
type Mock struct {
	User *model.User
}

// NewDaoMock return a new Mock
func NewDaoMock() (Dao, error) {
	dm := &Mock{
		User: user,
	}

	return dm, nil
}

// AddUser return a mocked user and log its ID
func (dm *Mock) AddUser(user *model.User) (*model.User, error) {
	logrus.WithField("user", user.ID).Debug("AddUser")
	return dm.User, nil
}

// DeleteUser log user ID for Delete method
func (dm *Mock) DeleteUser(userID int) error {
	logrus.WithField("user ID", userID).Debug("DeleteUser")
	return nil
}

// GetUserByID returns a mocked user
func (dm *Mock) GetUserByID(userID int) (*model.User, error) {
	logrus.WithField("user ID", userID).Debug("GetUserByID")
	return dm.User, nil
}

// UpdateUser returns a mocked user and logs its user ID
func (dm *Mock) UpdateUser(user *model.User) (*model.User, error) {
	logrus.WithField("user ID", dm.User.ID).Debug("UpdateUser")
	return dm.User, nil
}
