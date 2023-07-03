package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-pg/pg/v10"
	"github.com/rafaelcarvalhocaetano/meetup/graphql"
	"github.com/rafaelcarvalhocaetano/meetup/postgres"
)

const defaultPort = "8080"

func main() {

	DB := postgres.New(&pg.Options{
		User:     "root",
		Password: "postgres",
		Database: "meetup_dev_demo",
	})

	defer DB.Close()

	DB.AddQueryHook(postgres.DBLogger{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{Resolvers: &graphql.Resolver{
		MeetupRepo: &postgres.MeetupRepo{DB: DB},
		UserRepo:   &postgres.UserRepo{DB: DB},
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
