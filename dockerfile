# Stage 1: Build the Go application
FROM golang:1.20-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Stage 2: Create a minimal image
FROM scratch

# Copy the binary from the builder stage
COPY --from=builder /app/main /main

# Expose the port your application listens on (if any)
EXPOSE 8080

# Set the entrypoint for the container
ENTRYPOINT ["/main"]