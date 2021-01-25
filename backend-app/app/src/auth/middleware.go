package auth

import (
	"log"
	"net/http"
	"context"
)

var userCtxKey = &contextKey{"user"}
type contextKey struct {
	name string
}

func Middleware(message string) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			log.Print(message)
            accessToken := req.Header.Get("Authorization") 
			log.Print(accessToken)

            if IsAccessTokenValid(&accessToken) {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
                return
            }

			ctx := context.WithValue(req.Context(), userCtxKey, accessToken)
			req = req.WithContext(ctx)

            next.ServeHTTP(w, req)
        })
    }
}

func IsAccessTokenValid(token *string) bool {
	log.Print(*token)
	return true
}