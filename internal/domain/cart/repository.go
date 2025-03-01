package cart

import (
	"context"
	"e-commerce/internal/domain/user"
)

// Repository defines the interface for cart persistence operations
type Repository interface {
	// Save persists a cart to the repository
	Save(ctx context.Context, cart *Cart) error

	// FindByID retrieves a cart by ID
	FindByID(ctx context.Context, id ID) (*Cart, error)

	// FindByUserID retrieves a cart by user ID
	FindByUserID(ctx context.Context, userID user.ID) (*Cart, error)

	// Update updates an existing cart
	Update(ctx context.Context, cart *Cart) error

	// Delete removes a cart from the repository
	Delete(ctx context.Context, id ID) error
}
