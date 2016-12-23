package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Sirupsen/logrus"
	jwt "github.com/dgrijalva/jwt-go"
)

const (
	SecretKey = "secret"
	// ResponseHeaderContentTypeKey is the key used for response content type
	ResponseHeaderContentTypeKey = "Content-Type"
	// ResponseHeaderContentTypeJSONUTF8 is the key used for UTF8 JSON response
	ResponseHeaderContentTypeJSONUTF8 = "application/json; charset=UTF-8"
)

// Claims claims of the jwt
type Claims struct {
	UserName string
	jwt.StandardClaims
}

// JWTValidationMiddleware validation of Jwt
func JWTValidationMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// do some stuff before
	authorization := r.Header.Get("Authorization")
	token := ""

	if authorization != "" {
		el := strings.Split(authorization, " ")
		fmt.Println(el)
		if len(el) == 2 {
			token = el[1]
		} else {
			logrus.WithField("Authorization header", authorization).Debug("Authorization header is wrong or malformed")
			JSONWithHTTPCode(rw, MsgTokenMalformed, http.StatusUnauthorized)
			return
		}
	}

	fmt.Println(token)

	if token == "" {
		logrus.WithField("token", token).Debug("Unable to get token.")
		JSONWithHTTPCode(rw, MsgMissingAuth, http.StatusUnauthorized)
	}

	c := &Claims{}
	jwtToken, err := jwt.ParseWithClaims(token, c, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		logrus.WithField("Authorization header", authorization).WithField("error", err).Debug("Error parsing jwt token")
		JSONWithHTTPCode(rw, MsgTokenMalformed, http.StatusUnauthorized)
		return
	}

	if claims, ok := jwtToken.Claims.(*Claims); ok && jwtToken.Valid {
		// authSum = &claims.AuthenticationSummaryLight
		ctx := context.WithValue(r.Context(), "claims", claims)
		next(rw, r.WithContext(ctx))
	} else {
		logrus.WithField("Authorization header", authorization).Debug("Invalid jwt token, check key or alg")
		JSONWithHTTPCode(rw, MsgTokenMalformed, http.StatusUnauthorized)
		return
	}

	next(rw, r)
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
