package utils

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		r.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}

	return r
}
