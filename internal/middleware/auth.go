package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/go-fuego/fuego"
	"github.com/jackc/pgx/v5"
	"github.com/litegral/freepass-2025/internal/lib/db"
	"github.com/litegral/freepass-2025/internal/lib/jwt"
)

func AuthMiddleware(queries *db.Queries) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
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

			// Verify user exists in database
			_, err = queries.GetUserByID(r.Context(), int32(claims.UserID))
			if err != nil {
				if err == pgx.ErrNoRows {
					log.Printf("User not found in database: %d", claims.UserID)
					fuego.SendError(w, r, fuego.UnauthorizedError{Title: "User not found"})
					return
				}
				log.Printf("Database error: %v", err)
				fuego.SendError(w, r, fuego.InternalServerError{Title: "Internal server error"})
				return
			}

			// Add claims to request context
			ctx := context.WithValue(r.Context(), jwt.ClaimsContextKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
} 