package auth

import (
	"context"
	"net/http"
	"strings"
)

type authenticationMiddleware struct {
	SingingKey string
}

func Register(singingKey string) *authenticationMiddleware {
	return &authenticationMiddleware{
		SingingKey: singingKey,
	}
}

func (amw *authenticationMiddleware) Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

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

		email, err := ParseToken(headerParts[1], []byte(amw.SingingKey))

		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user", email)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}
