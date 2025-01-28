package middleware

import (
	"log"
	"net/http"

	"github.com/go-fuego/fuego"
	"github.com/litegral/freepass-2025/internal/lib/jwt"
)

// RequireRole creates a middleware that checks if the user has the required role
func RequireRole(role string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, ok := r.Context().Value(jwt.ClaimsContextKey).(*jwt.Claims)
			if !ok {
				log.Printf("No claims found in context")
				fuego.SendError(w, r, fuego.UnauthorizedError{Title: "No authorization found"})
				return
			}

			if claims.Role != role {
				log.Printf("User role %s does not match required role %s", claims.Role, role)
				fuego.SendError(w, r, fuego.UnauthorizedError{Title: "Unauthorized: Insufficient privileges"})
				return
			}

			next.ServeHTTP(w, r)
		})
	}
} 