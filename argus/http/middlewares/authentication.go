package middlewares

import (
	"net/http"
	"os"
	"strings"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if has authentication header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
			return
		}

		// Extract token if needed
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token != os.Getenv("GLAUCUS_SMS_SERVICE_TOKEN") { // Use same token as sms service for now
			http.Error(w, "Empty Bearer token", http.StatusUnauthorized)
			return
		}

		// Token is present and non-empty
		next.ServeHTTP(w, r)
	})
}
