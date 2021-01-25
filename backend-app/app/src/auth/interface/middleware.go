package auth

import (
	"net/http"
)

type MiddlewareIntarface interface {
	VerifyAccessToken(message string) func(http.Handler) http.Handler
}