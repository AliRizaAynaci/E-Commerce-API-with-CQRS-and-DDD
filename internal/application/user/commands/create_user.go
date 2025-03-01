package commands

import (
	"context"
	"e-commerce/internal/domain/user"
)

// CreateUserCommand represents the command to create a new user
type CreateUserCommand struct {
	Email    string
	Password string
	Name     string
}

// CreateUserHandler handles the CreateUserCommand
type CreateUserHandler struct {
	userRepo user.Repository
}

// NewCreateUserHandler creates a new CreateUserHandler
func NewCreateUserHandler(userRepo user.Repository) *CreateUserHandler {
	return &CreateUserHandler{
		userRepo: userRepo,
	}
}

// Handle processes the CreateUserCommand
func (h *CreateUserHandler) Handle(ctx context.Context, cmd CreateUserCommand) (string, error) {
	// Check if user with the same email already exists
	existingUser, err := h.userRepo.FindByEmail(ctx, user.Email(cmd.Email))
	if err == nil && existingUser != nil {
		return "", user.ErrInvalidEmail
	}

	// Create a new user
	newUser, err := user.NewUser(cmd.Email, cmd.Password, cmd.Name)
	if err != nil {
		return "", err
	}

	// Save the user
	if err := h.userRepo.Save(ctx, newUser); err != nil {
		return "", err
	}

	return newUser.ID().String(), nil
}
