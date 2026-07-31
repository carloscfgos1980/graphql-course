# STEPS

printf '//go:build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go
go mod tidy
go run github.com/99designs/gqlgen init
go mod tidy


go get github.com/99designs/gqlgen
go run github.com/99designs/gqlgen init
go run github.com/99designs/gqlgen generate

goose -dir ./migrations sqlite3 ./data/school.db up
goose -dir ./migrations sqlite3 ./data/school.db status
goose -dir ./migrations sqlite3 ./data/school.db down

sqlite3 ./data/school.db "SELECT 1;"