package handler

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!\n")
}

func GetGlucoses(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GET /glucoses \n")
}

func GetGlucoseByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	glucoseID := vars["id"]
	fmt.Fprintf(w, "GET /glucose/%v", glucoseID)
}

func AddGlucose(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "POST /glucoses")
}

func DeleteGlucose(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	glucoseID := vars["id"]
	fmt.Fprintf(w, "DELETE /glucose/%v", glucoseID)
}