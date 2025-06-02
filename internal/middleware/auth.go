package middleware

import (
	"context"
	"net/http"
	"stock/pkg/auth" // import your JWT helper
)

type contextKey string

const userKey contextKey = "user"

// AuthMiddleware validates JWT from Authorization header
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		const prefix = "Bearer "
		if len(authHeader) < len(prefix) || authHeader[:len(prefix)] != prefix {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		// Strip "Bearer " prefix
		tokenStr := authHeader[len(prefix):]

		claims, err := auth.ValidateJWT(tokenStr)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUserFromContext retrieves JWT claims from request context
func GetUserFromContext(ctx context.Context) (*auth.Claims, bool) {
	claims, ok := ctx.Value(userKey).(*auth.Claims)
	return claims, ok
}
