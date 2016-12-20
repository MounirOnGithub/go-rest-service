package main

import (
	"net/http"

	"github.com/MounirOnGithub/go-rest-service/handler"
	"github.com/MounirOnGithub/go-rest-service/utils"
	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/meatballhat/negroni-logrus"
	"github.com/urfave/negroni"
)

func main() {
	n := negroni.New()
	n.Use(negronilogrus.NewMiddlewareFromLogger(logrus.StandardLogger(), "go-rest-service"))
	// Recovery middleware for responding 500 while having a panic
	recovery := negroni.NewRecovery()
	recovery.PrintStack = false
	n.Use(recovery)

	// Router
	r := mux.NewRouter()
	r.HandleFunc("/login", handler.LogIn).Methods(http.MethodPost)

	// User sub router
	userSubRouter := mux.NewRouter().PathPrefix("/user").Subrouter().StrictSlash(true)
	userSubRouter.HandleFunc("/", handler.Hello).Methods(http.MethodGet)

	// Using middleware for the user sub router
	r.PathPrefix("/user").Handler(negroni.New(
		negroni.HandlerFunc(utils.MyMiddleware),
		negroni.Wrap(userSubRouter),
	))

	n.UseHandler(r)
	n.Run(":8080")
}
