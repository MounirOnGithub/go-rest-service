package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/MounirOnGithub/go-rest-service/model"
	"github.com/MounirOnGithub/go-rest-service/utils"
	"github.com/Sirupsen/logrus"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

// AddUser POST a new user
func AddUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST /user")
}

// GetUserByID fetch a user by its ID
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	fmt.Fprintf(w, "GET /user/%v", userID)
}

// UpdateUserByID PUT modify a user by ID
func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	fmt.Fprintf(w, "PUT /user/%v", userID)
}

// DeleteUserByID deleting a user by its ID
func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	fmt.Fprintf(w, "DELETE /user/%v", userID)
}

// LogIn logging in the user
func LogIn(w http.ResponseWriter, r *http.Request) {
	user := model.User{}

	err := utils.GetJSONContent(&user, r)
	if err != nil {
		logrus.WithField("error= ", err).Warn("Error while retrieving user")
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
func Hello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context().Value("claims")
	fmt.Fprintf(w, "Context : \n %+v", ctx)
}
