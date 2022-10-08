package handler

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
)

func PlaygroundHandler(w http.ResponseWriter, r *http.Request) {
	srv := playground.Handler("GraphQL playground", "/api/graphql")

	srv.ServeHTTP(w, r)
}
