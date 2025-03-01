.PHONY: build run test clean docker-build docker-up docker-down lint migrate-up migrate-down

# Go related variables
BINARY_NAME=ecommerce
MAIN_PACKAGE=./cmd/api

# Docker related variables
DOCKER_COMPOSE=docker-compose

# Build the application
build:
	go build -o $(BINARY_NAME) $(MAIN_PACKAGE)

# Run the application
run:
	go run $(MAIN_PACKAGE)

# Run tests
test:
	go test -v ./...

# Clean build files
clean:
	go clean
	rm -f $(BINARY_NAME)

# Build docker image
docker-build:
	$(DOCKER_COMPOSE) build

# Start docker containers
docker-up:
	$(DOCKER_COMPOSE) up -d

# Stop docker containers
docker-down:
	$(DOCKER_COMPOSE) down

# Start docker containers and follow logs
docker-up-logs:
	$(DOCKER_COMPOSE) up

# Restart docker containers
docker-restart:
	$(DOCKER_COMPOSE) restart

# Show docker logs
docker-logs:
	$(DOCKER_COMPOSE) logs -f

# Run linter
lint:
	golangci-lint run ./...

# Install dependencies
deps:
	go mod download

# Update dependencies
deps-update:
	go get -u ./...
	go mod tidy

# Create database migrations
migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir migrations -seq $$name

# Apply database migrations
migrate-up:
	migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/ecommerce?sslmode=disable" up

# Rollback database migrations
migrate-down:
	migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/ecommerce?sslmode=disable" down 1

# Default target
all: clean build 