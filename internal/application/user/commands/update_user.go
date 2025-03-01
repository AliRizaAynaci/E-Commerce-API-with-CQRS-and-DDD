package commands

import (
	"context"
	"e-commerce/internal/domain/user"
)

// UpdateUserCommand represents the command to update a user
type UpdateUserCommand struct {
	ID       string
	Email    string
	Name     string
	Password string
}

// UpdateUserHandler handles the UpdateUserCommand
type UpdateUserHandler struct {
	userRepo user.Repository
}

// NewUpdateUserHandler creates a new UpdateUserHandler
func NewUpdateUserHandler(userRepo user.Repository) *UpdateUserHandler {
	return &UpdateUserHandler{
		userRepo: userRepo,
	}
}

// Handle processes the UpdateUserCommand
func (h *UpdateUserHandler) Handle(ctx context.Context, cmd UpdateUserCommand) error {
	// Convert ID string to domain ID
	id, err := user.NewID(cmd.ID)
	if err != nil {
		return err
	}

	// Find the user
	existingUser, err := h.userRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	// Update user fields if provided
	if cmd.Email != "" && cmd.Email != existingUser.Email().String() {
		if err := existingUser.ChangeEmail(cmd.Email); err != nil {
			return err
		}
	}

	if cmd.Name != "" && cmd.Name != existingUser.Name().String() {
		if err := existingUser.ChangeName(cmd.Name); err != nil {
			return err
		}
	}

	if cmd.Password != "" {
		if err := existingUser.ChangePassword(cmd.Password); err != nil {
			return err
		}
	}

	// Save the updated user
	return h.userRepo.Update(ctx, existingUser)
}
