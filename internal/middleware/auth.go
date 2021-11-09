package middleware

import (
	"encoding/json"
	"net/http"
	"os"
	"together-backend/pkg"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func Auth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Authorization"] == nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "unauthorized",
			})
			return
		}

		token := pkg.BearerAuthHeader(r.Header["Authorization"][0])

		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "unauthorized",
			})
			return
		}

		var jwtKey = []byte(os.Getenv("SECRET_KEY"))
		tokenParse, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "unauthorized",
			})
			return
		}

		if !tokenParse.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "unauthorized token is invalid",
			})
			return
		}

		handler.ServeHTTP(w, r)
	}
}
