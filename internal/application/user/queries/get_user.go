package queries

import (
	"context"
	"e-commerce/internal/domain/user"
	"time"
)

// UserDTO represents the data transfer object for user information
type UserDTO struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetUserQuery represents the query to get a user by ID
type GetUserQuery struct {
	ID string
}

// GetUserHandler handles the GetUserQuery
type GetUserHandler struct {
	userRepo user.Repository
}

// NewGetUserHandler creates a new GetUserHandler
func NewGetUserHandler(userRepo user.Repository) *GetUserHandler {
	return &GetUserHandler{
		userRepo: userRepo,
	}
}

// Handle processes the GetUserQuery
func (h *GetUserHandler) Handle(ctx context.Context, query GetUserQuery) (*UserDTO, error) {
	// Convert ID string to domain ID
	id, err := user.NewID(query.ID)
	if err != nil {
		return nil, err
	}

	// Find the user
	u, err := h.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Map domain user to DTO
	return &UserDTO{
		ID:        u.ID().String(),
		Email:     u.Email().String(),
		Name:      u.Name().String(),
		CreatedAt: u.CreatedAt(),
		UpdatedAt: u.UpdatedAt(),
	}, nil
}
