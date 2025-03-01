package persistence

import (
	"context"
	"database/sql"
	"e-commerce/internal/domain/user"
	"errors"
	"time"
)

// UserRepository implements the user.Repository interface
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new UserRepository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Save persists a user to the database
func (r *UserRepository) Save(ctx context.Context, user *user.User) error {
	query := `
		INSERT INTO users (id, email, password, name, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		user.ID().String(),
		user.Email().String(),
		user.Password().String(),
		user.Name().String(),
		user.CreatedAt(),
		user.UpdatedAt(),
	)

	return err
}

// FindByID retrieves a user by ID
func (r *UserRepository) FindByID(ctx context.Context, id user.ID) (*user.User, error) {
	query := `
		SELECT id, email, password, name, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	row := r.db.QueryRowContext(ctx, query, id.String())
	return r.scanUser(row)
}

// FindByEmail retrieves a user by email
func (r *UserRepository) FindByEmail(ctx context.Context, email user.Email) (*user.User, error) {
	query := `
		SELECT id, email, password, name, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	row := r.db.QueryRowContext(ctx, query, email.String())
	return r.scanUser(row)
}

// Update updates an existing user
func (r *UserRepository) Update(ctx context.Context, user *user.User) error {
	query := `
		UPDATE users
		SET email = $1, password = $2, name = $3, updated_at = $4
		WHERE id = $5
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		user.Email().String(),
		user.Password().String(),
		user.Name().String(),
		user.UpdatedAt(),
		user.ID().String(),
	)

	return err
}

// Delete removes a user from the database
func (r *UserRepository) Delete(ctx context.Context, id user.ID) error {
	query := `
		DELETE FROM users
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query, id.String())
	return err
}

// List retrieves all users with pagination
func (r *UserRepository) List(ctx context.Context, limit, offset int) ([]*user.User, error) {
	query := `
		SELECT id, email, password, name, created_at, updated_at
		FROM users
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*user.User
	for rows.Next() {
		user, err := r.scanUserFromRows(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// scanUser scans a user from a row
func (r *UserRepository) scanUser(row *sql.Row) (*user.User, error) {
	var id, email, password, name string
	var createdAt, updatedAt time.Time

	if err := row.Scan(&id, &email, &password, &name, &createdAt, &updatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// Reconstruct the user from database values
	// In a real implementation, we would have a method to reconstruct a user
	// from persistence with the exact same ID and timestamps
	u, err := user.NewUser(email, password, name)
	if err != nil {
		return nil, err
	}

	// Note: In a real implementation, you would have a method like:
	// user.Reconstruct(id, email, password, name, createdAt, updatedAt)

	return u, nil
}

// scanUserFromRows scans a user from rows
func (r *UserRepository) scanUserFromRows(rows *sql.Rows) (*user.User, error) {
	var id, email, password, name string
	var createdAt, updatedAt time.Time

	if err := rows.Scan(&id, &email, &password, &name, &createdAt, &updatedAt); err != nil {
		return nil, err
	}

	// Reconstruct the user from database values
	u, err := user.NewUser(email, password, name)
	if err != nil {
		return nil, err
	}

	return u, nil
}
