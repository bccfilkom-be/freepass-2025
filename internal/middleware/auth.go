package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/go-fuego/fuego"
	"github.com/litegral/freepass-2025/internal/lib/jwt"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get token from Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			log.Printf("No auth header found")
			fuego.SendError(w, r, fuego.UnauthorizedError{Title: "No authorization header"})
			return
		}

		// Check if the header has the Bearer prefix
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			log.Printf("Invalid auth format: %s", authHeader)
			fuego.SendError(w, r, fuego.UnauthorizedError{Title: "Invalid authorization header format"})
			return
		}

		// Validate the token
		claims, err := jwt.ValidateToken(parts[1])
		if err != nil {
			log.Printf("Token validation failed: %v", err)
			fuego.SendError(w, r, fuego.UnauthorizedError{Title: "Invalid token"})
			return
		}

		// Add claims to request context
		ctx := context.WithValue(r.Context(), jwt.ClaimsContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
} 