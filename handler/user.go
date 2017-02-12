package handler

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/MounirOnGithub/go-rest-service/dao"
	"github.com/MounirOnGithub/go-rest-service/model"
	"github.com/MounirOnGithub/go-rest-service/utils"
	"github.com/Sirupsen/logrus"
	"github.com/asaskevich/govalidator"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

// UserHandler handler containing dao
type UserHandler struct {
	dao dao.Dao
}

// NewUserHandler returns a new UserHandler
func NewUserHandler(dao dao.Dao) UserHandler {
	return UserHandler{
		dao: dao,
	}
}

// IfUserExist check if the user already exist in db
func (uh *UserHandler) IfUserExist(userName string) bool {
	_, err := uh.dao.GetUserByUserName(userName)
	if err != nil {
		return false
	}
	return true
}

// AddUser POST a new user
func (uh *UserHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	name := r.FormValue("name")
	surname := r.FormValue("surname")

	user := &model.User{
		Username: username,
		Name:     name,
		Surname:  surname,
		Enabled:  true,
		Roles: []string{
			model.RoleUser,
		},
	}
	user.Password = base64.StdEncoding.EncodeToString([]byte(r.FormValue("password")))

	v, err := govalidator.ValidateStruct(user)
	if !v {
		logrus.WithField("valid struct", v).Error("Invalid structure")
		utils.JSONWithHTTPCode(w, utils.MsgBadParameter, http.StatusBadRequest)
		return
	}

	// Verify if the user exist
	exist := uh.IfUserExist(username)
	if exist {
		logrus.WithField("username", username).Warn("User already exists")
		utils.JSONWithHTTPCode(w, utils.MsgBadParameter, http.StatusBadRequest)
		return
	}

	user, err = uh.dao.AddUser(user)
	if err != nil {
		logrus.WithField("err=", err).Warn("Error while updating user")
	}

	logrus.Info("New user created")
	utils.JSON(w, user)
}

// GetUserByID fetch a user by its ID
func (uh *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	vars := mux.Vars(r)
	userID := vars["id"]
	user := &model.User{}
	user, err := uh.dao.GetUserByID(userID)
	if err != nil {
		logrus.WithField("err", err).Warn("Error while fetching user")
		utils.JSONWithHTTPCode(w, utils.MsgBadParameter, http.StatusBadRequest)
		return
	}

	logrus.Info("Fetched a user")
	utils.JSONWithHTTPCode(w, user, http.StatusOK)
}

// UpdateUserByID PUT modify a user by ID
func (uh *UserHandler) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	existingUser, err := uh.dao.GetUserByID(userID)
	if err != nil {
		logrus.WithField("user ID", userID).Error(err)
		utils.JSONWithHTTPCode(w, utils.MsgEntityDoesNotExist, http.StatusNotFound)
		return
	}

	user := &model.User{}
	err = utils.GetJSONContent(&user, r)
	if err != nil {
		logrus.WithField("err= ", err).Warn("Error while retrieving user")
		utils.JSONWithHTTPCode(w, utils.MsgBadParameter, http.StatusBadRequest)
		return
	}

	if len(user.Roles) != len(existingUser.Roles) {
		logrus.WithField("lenght roles", len(user.Roles)).Error("Cannot modify Roles")
		utils.JSONWithHTTPCode(w, utils.MsgBadParameter, http.StatusBadRequest)
		return
	}

	userModified, err := uh.dao.UpdateUser(user)
	fmt.Printf("user modified roles %v", userModified.Roles)
	if err != nil {
		logrus.WithField("err=", err).Warn("Error while updating user")
		utils.JSONWithHTTPCode(w, utils.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	logrus.Info("Updated successfully")
	utils.JSONWithHTTPCode(w, userModified, http.StatusOK)
}

// DeleteUserByID deleting a user by its ID
func (uh *UserHandler) DeleteUserByID(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	vars := mux.Vars(r)
	userID := vars["id"]

	err := uh.dao.DeleteUser(userID)
	if err != nil {
		logrus.WithField("err=", err).Warn("Error while deleting user")
		utils.JSONWithHTTPCode(w, utils.MsgBadParameter, http.StatusBadRequest)
		return
	}

	logrus.Info("Remove successfully")
	utils.JSONWithHTTPCode(w, nil, http.StatusNoContent)
}

// LogIn logging in the user
func (uh *UserHandler) LogIn(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := base64.StdEncoding.EncodeToString([]byte(r.FormValue("password")))
	user := &model.User{}

	if username == "" || password == "" {
		logrus.Warn("username or password empty")
		utils.JSONWithHTTPCode(w, utils.MsgBadParameter, http.StatusBadRequest)
		return
	}

	user, err := uh.dao.GetUserByUserName(username)
	if err != nil {
		logrus.WithField("err", err).Warn("Error while fetching user")
		utils.JSONWithHTTPCode(w, utils.MsgBadParameter, http.StatusBadRequest)
		return
	}

	if user.Password != password {
		logrus.Warn("Wrong password or username")
		utils.JSONWithHTTPCode(w, utils.MsgBadParameter, http.StatusBadRequest)
		return
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	c := utils.Claims{}

	if user != nil {
		c.UserName = user.Username
		c.Roles = user.Roles
		c.Enabled = user.Enabled
		c.StandardClaims = jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		}
	}

	mySigningKey := []byte(utils.SecretKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	auth := model.Authentification{}
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		logrus.WithField("token", tokenString).Warn(err)
	}
	auth.Token = tokenString

	utils.JSONWithHTTPCode(w, auth, http.StatusCreated)
}

// Hello handler saying hello to the user red from claims
func (uh *UserHandler) Hello(w http.ResponseWriter, r *http.Request) {
	c := utils.GetClaimsFromContext(r)
	fmt.Fprintf(w, " Hello %v", c.UserName)
}
