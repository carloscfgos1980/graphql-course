package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/carloscfgos1980/graphql-course/internal/database"
	"github.com/carloscfgos1980/graphql-course/internal/graph/generated"
	"github.com/carloscfgos1980/graphql-course/internal/graph/resolvers"
	"github.com/carloscfgos1980/graphql-course/internal/repository"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Load environment variables from the repo root when running from cmd/.
	_ = godotenv.Load("../.env", ".env")
	// get the database path from environment variable or use default
	dbPath := os.Getenv("DATABASE_PATH")
	if dbPath == "" {
		dbPath = "../data/school.db"
	}
	// Initialize the database
	db, err := database.InitDB(dbPath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	log.Printf("Database initialized successfully: %s", dbPath)

	categoryRepo := repository.NewCategoryRepository(db)
	courseRepo := repository.NewCourseRepository(db)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{
		CategoryDB: categoryRepo,
		CourseDB:   courseRepo,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	// get the port from environment variable
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not set")
	}
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	// Start the server and log any errors
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
