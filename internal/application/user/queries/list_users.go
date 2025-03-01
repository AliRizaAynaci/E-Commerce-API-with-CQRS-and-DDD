package queries

import (
	"context"
	"e-commerce/internal/domain/user"
)

// ListUsersQuery represents the query to list users with pagination
type ListUsersQuery struct {
	Limit  int
	Offset int
}

// ListUsersHandler handles the ListUsersQuery
type ListUsersHandler struct {
	userRepo user.Repository
}

// NewListUsersHandler creates a new ListUsersHandler
func NewListUsersHandler(userRepo user.Repository) *ListUsersHandler {
	return &ListUsersHandler{
		userRepo: userRepo,
	}
}

// Handle processes the ListUsersQuery
func (h *ListUsersHandler) Handle(ctx context.Context, query ListUsersQuery) ([]*UserDTO, error) {
	// Set default values if not provided
	limit := query.Limit
	if limit <= 0 {
		limit = 10
	}

	offset := query.Offset
	if offset < 0 {
		offset = 0
	}

	// Get users from repository
	users, err := h.userRepo.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	// Map domain users to DTOs
	result := make([]*UserDTO, len(users))
	for i, u := range users {
		result[i] = &UserDTO{
			ID:        u.ID().String(),
			Email:     u.Email().String(),
			Name:      u.Name().String(),
			CreatedAt: u.CreatedAt(),
			UpdatedAt: u.UpdatedAt(),
		}
	}

	return result, nil
}
