package auth

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		// TODO: virtual variables environment
		isSingingKey := os.Getenv("singing_key")
		if isSingingKey == "" {
			fmt.Errorf("No encryption token")
		}

		singingKey := []byte(isSingingKey)

		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		headerParts := strings.Split(authHeader, " ")

		if len(headerParts) != 2 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if headerParts[0] != "Bearer" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		_, err := ParseToken(headerParts[1], singingKey)

		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

	})
}
