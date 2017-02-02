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
	return
}

// RolesVerificationMiddleware check permissions
func RolesVerificationMiddleware(s []string) func(http.ResponseWriter, *http.Request, http.HandlerFunc) {
	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		logrus.Info("Role", s)
		claims := GetClaimsFromContext(r)
		roles := claims.Roles

		if !claims.Enabled {
			logrus.WithField("User enabled", claims.Enabled).Warn("User not enabled")
			JSONWithHTTPCode(rw, MsgTokenIsRevoked, http.StatusForbidden)
			return
		}

		for _, v := range s {
			if !isAllowed(roles, v) {
				logrus.WithField("role", s).Warn("Forbidden")
				JSONWithHTTPCode(rw, MsgTokenIsRevoked, http.StatusForbidden)
				return
			}
		}

		next(rw, r)
	}
}

// isAllowed check if the role is expected or not
func isAllowed(r []string, expected string) bool {
	if len(r) == 0 {
		return false
	}

	for _, v := range r {
		if v == expected {
			return true
		}
	}
	return false
}
