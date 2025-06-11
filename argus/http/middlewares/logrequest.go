package middlewares

import (
	"fmt"
	"net/http"
)

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		go fmt.Printf("[HTTP][Request] (%s '%s') Received.\n", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
