package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"together-backend/pkg"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserId int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

func Auth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

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
		var claims = &Claims{}
		tokenParse, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
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

		ctx := context.WithValue(r.Context(), "currentUserID", claims.UserId)
		handler.ServeHTTP(w, r.WithContext(ctx))
	}
}
