package handler

import (
	"fmt"
	"net/http"

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
	fmt.Fprint(w, "POST Login handler")
}
