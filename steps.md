# STEPS

## Set up

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

## 2 migrations

1. create the migrations to categories and courses tables in migrations directory
2. run the cli to create or delete the tables

### goose migrations for sqlite3
goose -dir ./migrations sqlite3 ./data/school.db up
goose -dir ./migrations sqlite3 ./data/school.db status
goose -dir ./migrations sqlite3 ./data/school.db down


## main
- Load environment variables from the repo root when running from cmd/.
- get the database path from environment variable or use default
- Initialize the database
- Initialize repositories and GraphQL server (todo)
- Create a new GraphQL server with the generated schema and resolvers (todo)
- Set up the HTTP handlers for the GraphQL playground and query endpoint
- get the port from environment variable
- Start the server and log any errors