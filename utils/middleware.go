package utils

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/Sirupsen/logrus"
	jwt "github.com/dgrijalva/jwt-go"
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
			logrus.WithField("Authorization header", authorization).Warn("Authorization header is wrong or malformed")
			JSONWithHTTPCode(rw, MsgTokenMalformed, http.StatusUnauthorized)
			return
		}
	}

	fmt.Println(token)

	if token == "" {
		logrus.WithField("token", token).Warn("Unable to get token.")
		JSONWithHTTPCode(rw, MsgMissingAuth, http.StatusUnauthorized)
	}

	c := &Claims{}
	jwtToken, err := jwt.ParseWithClaims(token, c, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		logrus.WithField("Authorization header", authorization).WithField("error", err).Warn("Error parsing jwt token")
		JSONWithHTTPCode(rw, MsgTokenMalformed, http.StatusUnauthorized)
		return
	}

	if claims, ok := jwtToken.Claims.(*Claims); ok && jwtToken.Valid {
		// authSum = &claims.AuthenticationSummaryLight
		ctx := context.WithValue(r.Context(), "claims", claims)
		next(rw, r.WithContext(ctx))
		return
	}

	logrus.WithField("Authorization header", authorization).Warn("Invalid jwt token, check key or alg")
	JSONWithHTTPCode(rw, MsgTokenMalformed, http.StatusUnauthorized)

	next(rw, r)
}
