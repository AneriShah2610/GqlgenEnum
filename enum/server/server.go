package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/aneri/enum"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("Enum", "/query"))
	http.Handle("/query", handler.GraphQL(
		enum.NewExecutableSchema(
			enum.Config{
				Resolvers: &enum.Resolver{},
				Directives: enum.DirectiveRoot{
					EnumLogging:  enum.EnumLogging,
					FieldLogging: enum.FieldLogging,
				},
			})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
