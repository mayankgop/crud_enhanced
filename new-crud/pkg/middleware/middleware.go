package middleware

import (
	"crud/pkg/logger"
	login "crud/pkg/signin"
	"strings"

	// "go/token"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("mayank")

func Authorize(s http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// map value mv
		mv := r.Header.Get("Authorization")
		if mv == "" {
			w.WriteHeader(http.StatusBadRequest)
			logger.Logger.DPanic("not Authorized")
			return
		}
		array := strings.Fields(mv)
		if len(array) != 2 {
			w.WriteHeader(http.StatusBadRequest)
			logger.Logger.DPanic("not Authorized")
			return
		}

		if strings.ToLower(array[0]) != "bearer" {
			w.WriteHeader(http.StatusBadRequest)
			logger.Logger.DPanic("not Authorized")
			return
		}
		tkstr := array[1]

		claims := &login.Claims{}

		tkn, err := jwt.ParseWithClaims(tkstr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		s(w, r)
	}
}
