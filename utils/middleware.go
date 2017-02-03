package utils

import (
	"context"
	"net/http"
	"strings"

	"github.com/MounirOnGithub/go-rest-service/dao"
	"github.com/MounirOnGithub/go-rest-service/model"
	"github.com/Sirupsen/logrus"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
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
		claims := GetClaimsFromContext(r)
		roles := claims.Roles

		if !claims.Enabled {
			logrus.WithField("User enabled", claims.Enabled).Warn("User not enabled")
			JSONWithHTTPCode(rw, MsgTokenIsRevoked, http.StatusForbidden)
			return
		}

		p := false
		for _, v := range s {
			p = isAllowed(roles, v)
			if p {
				break
			}
		}

		if !p {
			logrus.WithField("role expected", s).Warn("Forbidden")
			JSONWithHTTPCode(rw, MsgTokenIsRevoked, http.StatusForbidden)
			return
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

// OwningResourceMiddleware verify that the client is requesting only a ressource that it's owning
func OwningResourceMiddleware() func(http.ResponseWriter, *http.Request, http.HandlerFunc) {
	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		claims := GetClaimsFromContext(r)

		if isAllowed(claims.Roles, model.RoleAdmin) {
			next(rw, r)
		}

		var d dao.Dao
		session, err := dao.GetSession()
		if err != nil {
			logrus.WithError(err).Warn("Error while retrieving mongo session")
			JSONWithHTTPCode(rw, MsgInternalServerError, http.StatusInternalServerError)
			return
		}

		d, err = dao.NewDao(session)
		if err != nil {
			logrus.WithError(err).Warn("Error while creation a new Dao")
			JSONWithHTTPCode(rw, MsgInternalServerError, http.StatusInternalServerError)
			return
		}

		usrName := claims.UserName
		vars := mux.Vars(r)
		userID := vars["id"]

		u, err := d.GetUserByUserName(usrName)
		if err != nil {
			logrus.WithError(err).Warn("Error while retrieving user in database")
			JSONWithHTTPCode(rw, MsgEntityDoesNotExist, http.StatusNotFound)
			return
		}

		if u.ID != userID {
			logrus.WithField("user ID", userID).Error("User is not the owner of this resource")
			JSONWithHTTPCode(rw, MsgTokenIsRevoked, http.StatusForbidden)
			return
		}

		next(rw, r)
	}
}
