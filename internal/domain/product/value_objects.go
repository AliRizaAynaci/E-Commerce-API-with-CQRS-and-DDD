package product

import (
	"strings"
)

// ID represents a product identifier
type ID string

// NewID creates a new ID
func NewID(id string) (ID, error) {
	if id == "" {
		return "", ErrInvalidID
	}
	return ID(id), nil
}

// String returns the string representation of the ID
func (id ID) String() string {
	return string(id)
}

// Name represents a product name
type Name string

// NewName creates a new Name
func NewName(name string) (Name, error) {
	trimmedName := strings.TrimSpace(name)
	if trimmedName == "" || len(trimmedName) < 2 {
		return "", ErrInvalidName
	}
	return Name(trimmedName), nil
}

// String returns the string representation of the Name
func (n Name) String() string {
	return string(n)
}

// Description represents a product description
type Description string

// NewDescription creates a new Description
func NewDescription(description string) (Description, error) {
	trimmedDescription := strings.TrimSpace(description)
	if trimmedDescription == "" {
		return "", ErrInvalidDescription
	}
	return Description(trimmedDescription), nil
}

// String returns the string representation of the Description
func (d Description) String() string {
	return string(d)
}

// Price represents a product price
type Price float64

// NewPrice creates a new Price
func NewPrice(price float64) (Price, error) {
	if price <= 0 {
		return 0, ErrInvalidPrice
	}
	return Price(price), nil
}

// Value returns the float64 value of the Price
func (p Price) Value() float64 {
	return float64(p)
}

// Stock represents a product stock
type Stock int

// NewStock creates a new Stock
func NewStock(stock int) (Stock, error) {
	if stock < 0 {
		return 0, ErrInvalidStock
	}
	return Stock(stock), nil
}

// Value returns the int value of the Stock
func (s Stock) Value() int {
	return int(s)
}
