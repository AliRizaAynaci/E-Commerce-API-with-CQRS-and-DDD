package cart

import (
	"e-commerce/internal/domain/product"
	"e-commerce/internal/domain/user"
	"errors"
	"time"

	"github.com/google/uuid"
)

// Cart errors
var (
	ErrInvalidID        = errors.New("invalid cart ID")
	ErrInvalidUserID    = errors.New("invalid user ID")
	ErrInvalidProductID = errors.New("invalid product ID")
	ErrInvalidQuantity  = errors.New("invalid quantity")
	ErrProductNotFound  = errors.New("product not found in cart")
)

// CartItem represents an item in a cart
type CartItem struct {
	id        ID
	productID product.ID
	quantity  int
	createdAt time.Time
	updatedAt time.Time
}

// NewCartItem creates a new cart item
func NewCartItem(productID string, quantity int) (*CartItem, error) {
	id, err := NewID(uuid.New().String())
	if err != nil {
		return nil, err
	}

	productIDVO, err := product.NewID(productID)
	if err != nil {
		return nil, ErrInvalidProductID
	}

	if quantity <= 0 {
		return nil, ErrInvalidQuantity
	}

	now := time.Now()

	return &CartItem{
		id:        id,
		productID: productIDVO,
		quantity:  quantity,
		createdAt: now,
		updatedAt: now,
	}, nil
}

// ID returns the cart item ID
func (ci *CartItem) ID() ID {
	return ci.id
}

// ProductID returns the product ID
func (ci *CartItem) ProductID() product.ID {
	return ci.productID
}

// Quantity returns the quantity
func (ci *CartItem) Quantity() int {
	return ci.quantity
}

// CreatedAt returns the cart item creation time
func (ci *CartItem) CreatedAt() time.Time {
	return ci.createdAt
}

// UpdatedAt returns the cart item last update time
func (ci *CartItem) UpdatedAt() time.Time {
	return ci.updatedAt
}

// UpdateQuantity updates the cart item quantity
func (ci *CartItem) UpdateQuantity(quantity int) error {
	if quantity <= 0 {
		return ErrInvalidQuantity
	}

	ci.quantity = quantity
	ci.updatedAt = time.Now()
	return nil
}

// IncreaseQuantity increases the cart item quantity
func (ci *CartItem) IncreaseQuantity(amount int) error {
	if amount <= 0 {
		return ErrInvalidQuantity
	}

	ci.quantity += amount
	ci.updatedAt = time.Now()
	return nil
}

// Cart represents the cart aggregate root
type Cart struct {
	id        ID
	userID    user.ID
	items     []*CartItem
	createdAt time.Time
	updatedAt time.Time
}

// NewCart creates a new cart
func NewCart(userID string) (*Cart, error) {
	id, err := NewID(uuid.New().String())
	if err != nil {
		return nil, err
	}

	userIDVO, err := user.NewID(userID)
	if err != nil {
		return nil, ErrInvalidUserID
	}

	now := time.Now()

	return &Cart{
		id:        id,
		userID:    userIDVO,
		items:     []*CartItem{},
		createdAt: now,
		updatedAt: now,
	}, nil
}

// ID returns the cart ID
func (c *Cart) ID() ID {
	return c.id
}

// UserID returns the user ID
func (c *Cart) UserID() user.ID {
	return c.userID
}

// Items returns the cart items
func (c *Cart) Items() []*CartItem {
	return c.items
}

// CreatedAt returns the cart creation time
func (c *Cart) CreatedAt() time.Time {
	return c.createdAt
}

// UpdatedAt returns the cart last update time
func (c *Cart) UpdatedAt() time.Time {
	return c.updatedAt
}

// AddItem adds an item to the cart
func (c *Cart) AddItem(productID string, quantity int) error {
	if quantity <= 0 {
		return ErrInvalidQuantity
	}

	// Check if product already exists in cart
	for _, item := range c.items {
		if item.productID.String() == productID {
			// Update quantity
			return item.IncreaseQuantity(quantity)
		}
	}

	// Add new item to cart
	item, err := NewCartItem(productID, quantity)
	if err != nil {
		return err
	}

	c.items = append(c.items, item)
	c.updatedAt = time.Now()
	return nil
}

// RemoveItem removes an item from the cart
func (c *Cart) RemoveItem(productID string) error {
	for i, item := range c.items {
		if item.productID.String() == productID {
			// Remove item from cart
			c.items = append(c.items[:i], c.items[i+1:]...)
			c.updatedAt = time.Now()
			return nil
		}
	}
	return ErrProductNotFound
}

// UpdateItemQuantity updates the quantity of an item in the cart
func (c *Cart) UpdateItemQuantity(productID string, quantity int) error {
	if quantity <= 0 {
		return ErrInvalidQuantity
	}

	for _, item := range c.items {
		if item.productID.String() == productID {
			return item.UpdateQuantity(quantity)
		}
	}
	return ErrProductNotFound
}

// Clear removes all items from the cart
func (c *Cart) Clear() {
	c.items = []*CartItem{}
	c.updatedAt = time.Now()
}

// ItemCount returns the number of items in the cart
func (c *Cart) ItemCount() int {
	return len(c.items)
}

// TotalItems returns the total number of items in the cart
func (c *Cart) TotalItems() int {
	total := 0
	for _, item := range c.items {
		total += item.quantity
	}
	return total
}

// HasItem checks if the cart has a specific product
func (c *Cart) HasItem(productID string) bool {
	for _, item := range c.items {
		if item.productID.String() == productID {
			return true
		}
	}
	return false
}

// GetItem returns a cart item by product ID
func (c *Cart) GetItem(productID string) (*CartItem, error) {
	for _, item := range c.items {
		if item.productID.String() == productID {
			return item, nil
		}
	}
	return nil, ErrProductNotFound
}
