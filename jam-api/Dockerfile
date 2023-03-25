FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o main ./cmd/api

EXPOSE 5500

CMD ["/app/main"]