# jam

## create and update graphql models and resolvers
go run github.com/99designs/gqlgen init
go run github.com/99designs/gqlgen generate

## running
go run server.go

## unit testing
go test $(go list ./... | grep -v /tools)
or
go test ./src/...

## .env file (example)\
GIN_MODE=debug

swag init -g src/cmd/api/main.go --output src/cmd/api/docs