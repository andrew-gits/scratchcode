package handler

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/andrew-gits/scratchcode/backend/graph"
	"github.com/andrew-gits/scratchcode/backend/graph/generated"
)

func ServerHandler(w http.ResponseWriter, r *http.Request) {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	srv.ServeHTTP(w, r)
}
