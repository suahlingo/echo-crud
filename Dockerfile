FROM golang:1.23-alpine as builder

# Create and set working directory
WORKDIR /app

# Copy go.mod and go.sum files first for caching dependencies
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the entire project
COPY . .

# Set the working directory to /app/cmd where the main.go file is located
WORKDIR /app/cmd

# Build the application
RUN go build -buildvcs=false -o /app/appbin

# Final stage
FROM alpine:latest

# Create the directory for the application
RUN mkdir -p /app

# Set working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/appbin /app/

# Expose the port the application will run on
EXPOSE 8080

# Command to run the application
ENTRYPOINT ["/app/appbin"]