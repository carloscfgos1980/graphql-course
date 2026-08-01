# STEPS

## 1. Set up

### start project
 go mod init github.com/carloscfgos1980/graphql-course

### Create the boiler plate for the project
printf '//go:build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go
go mod tidy
go run github.com/99designs/gqlgen init
go mod tidy

go run github.com/99designs/gqlgen generate

### create database
sqlite3 ./data/school.db "SELECT 1;"

### Copy files from other projects to use here:
- .env
- .gitignore
- internal/database/sqlite.go

### Adjust api architecture:
1. create internal directory and move graph to this new directory
2. Adjust file gqlgen.yaml for the desired architecture
3. create cmd directory and move server.go to this directory. Also rename server as main, otherwhise I can not user "air"

### Replace the default types, input, query and mutations for the one neeed for the project and the run the cli to generate a new schema adpated to this project


## 2. migrations

1. create the migrations to categories and courses tables in migrations directory
2. run the cli to create or delete the tables

### goose migrations for sqlite3
goose -dir ./migrations sqlite3 ./data/school.db up
goose -dir ./migrations sqlite3 ./data/school.db status
goose -dir ./migrations sqlite3 ./data/school.db down


## 3. main
- Load environment variables from the repo root when running from cmd/.
- get the database path from environment variable or use default
- Initialize the database
- Initialize repositories and GraphQL server (todo)
- Create a new GraphQL server with the generated schema and resolvers (todo)
- Set up the HTTP handlers for the GraphQL playground and query endpoint
- get the port from environment variable
- Start the server and log any errors

## 4. Create category

1. Repository
CreateCategory creates a new category in the database and returns the created category.
2. resolver
type Resolver struct {
	CategoryDB *repository.CategoryRepository
}
3. main wire database to api
	categoryRepo := repository.NewCategoryRepository(db)
	// Create a new GraphQL server with the generated schema and resolvers
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{
		CategoryDB: categoryRepo,
	}}))
4. resolver
CreateCategory is the resolver for the createCategory field.

## Get categories

1. Respository
GetCategories retrieves all categories from the database.
2. Query
Categories is the resolver for the categories field.
- Retrieve all categories from the database
- For each category, retrieve the associated courses and populate the Courses field (todo)
- Return the list of categories with their associated courses

## Get single category

1. Respository
GetCategoryByID retrieves a category by its ID from the database.
2. Query
GetCategoryByID retrieves a category by its ID from the database.
- Retrieve a category by its ID from the database
- If the category is not found, return an error
- Return the category with its associated courses (if any)