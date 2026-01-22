# Build Stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the Go app
# Adjust the path to your main.go if different
RUN go build -o main ./cmd/server/main.go

# Run Stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Expose the application port
EXPOSE 8080

# Run the binary
CMD ["./main"]
