package middleware

import (
	"fmt"
	"net/http"
)

// TODO: move this to the env
var bearerToken = "abc123"

func BearerTokenValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// special case the healthcheck endpoint for now since we the middleware applies to all endpoints
		if r.URL.Path == "/healthcheck" {
			next.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		var token string
		_, err := fmt.Sscanf(authHeader, "Bearer %s", &token)
		if err != nil {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		if !isValidToken(token) {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func isValidToken(token string) bool {
	return token == bearerToken
}
