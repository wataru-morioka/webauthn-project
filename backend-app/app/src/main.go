package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/wataru-morioka/webauthn-project/backend-app/app/src/graphql/generated"
	"github.com/wataru-morioka/webauthn-project/backend-app/app/src/graphql/resolver"
	"github.com/wataru-morioka/webauthn-project/backend-app/app/src/auth"
	. "github.com/wataru-morioka/webauthn-project/backend-app/app/src/auth/interface"
	"github.com/go-chi/chi"
)

const defaultPort = "8080"

func main() {
	router := chi.NewRouter()
	var validation MiddlewareIntarface = auth.NewMiddleware()
	router.Use(validation.VerifyAccessToken("start"))

	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	router.Handle("/graphql", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, router))
}