# Steps

[Getting Started with GraphQL and Golang](https://gabrielgomes61320.medium.com/getting-started-with-graphql-and-golang-e8f7104b1d0b)
go mod init github.com/carloscfgos1980/graphql-course

printf '//go:build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go
go mod tidy
go run github.com/99designs/gqlgen init
go mod tidy

go run github.com/99designs/gqlgen generate 

sqlite3 data.db
create table categories (name string, id string, description string);

go run cmd/server.go
