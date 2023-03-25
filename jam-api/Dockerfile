FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o main ./api

EXPOSE 5500

CMD ["/app/main"]