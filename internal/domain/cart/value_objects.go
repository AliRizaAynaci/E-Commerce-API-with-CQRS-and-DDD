package cart

import (
	"errors"
	"strings"
)

// ID represents a cart ID value object
type ID string

// NewID creates a new cart ID
func NewID(id string) (ID, error) {
	if strings.TrimSpace(id) == "" {
		return "", errors.New("cart ID cannot be empty")
	}
	return ID(id), nil
}

// String returns the string representation of the cart ID
func (id ID) String() string {
	return string(id)
}
