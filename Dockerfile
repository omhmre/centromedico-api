# Stage 1: Build the Go binary
FROM golang:1.22-alpine AS builder

# Install git and other dependencies
RUN apk add --no-cache git

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Create a minimal image
FROM alpine:latest

# Install CA certificates for HTTPS/SSL
RUN apk add --no-cache ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Copy assets and necessary files
COPY --from=builder /app/app/domain/database/medical_history_tables.sql ./app/domain/database/
COPY --from=builder /app/app/templates ./app/templates
COPY --from=builder /app/menu ./menu

# Expose the port (Cloud Run defaults to 8080)
EXPOSE 8080

# Command to run the application
CMD ["./main"]
