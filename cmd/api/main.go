package main

import (
	"e-commerce/internal/application/user/commands"
	"e-commerce/internal/application/user/queries"
	"e-commerce/internal/infrastructure/api/handlers"
	"e-commerce/internal/infrastructure/cache"
	"e-commerce/internal/infrastructure/database"
	"e-commerce/internal/infrastructure/messaging"
	"e-commerce/internal/infrastructure/persistence"
	"e-commerce/pkg/config"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/lib/pq"           // PostgreSQL driver
	_ "github.com/mattn/go-sqlite3" // SQLite driver for backward compatibility
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database
	db, err := database.NewPostgresConnection(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close(db)

	// Initialize Redis
	redisClient, err := cache.NewRedisClient(&cfg.Redis)
	if err != nil {
		log.Fatalf("Failed to initialize Redis: %v", err)
	}
	defer redisClient.Close()

	// Initialize RabbitMQ (optional)
	var rabbitMQClient *messaging.RabbitMQClient
	rabbitMQClient, err = messaging.NewRabbitMQClient(&cfg.RabbitMQ)
	if err != nil {
		log.Printf("Warning: Failed to initialize RabbitMQ: %v", err)
		log.Println("Continuing without RabbitMQ...")
	} else {
		defer rabbitMQClient.Close()
	}

	// Initialize repositories
	userRepo := persistence.NewUserRepository(db)

	// Initialize command handlers
	createUserHandler := commands.NewCreateUserHandler(userRepo)
	updateUserHandler := commands.NewUpdateUserHandler(userRepo)
	deleteUserHandler := commands.NewDeleteUserHandler(userRepo)

	// Initialize query handlers
	getUserHandler := queries.NewGetUserHandler(userRepo)
	listUsersHandler := queries.NewListUsersHandler(userRepo)

	// Initialize API handlers
	userHandler := handlers.NewUserHandler(
		createUserHandler,
		updateUserHandler,
		deleteUserHandler,
		getUserHandler,
		listUsersHandler,
	)

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	// Register routes
	userHandler.RegisterRoutes(app)

	// Default route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to the e-commerce API",
		})
	})

	// Health check route
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	// Start server in a goroutine
	go func() {
		port := cfg.Server.Port
		log.Printf("Server starting on port %s", port)
		if err := app.Listen(":" + port); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Close the server
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}
	log.Println("Server gracefully stopped")
}
