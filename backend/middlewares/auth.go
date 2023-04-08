package middlewares

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		tokenString := authHeader[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})
		if err != nil {
			log.Fatal(err)
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx := context.WithValue(r.Context(), "props", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			json.NewEncoder(w).Encode("Unauthorized")
		}
	})
}
