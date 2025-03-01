package commands

import (
	"context"
	"e-commerce/internal/domain/user"
)

// DeleteUserCommand represents the command to delete a user
type DeleteUserCommand struct {
	ID string
}

// DeleteUserHandler handles the DeleteUserCommand
type DeleteUserHandler struct {
	userRepo user.Repository
}

// NewDeleteUserHandler creates a new DeleteUserHandler
func NewDeleteUserHandler(userRepo user.Repository) *DeleteUserHandler {
	return &DeleteUserHandler{
		userRepo: userRepo,
	}
}

// Handle processes the DeleteUserCommand
func (h *DeleteUserHandler) Handle(ctx context.Context, cmd DeleteUserCommand) error {
	// Convert ID string to domain ID
	id, err := user.NewID(cmd.ID)
	if err != nil {
		return err
	}

	// Check if user exists
	_, err = h.userRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	// Delete the user
	return h.userRepo.Delete(ctx, id)
}
