# Use the official Golang image as the base image
FROM golang:latest AS build

# Set the working directory in the container
WORKDIR /app

# Copy the Go modules manifest and lock file
COPY go.mod ./
COPY go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY src/ .

# Build the Go app
RUN go build -o main cmd/main.go

# Start a new stage from scratch
FROM golang:latest

# Set the working directory in the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=build /app/main .

# Command to run the executable
CMD ["./main"]
