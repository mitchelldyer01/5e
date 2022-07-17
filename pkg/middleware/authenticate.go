package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/mitchelldyer01/5e/pkg/models"
)

func Authenticate(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ignore these paths
		switch r.URL.Path {
		case "/player/login", "/player/register":
			h.ServeHTTP(w, r)
			return
		}

		var t string

		c, err := r.Cookie("token")
		a := r.Header.Get("Authorization")

		if err != nil {
			if len(a) < 1 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			t = strings.Split(a, " ")[1]
		} else {
			t = c.Value
		}

		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(t, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(w, r)
	})
}
