package main

import (
	"juice-subgraph/graph"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "3009"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Create the GraphQL schema
	schema := graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}})

	// Create the GraphQL server
	srv := handler.NewDefaultServer(schema)

	// Configure CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Allow all origins
		AllowedHeaders: []string{"*"},
	})

	// Create a handler with CORS middleware and error handling
	handlerWithCors := c.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rerr := recover(); rerr != nil {
				log.Println("Error:", rerr)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		}()

		srv.ServeHTTP(w, r)
	}))

	// Handle GraphQL playground
	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))

	// Handle GraphQL queries and mutations
	http.Handle("/graphql", handlerWithCors)

	log.Printf("Connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
