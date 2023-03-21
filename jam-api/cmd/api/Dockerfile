FROM golang:latest

# Set the working directory to /go/libs/app
WORKDIR /app

# Copy the go.mod and go.sum files to the container
# COPY go.mod go.sum ./

# Download dependencies
# RUN go mod download

# Copy the rest of the application to the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./libs/cmd/api

# Expose the port your service will listen on
EXPOSE 8080

# Set the startup command for your container
CMD ["/app/main"]