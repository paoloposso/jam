# jam

## create and update graphql models and resolvers
go run github.com/99designs/gqlgen init
go run github.com/99designs/gqlgen generate

## running
go run server.go

## unit testing
go test $(go list ./... | grep -v /tools)
or
go test ./libs/...

## .env file (example)\
GIN_MODE=debug


export PATH=$(go env GOPATH)/bin:$PATH
swag init

## ENVIRONMENT Variables / Access AWS
To run it locally, it's necessary to have one of the options:
- credentials file configured to access AWS
- Set environment variables by:
    - adding .env file on the same folder of the main file (cmd/api)
    - setting all necessary environment variables when running it

AWS_ACCESS_KEY=
AWS_ACCESS_KEY_SECRET=
AWS_REGION=us-east-2
GO_ENV=development