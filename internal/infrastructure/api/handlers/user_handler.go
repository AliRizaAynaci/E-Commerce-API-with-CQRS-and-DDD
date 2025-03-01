package handlers

import (
	"e-commerce/internal/application/user/commands"
	"e-commerce/internal/application/user/queries"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// UserHandler handles HTTP requests related to users
type UserHandler struct {
	createUserHandler *commands.CreateUserHandler
	updateUserHandler *commands.UpdateUserHandler
	deleteUserHandler *commands.DeleteUserHandler
	getUserHandler    *queries.GetUserHandler
	listUsersHandler  *queries.ListUsersHandler
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(
	createUserHandler *commands.CreateUserHandler,
	updateUserHandler *commands.UpdateUserHandler,
	deleteUserHandler *commands.DeleteUserHandler,
	getUserHandler *queries.GetUserHandler,
	listUsersHandler *queries.ListUsersHandler,
) *UserHandler {
	return &UserHandler{
		createUserHandler: createUserHandler,
		updateUserHandler: updateUserHandler,
		deleteUserHandler: deleteUserHandler,
		getUserHandler:    getUserHandler,
		listUsersHandler:  listUsersHandler,
	}
}

// RegisterRoutes registers the user routes
func (h *UserHandler) RegisterRoutes(app *fiber.App) {
	users := app.Group("/api/users")

	users.Post("/", h.CreateUser)
	users.Get("/", h.ListUsers)
	users.Get("/:id", h.GetUser)
	users.Put("/:id", h.UpdateUser)
	users.Delete("/:id", h.DeleteUser)
}

// CreateUser handles the creation of a new user
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var cmd commands.CreateUserCommand
	if err := c.BodyParser(&cmd); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	userID, err := h.createUserHandler.Handle(c.Context(), cmd)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": userID,
	})
}

// GetUser handles retrieving a user by ID
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User ID is required",
		})
	}

	query := queries.GetUserQuery{
		ID: id,
	}

	user, err := h.getUserHandler.Handle(c.Context(), query)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.JSON(user)
}

// UpdateUser handles updating a user
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User ID is required",
		})
	}

	var cmd commands.UpdateUserCommand
	if err := c.BodyParser(&cmd); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	cmd.ID = id

	if err := h.updateUserHandler.Handle(c.Context(), cmd); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User updated successfully",
	})
}

// DeleteUser handles deleting a user
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User ID is required",
		})
	}

	cmd := commands.DeleteUserCommand{
		ID: id,
	}

	if err := h.deleteUserHandler.Handle(c.Context(), cmd); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}

// ListUsers handles listing users with pagination
func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	limitStr := c.Query("limit", "10")
	offsetStr := c.Query("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}

	query := queries.ListUsersQuery{
		Limit:  limit,
		Offset: offset,
	}

	users, err := h.listUsersHandler.Handle(c.Context(), query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(users)
}
