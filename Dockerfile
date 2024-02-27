# Use the official Go image as a base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to ensure dependencies are downloaded efficiently
COPY go.mod go.sum ./

# Download and install dependencies
RUN go mod download

# Copy the entire application source code to the container's workspace
COPY . .

# Build the Go application
RUN go build -o main

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
