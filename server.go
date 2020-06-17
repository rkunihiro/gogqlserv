package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rkunihiro/gogqlserv/repos"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rkunihiro/gogqlserv/graph"
	"github.com/rkunihiro/gogqlserv/graph/generated"
	"github.com/rkunihiro/gogqlserv/internal"
)

const defaultPort = "8080"

func injectResolver() (*graph.Resolver, error) {
	dsn := internal.GetDSN()
	db, err := internal.NewDBClient(dsn)
	if err != nil {
		return nil, err
	}

	userRepo, err := repos.NewUserRepo(db)
	if err != nil {
		return nil, err
	}

	return graph.NewResolver(
		userRepo,
	)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	resolver, err := injectResolver()
	if err != nil {
		panic(err)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: resolver,
	}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
