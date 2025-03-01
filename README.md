# E-Commerce API with CQRS and DDD

A modern e-commerce REST API built with Go, implementing Command Query Responsibility Segregation (CQRS) and Domain-Driven Design (DDD) principles.

## Table of Contents

- [Features](#features)
- [Architecture](#architecture)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Running with Docker](#running-with-docker)
  - [Helper Scripts](#helper-scripts)
- [API Documentation](#api-documentation)
  - [User Endpoints](#user-endpoints)
  - [Product Endpoints](#product-endpoints)
  - [Cart Endpoints](#cart-endpoints)
  - [Order Endpoints](#order-endpoints)
- [Testing with Postman](#testing-with-postman)
- [Development](#development)
  - [Local Development](#local-development)
  - [Running Tests](#running-tests)
- [License](#license)

## Features

- **Clean Architecture**: Separation of concerns with domain, application, and infrastructure layers
- **CQRS Pattern**: Separate command and query models for better scalability
- **Domain-Driven Design**: Rich domain models with encapsulated business logic
- **RESTful API**: HTTP endpoints for all e-commerce operations
- **PostgreSQL Database**: Persistent storage with migrations
- **Redis Cache**: High-performance caching
- **RabbitMQ**: Message queue for asynchronous processing
- **Docker Support**: Easy deployment with Docker and Docker Compose
- **Comprehensive Testing**: Unit and integration tests

## Architecture

The project follows a clean architecture approach with the following layers:

### Domain Layer
- Contains the core business logic and domain models
- Defines value objects, entities, and aggregates
- Implements domain services and repository interfaces

### Application Layer
- Implements the CQRS pattern with separate command and query models
- Commands: Create, Update, Delete operations
- Queries: Read operations with DTOs for data transfer

### Infrastructure Layer
- Implements the repository interfaces
- Provides database access and persistence
- Handles HTTP requests and responses

## Project Structure

```
├── cmd
│   └── api
│       └── main.go           # Application entry point
├── internal
│   ├── domain
│   │   ├── user              # User domain model
│   │   ├── product           # Product domain model
│   │   ├── cart              # Cart domain model
│   │   └── order             # Order domain model
│   ├── application
│   │   ├── user              # User application services
│   │   ├── product           # Product application services
│   │   ├── cart              # Cart application services
│   │   └── order             # Order application services
│   └── infrastructure
│       ├── persistence       # Repository implementations
│       ├── api               # HTTP handlers
│       ├── database          # Database connections
│       ├── cache             # Redis client
│       └── messaging         # RabbitMQ client
├── pkg                       # Shared packages
│   └── config                # Configuration
├── migrations                # Database migrations
├── scripts                   # Helper scripts
├── docker-compose.yml        # Docker Compose configuration
├── Dockerfile                # Docker configuration
├── go.mod                    # Go module file
└── go.sum                    # Go dependencies checksum
```

## Getting Started

### Prerequisites
- Docker and Docker Compose
- Go 1.24 or higher (for local development)
- Postman (for API testing)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/AliRizaAynaci/e-commerce.git
   cd e-commerce
   ```

2. Build the application:
   ```bash
   go mod download
   ```

### Running with Docker

1. Start the application using the development script:
   ```bash
   scripts\dev.bat
   ```
   This will:
   - Build and start all containers (app, PostgreSQL, Redis, RabbitMQ)
   - Show logs from the application

2. Run database migrations:
   ```bash
   scripts\migrate.bat up
   ```

3. Access the API at http://localhost:3000

### Helper Scripts

The project includes several helper scripts in the `scripts` directory to simplify common tasks:

- `dev.bat`: Builds and starts all containers with logs
- `start.bat`: Starts containers without rebuilding
- `stop.bat`: Stops all containers
- `migrate.bat up|down`: Runs database migrations (up or down)
- `status.bat`: Checks container status
- `logs.bat [service]`: Views logs for all services or a specific service

## API Documentation

### User Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/users` | Create a new user |
| GET | `/api/users/:id` | Get a user by ID |
| PUT | `/api/users/:id` | Update a user |
| DELETE | `/api/users/:id` | Delete a user |
| GET | `/api/users?limit=10&offset=0` | List users with pagination |

### Product Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/products` | Create a new product |
| GET | `/api/products/:id` | Get a product by ID |
| PUT | `/api/products/:id` | Update a product |
| DELETE | `/api/products/:id` | Delete a product |
| GET | `/api/products?limit=10&offset=0` | List products with pagination |
| GET | `/api/products/search?query=keyword` | Search products |

### Cart Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/carts` | Create a new cart |
| GET | `/api/carts/:id` | Get a cart by ID |
| PUT | `/api/carts/:id/items` | Add item to cart |
| DELETE | `/api/carts/:id/items/:productId` | Remove item from cart |
| GET | `/api/carts/user/:userId` | Get cart by user ID |

### Order Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/orders` | Create a new order |
| GET | `/api/orders/:id` | Get an order by ID |
| PUT | `/api/orders/:id/status` | Update order status |
| GET | `/api/orders/user/:userId` | Get orders by user ID |

## Testing with Postman

You can test the API endpoints using Postman:

1. **Create a User**
   - Method: POST
   - URL: `http://localhost:3000/api/users`
   - Body (raw JSON):
     ```json
     {
       "email": "test@example.com",
       "password": "Password123",
       "name": "Test User"
     }
     ```

2. **Get User by ID**
   - Method: GET
   - URL: `http://localhost:3000/api/users/{user_id}`
   - Replace `{user_id}` with the ID from the create user response

3. **Update User**
   - Method: PUT
   - URL: `http://localhost:3000/api/users/{user_id}`
   - Body (raw JSON):
     ```json
     {
       "name": "Updated User Name"
     }
     ```

4. **List Users**
   - Method: GET
   - URL: `http://localhost:3000/api/users?limit=10&offset=0`

5. **Health Check**
   - Method: GET
   - URL: `http://localhost:3000/health`
   - Expected response: `{"status":"ok"}`

For more detailed examples of all endpoints, import the Postman collection from the `postman` directory.

## Development

### Local Development

1. Install dependencies:
   ```bash
   go mod download
   ```

2. Run the application:
   ```bash
   go run cmd/api/main.go
   ```

### Running Tests

```bash
go test ./...
```