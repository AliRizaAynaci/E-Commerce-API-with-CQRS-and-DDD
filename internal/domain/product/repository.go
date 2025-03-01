package product

import (
	"context"
)

// Repository defines the interface for product persistence operations
type Repository interface {
	// Save persists a product to the repository
	Save(ctx context.Context, product *Product) error

	// FindByID retrieves a product by ID
	FindByID(ctx context.Context, id ID) (*Product, error)

	// Update updates an existing product
	Update(ctx context.Context, product *Product) error

	// Delete removes a product from the repository
	Delete(ctx context.Context, id ID) error

	// List retrieves all products with pagination
	List(ctx context.Context, limit, offset int) ([]*Product, error)

	// Search searches for products by name or description
	Search(ctx context.Context, query string, limit, offset int) ([]*Product, error)
}
