package middleware

import (
	"net/http"
	"os"
	"strings"
)

func CORS(next http.Handler) http.Handler {
	origins := getAllowedOrigins()
	allowed := make(map[string]bool, len(origins))
	for _, o := range origins {
		allowed[o] = true
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if allowed[origin] {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		w.Header().Set("Vary", "Origin")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func getAllowedOrigins() []string {
	if val := os.Getenv("ALLOWED_ORIGINS"); val != "" {
		return strings.Split(val, ",")
	}
	return []string{"http://localhost:5173", "http://localhost:3000"}
}
