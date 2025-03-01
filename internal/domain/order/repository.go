package order

import (
	"context"
	"e-commerce/internal/domain/user"
)

// Repository defines the interface for order persistence operations
type Repository interface {
	// Save persists an order to the repository
	Save(ctx context.Context, order *Order) error

	// FindByID retrieves an order by ID
	FindByID(ctx context.Context, id ID) (*Order, error)

	// FindByUserID retrieves orders by user ID
	FindByUserID(ctx context.Context, userID user.ID, limit, offset int) ([]*Order, error)

	// Update updates an existing order
	Update(ctx context.Context, order *Order) error

	// Delete removes an order from the repository
	Delete(ctx context.Context, id ID) error

	// List retrieves all orders with pagination
	List(ctx context.Context, limit, offset int) ([]*Order, error)

	// FindByStatus retrieves orders by status
	FindByStatus(ctx context.Context, status Status, limit, offset int) ([]*Order, error)
}
