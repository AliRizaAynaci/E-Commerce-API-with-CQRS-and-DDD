FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install necessary build dependencies
RUN apk add --no-cache gcc musl-dev libc-dev

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application with CGO enabled
RUN CGO_ENABLED=1 GOOS=linux go build -a -o ecommerce ./cmd/api

# Use a smaller image for the final stage
FROM alpine:latest

# Install required dependencies for runtime
RUN apk --no-cache add ca-certificates tzdata libc6-compat

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/ecommerce .

# Verify the binary exists (will fail the build if not found)
RUN ls -la /app/ecommerce && chmod +x /app/ecommerce

# Expose the application port
EXPOSE 3000

# Command to run the application
CMD ["/app/ecommerce"]
