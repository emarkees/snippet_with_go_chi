package internal

import (
	"net/http"
)

func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-policy",
			"default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts-gstatic.com")
		w.Header().Set("Refferrer-Policy", "origin-when-cross-gin")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "deny")
		w.Header().Set("X-XSS-Protection", "0")

		next.ServeHTTP(w, r)
	})
}
