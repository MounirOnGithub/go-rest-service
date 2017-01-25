package utils

import (
	"context"
	"net/http"
	"strings"

	"github.com/Sirupsen/logrus"
	jwt "github.com/dgrijalva/jwt-go"
)

// Claims claims of the jwt
type Claims struct {
	UserName string
	Roles    []string
	Enabled  bool
	jwt.StandardClaims
}

// JWTValidationMiddleware validation of Jwt
func JWTValidationMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// do some stuff before
	authorization := r.Header.Get("Authorization")
	token := ""

	if authorization != "" {
		el := strings.Split(authorization, " ")
		if len(el) == 2 {
			token = el[1]
		} else {
			logrus.WithField("Authorization header", authorization).Warn("Authorization header is wrong or malformed")
			JSONWithHTTPCode(rw, MsgTokenMalformed, http.StatusUnauthorized)
			return
		}
	}

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
		ctx := context.WithValue(r.Context(), "claims", claims)
		next(rw, r.WithContext(ctx))
		return
	}

	logrus.WithField("Authorization header", authorization).Warn("Invalid jwt token, check key or alg")
	JSONWithHTTPCode(rw, MsgTokenMalformed, http.StatusUnauthorized)

	next(rw, r)
}

// RolesAndUserVerification validation of user roles and check if it is enabled
func RolesAndUserVerification(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	c := GetClaimsFromContext(r)

	// Not enabled user receive a 403 error code
	if !c.Enabled {
		logrus.WithField("Enabled", c.Enabled).Warn("User not enabled.")
		JSONWithHTTPCode(rw, MsgTokenIsRevoked, http.StatusForbidden)
		return
	}

	if len(c.Roles) == 0 {
		logrus.WithField("Roles", c.Roles).Error("User does not have any rights")
		JSONWithHTTPCode(rw, MsgTokenIsRevoked, http.StatusForbidden)
		return
	}

	next(rw, r)
}
