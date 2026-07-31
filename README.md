# GraphQL Course

A small GraphQL API built with Go, gqlgen, and SQLite. It exposes queries and mutations for categories and courses, with database migrations managed through Goose.

## Features

- GraphQL schema with categories and courses
- CRUD operations for both entities
- SQLite persistence
- gqlgen-generated resolvers and models
- Playground available at runtime

## Prerequisites

- Go 1.26+
- SQLite CLI (optional, for inspecting the database)

## Setup

1. Install dependencies:

   ```bash
   go mod tidy
   ```

2. Run database migrations:

   ```bash
   goose -dir ./migrations sqlite3 ./data/school.db up
   ```

3. Set the required environment variables:

   ```bash
   export PORT=8080
   export DATABASE_PATH=./data/school.db
   ```

## Run the server

```bash
go run ./cmd
```

Once started, open:

- GraphQL Playground: http://localhost:8080/
- GraphQL endpoint: http://localhost:8080/query

## Project structure

- cmd/ - application entry point
- internal/graph/ - GraphQL schema, generated code, and resolvers
- internal/repository/ - repository implementations
- internal/database/ - database initialization
- migrations/ - SQL migrations
- data/ - SQLite database file

## Example query

```graphql
query {
  categories {
    id
    name
    description
  }
}
```

## Example mutation

```graphql
mutation {
  createCategory(input: { name: "Programming", description: "Programming courses" }) {
    id
    name
  }
}
```
