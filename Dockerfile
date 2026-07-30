# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy dependency files first to utilize Docker layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application statically
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/web

# Final stage
FROM alpine:latest

WORKDIR /app

# Add required CA certificates and timezone data
RUN apk --no-cache add ca-certificates tzdata

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Inform Docker that the container is listening on the specified port at runtime.
# Typically 3000 for Fiber apps in this starter kit, but actual port is controlled by WEB_PORT env variable.
EXPOSE 3000

# Run the executable
CMD ["./main"]
