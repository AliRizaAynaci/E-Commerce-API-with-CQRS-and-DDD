package order

import (
	"errors"
	"strings"
)

// ID represents an order ID value object
type ID string

// NewID creates a new order ID
func NewID(id string) (ID, error) {
	if strings.TrimSpace(id) == "" {
		return "", errors.New("order ID cannot be empty")
	}
	return ID(id), nil
}

// String returns the string representation of the order ID
func (id ID) String() string {
	return string(id)
}
