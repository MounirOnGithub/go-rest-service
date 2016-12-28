package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/MounirOnGithub/go-rest-service/dao"
	"github.com/MounirOnGithub/go-rest-service/model"
	"github.com/MounirOnGithub/go-rest-service/utils"
	"github.com/Sirupsen/logrus"
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

// AddUser POST a new user
func (uh *UserHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	// Add fields of user struct
	username := r.FormValue("username")
	password := r.FormValue("password")
	name := r.FormValue("name")
	surname := r.FormValue("surname")

	// TODO: Password encoding

	user := model.User{
		Username: username,
		Name:     name,
		Surname:  surname,
		Password: password,
	}

	// TODO: Create a new entry in db
	// TODO: Check if the user not already exist

	logrus.Info("New user created")
	utils.JSON(w, user)
}

// GetUserByID fetch a user by its ID
func (uh *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// userID := vars["id"]

	user := model.User{}

	// TODO: Find user in database

	logrus.Info("Fetched a user")
	utils.JSONWithHTTPCode(w, user, http.StatusOK)
}

// UpdateUserByID PUT modify a user by ID
func (uh *UserHandler) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	user := model.User{
		ID: userID,
	}
	// TODO: Find user in database to update

	err := utils.GetJSONContent(&user, r)
	if err != nil {
		logrus.WithField("err= ", err).Warn("Error while retrieving user")
		utils.JSONWithHTTPCode(w, utils.MsgBadParameter, http.StatusBadRequest)
		return
	}

	logrus.Info("Updated successfully")
	utils.JSONWithHTTPCode(w, user, http.StatusOK)
}

// DeleteUserByID deleting a user by its ID
func (uh *UserHandler) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// userID := vars["id"]

	// TODO: Remove user from database
	logrus.Info("Remove successfully")
	utils.JSONWithHTTPCode(w, nil, http.StatusNoContent)
}

// LogIn logging in the user
func (uh *UserHandler) LogIn(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	user := model.User{
		Username: username,
	}

	if username == "" || password == "" {
		logrus.Warn("Error while retrieving user")
		utils.JSONWithHTTPCode(w, utils.MsgBadParameter, http.StatusBadRequest)
		return
	}

	// TODO: verify in database the user exist

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	c := utils.Claims{
		UserName: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
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

// Hello handler saying hello
func (uh *UserHandler) Hello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context().Value("claims")
	fmt.Fprintf(w, "Context : \n %+v", ctx)
}
