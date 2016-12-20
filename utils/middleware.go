package utils

import (
	"fmt"
	"net/http"
)

// MyMiddleware middleware
func MyMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// do some stuff before
	fmt.Print("Hello Middleware before")
	next(rw, r)
	// do some stuff after
}
