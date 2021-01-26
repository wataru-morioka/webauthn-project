package auth

import (
	"net/http"
)

type MiddlewareIntarface interface {
	VerifyAccessToken() func(http.Handler) http.Handler
}