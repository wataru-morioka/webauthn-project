package auth

import (
	"log"
	"net/http"
	"context"
	"strings"

	"github.com/wataru-morioka/webauthn-project/backend-app/app/src/config"
)

type contextKey struct {
	name string
}

type Middleware struct{
	userCtxKey contextKey
}

func NewMiddleware() *Middleware {
	middleware := &Middleware {
		userCtxKey: contextKey{name: config.ContextKey,},
	}
	return middleware
}

func (m Middleware) VerifyAccessToken(message string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			log.Printf("validation %s", message)
			accessToken := constractAccessToken(req)

			if isAccessTokenValid(&accessToken) {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(req.Context(), m.userCtxKey, accessToken)
			req = req.WithContext(ctx)

			next.ServeHTTP(w, req)
		})
	}
}

func isAccessTokenValid(token *string) bool {
	log.Printf("access token: %s", *token)
	return true
}

func constractAccessToken(req *http.Request) string {
	bearerHeader := req.Header.Get("Authorization") 
	return strings.Replace(bearerHeader, "Bearer: ", "", 1)
}




