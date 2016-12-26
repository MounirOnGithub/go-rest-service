package utils

import (
	"encoding/json"
	"net/http"
)

const (
	// SecretKey the secret for the signature
	SecretKey = "secret"
	// ResponseHeaderContentTypeKey is the key used for response content type
	ResponseHeaderContentTypeKey = "Content-Type"
	// ResponseHeaderContentTypeJSONUTF8 is the key used for UTF8 JSON response
	ResponseHeaderContentTypeJSONUTF8 = "application/json; charset=UTF-8"
)

// GetJSONContent returns the JSON content of a request
func GetJSONContent(v interface{}, r *http.Request) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}

// JSONWithHTTPCode Json Output with an HTTP code
func JSONWithHTTPCode(w http.ResponseWriter, d interface{}, code int) {
	w.Header().Set(ResponseHeaderContentTypeKey, ResponseHeaderContentTypeJSONUTF8)
	w.WriteHeader(code)
	if d != nil {
		err := json.NewEncoder(w).Encode(d)
		if err != nil {
			// panic will cause the http.StatusInternalServerError to be send to users
			panic(err)
		}
	}
}

// JSON Outputs a JSON
func JSON(w http.ResponseWriter, d interface{}) {
	JSONWithHTTPCode(w, d, http.StatusOK)
}
