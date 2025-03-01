package user

import "context"

// Repository defines the interface for user persistence operations
type Repository interface {
	// Save persists a user to the repository
	Save(ctx context.Context, user *User) error

	// FindByID retrieves a user by ID
	FindByID(ctx context.Context, id ID) (*User, error)

	// FindByEmail retrieves a user by email
	FindByEmail(ctx context.Context, email Email) (*User, error)

	// Update updates an existing user
	Update(ctx context.Context, user *User) error

	// Delete removes a user from the repository
	Delete(ctx context.Context, id ID) error

	// List retrieves all users with pagination
	List(ctx context.Context, limit, offset int) ([]*User, error)
}
