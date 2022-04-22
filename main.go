package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/zzihyeon/golang-server.git/db"
	"github.com/zzihyeon/golang-server.git/graph"
	"github.com/zzihyeon/golang-server.git/graph/generated"
	mongodb "github.com/zzihyeon/golang-server.git/mongoDB"
	"github.com/zzihyeon/golang-server.git/restapi"

	_ "github.com/lib/pq"
)

const (
	defaultPort = "8080"
	dbDriver    = "postgres"
	dbSource    = "postgresql://test:test@localhost:5432/test?sslmode=disable"
	mongoDBURL  = "mongodb://localhost:27017"
)

func main() {
	go initGraphQLServer()
	go initMongoDB()
	go initAPI()
	go initPostgresqlDB()
	for {
	}
}

func initGraphQLServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func initAPI() {
	restapi.Initialize()
}

func initMongoDB() {
	mongodb.Init(mongoDBURL)
}

func initPostgresqlDB() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("db connection error :", err)
	}
	db.Initialize(conn)
}
