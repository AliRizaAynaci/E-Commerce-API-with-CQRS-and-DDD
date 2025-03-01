package user

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// User errors
var (
	ErrInvalidID       = errors.New("invalid user ID")
	ErrInvalidEmail    = errors.New("invalid email address")
	ErrInvalidPassword = errors.New("invalid password")
	ErrInvalidName     = errors.New("invalid name")
)

// CartItem represents an item in a user's cart
type CartItem struct {
	ProductID string
	Quantity  int
}

// OrderStatus represents the status of an order
type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusPaid      OrderStatus = "paid"
	OrderStatusShipped   OrderStatus = "shipped"
	OrderStatusDelivered OrderStatus = "delivered"
	OrderStatusCancelled OrderStatus = "cancelled"
)

// OrderItem represents an item in an order
type OrderItem struct {
	ProductID string
	Quantity  int
	Price     float64
}

// Order represents a user's order
type Order struct {
	ID              string
	Status          OrderStatus
	TotalAmount     float64
	ShippingAddress string
	BillingAddress  string
	PaymentMethod   string
	Items           []OrderItem
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// User represents the user aggregate root
type User struct {
	id        ID
	email     Email
	password  Password
	name      Name
	cart      []CartItem
	orders    []Order
	createdAt time.Time
	updatedAt time.Time
}

// NewUser creates a new user
func NewUser(email string, password string, name string) (*User, error) {
	id, err := NewID(uuid.New().String())
	if err != nil {
		return nil, err
	}

	emailVO, err := NewEmail(email)
	if err != nil {
		return nil, err
	}

	passwordVO, err := NewPassword(password)
	if err != nil {
		return nil, err
	}

	nameVO, err := NewName(name)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	return &User{
		id:        id,
		email:     emailVO,
		password:  passwordVO,
		name:      nameVO,
		cart:      []CartItem{},
		orders:    []Order{},
		createdAt: now,
		updatedAt: now,
	}, nil
}

// ID returns the user ID
func (u *User) ID() ID {
	return u.id
}

// Email returns the user email
func (u *User) Email() Email {
	return u.email
}

// Password returns the user password
func (u *User) Password() Password {
	return u.password
}

// Name returns the user name
func (u *User) Name() Name {
	return u.name
}

// Cart returns the user's cart
func (u *User) Cart() []CartItem {
	return u.cart
}

// Orders returns the user's orders
func (u *User) Orders() []Order {
	return u.orders
}

// CreatedAt returns the user creation time
func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

// UpdatedAt returns the user last update time
func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}

// AddToCart adds a product to the user's cart
func (u *User) AddToCart(productID string, quantity int) error {
	if quantity <= 0 {
		return errors.New("quantity must be greater than zero")
	}

	// Check if product already exists in cart
	for i, item := range u.cart {
		if item.ProductID == productID {
			// Update quantity
			u.cart[i].Quantity += quantity
			u.updatedAt = time.Now()
			return nil
		}
	}

	// Add new item to cart
	u.cart = append(u.cart, CartItem{
		ProductID: productID,
		Quantity:  quantity,
	})
	u.updatedAt = time.Now()
	return nil
}

// RemoveFromCart removes a product from the user's cart
func (u *User) RemoveFromCart(productID string) error {
	for i, item := range u.cart {
		if item.ProductID == productID {
			// Remove item from cart
			u.cart = append(u.cart[:i], u.cart[i+1:]...)
			u.updatedAt = time.Now()
			return nil
		}
	}
	return errors.New("product not found in cart")
}

// UpdateCartItemQuantity updates the quantity of a product in the user's cart
func (u *User) UpdateCartItemQuantity(productID string, quantity int) error {
	if quantity <= 0 {
		return errors.New("quantity must be greater than zero")
	}

	for i, item := range u.cart {
		if item.ProductID == productID {
			// Update quantity
			u.cart[i].Quantity = quantity
			u.updatedAt = time.Now()
			return nil
		}
	}
	return errors.New("product not found in cart")
}

// ClearCart removes all items from the user's cart
func (u *User) ClearCart() {
	u.cart = []CartItem{}
	u.updatedAt = time.Now()
}

// AddOrder adds an order to the user's orders
func (u *User) AddOrder(order Order) {
	u.orders = append(u.orders, order)
	u.updatedAt = time.Now()
}

// ChangeEmail changes the user email
func (u *User) ChangeEmail(email string) error {
	emailVO, err := NewEmail(email)
	if err != nil {
		return err
	}

	u.email = emailVO
	u.updatedAt = time.Now()
	return nil
}

// ChangeName changes the user name
func (u *User) ChangeName(name string) error {
	nameVO, err := NewName(name)
	if err != nil {
		return err
	}

	u.name = nameVO
	u.updatedAt = time.Now()
	return nil
}

// ChangePassword changes the user password
func (u *User) ChangePassword(password string) error {
	passwordVO, err := NewPassword(password)
	if err != nil {
		return err
	}

	u.password = passwordVO
	u.updatedAt = time.Now()
	return nil
}
