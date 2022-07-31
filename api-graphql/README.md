# jam

##create and update
go run github.com/99designs/gqlgen init
go run github.com/99designs/gqlgen generate

##run
go run server.go

##.env file
MONGO_URL="mongodb://localhost:27017"
MONGO_DATABASE="database_name_example"
ENV=DEV
PORT=8888