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

## .env file (example)
MONGO_URL="mongodb://localhost:27017"
MONGO_DATABASE="database_name_example"
ENV=DEV
PORT=8888
