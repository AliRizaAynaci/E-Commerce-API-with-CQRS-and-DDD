package order

import (
	"e-commerce/internal/domain/product"
	"e-commerce/internal/domain/user"
	"errors"
	"time"

	"github.com/google/uuid"
)

// Order errors
var (
	ErrInvalidID              = errors.New("invalid order ID")
	ErrInvalidUserID          = errors.New("invalid user ID")
	ErrInvalidStatus          = errors.New("invalid order status")
	ErrInvalidTotalAmount     = errors.New("invalid total amount")
	ErrInvalidShippingAddress = errors.New("invalid shipping address")
	ErrInvalidBillingAddress  = errors.New("invalid billing address")
	ErrInvalidPaymentMethod   = errors.New("invalid payment method")
	ErrInvalidProductID       = errors.New("invalid product ID")
	ErrInvalidQuantity        = errors.New("invalid quantity")
	ErrInvalidPrice           = errors.New("invalid price")
	ErrItemNotFound           = errors.New("item not found in order")
)

// Status represents the status of an order
type Status string

const (
	StatusPending   Status = "pending"
	StatusPaid      Status = "paid"
	StatusShipped   Status = "shipped"
	StatusDelivered Status = "delivered"
	StatusCancelled Status = "cancelled"
)

// OrderItem represents an item in an order
type OrderItem struct {
	id        ID
	productID product.ID
	quantity  int
	price     float64
	createdAt time.Time
	updatedAt time.Time
}

// NewOrderItem creates a new order item
func NewOrderItem(productID string, quantity int, price float64) (*OrderItem, error) {
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

	if price <= 0 {
		return nil, ErrInvalidPrice
	}

	now := time.Now()

	return &OrderItem{
		id:        id,
		productID: productIDVO,
		quantity:  quantity,
		price:     price,
		createdAt: now,
		updatedAt: now,
	}, nil
}

// ID returns the order item ID
func (oi *OrderItem) ID() ID {
	return oi.id
}

// ProductID returns the product ID
func (oi *OrderItem) ProductID() product.ID {
	return oi.productID
}

// Quantity returns the quantity
func (oi *OrderItem) Quantity() int {
	return oi.quantity
}

// Price returns the price
func (oi *OrderItem) Price() float64 {
	return oi.price
}

// Subtotal returns the subtotal for this item
func (oi *OrderItem) Subtotal() float64 {
	return float64(oi.quantity) * oi.price
}

// CreatedAt returns the order item creation time
func (oi *OrderItem) CreatedAt() time.Time {
	return oi.createdAt
}

// UpdatedAt returns the order item last update time
func (oi *OrderItem) UpdatedAt() time.Time {
	return oi.updatedAt
}

// Order represents the order aggregate root
type Order struct {
	id              ID
	userID          user.ID
	status          Status
	totalAmount     float64
	shippingAddress string
	billingAddress  string
	paymentMethod   string
	items           []*OrderItem
	createdAt       time.Time
	updatedAt       time.Time
}

// NewOrder creates a new order
func NewOrder(userID string, shippingAddress, billingAddress, paymentMethod string) (*Order, error) {
	id, err := NewID(uuid.New().String())
	if err != nil {
		return nil, err
	}

	userIDVO, err := user.NewID(userID)
	if err != nil {
		return nil, ErrInvalidUserID
	}

	if shippingAddress == "" {
		return nil, ErrInvalidShippingAddress
	}

	if billingAddress == "" {
		return nil, ErrInvalidBillingAddress
	}

	if paymentMethod == "" {
		return nil, ErrInvalidPaymentMethod
	}

	now := time.Now()

	return &Order{
		id:              id,
		userID:          userIDVO,
		status:          StatusPending,
		totalAmount:     0,
		shippingAddress: shippingAddress,
		billingAddress:  billingAddress,
		paymentMethod:   paymentMethod,
		items:           []*OrderItem{},
		createdAt:       now,
		updatedAt:       now,
	}, nil
}

// ID returns the order ID
func (o *Order) ID() ID {
	return o.id
}

// UserID returns the user ID
func (o *Order) UserID() user.ID {
	return o.userID
}

// Status returns the order status
func (o *Order) Status() Status {
	return o.status
}

// TotalAmount returns the total amount
func (o *Order) TotalAmount() float64 {
	return o.totalAmount
}

// ShippingAddress returns the shipping address
func (o *Order) ShippingAddress() string {
	return o.shippingAddress
}

// BillingAddress returns the billing address
func (o *Order) BillingAddress() string {
	return o.billingAddress
}

// PaymentMethod returns the payment method
func (o *Order) PaymentMethod() string {
	return o.paymentMethod
}

// Items returns the order items
func (o *Order) Items() []*OrderItem {
	return o.items
}

// CreatedAt returns the order creation time
func (o *Order) CreatedAt() time.Time {
	return o.createdAt
}

// UpdatedAt returns the order last update time
func (o *Order) UpdatedAt() time.Time {
	return o.updatedAt
}

// AddItem adds an item to the order
func (o *Order) AddItem(productID string, quantity int, price float64) error {
	if o.status != StatusPending {
		return errors.New("cannot modify a non-pending order")
	}

	item, err := NewOrderItem(productID, quantity, price)
	if err != nil {
		return err
	}

	o.items = append(o.items, item)
	o.recalculateTotalAmount()
	o.updatedAt = time.Now()
	return nil
}

// RemoveItem removes an item from the order
func (o *Order) RemoveItem(itemID string) error {
	if o.status != StatusPending {
		return errors.New("cannot modify a non-pending order")
	}

	for i, item := range o.items {
		if item.id.String() == itemID {
			o.items = append(o.items[:i], o.items[i+1:]...)
			o.recalculateTotalAmount()
			o.updatedAt = time.Now()
			return nil
		}
	}
	return ErrItemNotFound
}

// ChangeStatus changes the order status
func (o *Order) ChangeStatus(status Status) error {
	validStatuses := map[Status]bool{
		StatusPending:   true,
		StatusPaid:      true,
		StatusShipped:   true,
		StatusDelivered: true,
		StatusCancelled: true,
	}

	if !validStatuses[status] {
		return ErrInvalidStatus
	}

	o.status = status
	o.updatedAt = time.Now()
	return nil
}

// ChangeShippingAddress changes the shipping address
func (o *Order) ChangeShippingAddress(address string) error {
	if address == "" {
		return ErrInvalidShippingAddress
	}

	o.shippingAddress = address
	o.updatedAt = time.Now()
	return nil
}

// ChangeBillingAddress changes the billing address
func (o *Order) ChangeBillingAddress(address string) error {
	if address == "" {
		return ErrInvalidBillingAddress
	}

	o.billingAddress = address
	o.updatedAt = time.Now()
	return nil
}

// ChangePaymentMethod changes the payment method
func (o *Order) ChangePaymentMethod(method string) error {
	if method == "" {
		return ErrInvalidPaymentMethod
	}

	o.paymentMethod = method
	o.updatedAt = time.Now()
	return nil
}

// recalculateTotalAmount recalculates the total amount of the order
func (o *Order) recalculateTotalAmount() {
	total := 0.0
	for _, item := range o.items {
		total += item.Subtotal()
	}
	o.totalAmount = total
}

// ItemCount returns the number of items in the order
func (o *Order) ItemCount() int {
	return len(o.items)
}

// TotalItems returns the total number of items in the order
func (o *Order) TotalItems() int {
	total := 0
	for _, item := range o.items {
		total += item.quantity
	}
	return total
}

// IsPending checks if the order is pending
func (o *Order) IsPending() bool {
	return o.status == StatusPending
}

// IsPaid checks if the order is paid
func (o *Order) IsPaid() bool {
	return o.status == StatusPaid
}

// IsShipped checks if the order is shipped
func (o *Order) IsShipped() bool {
	return o.status == StatusShipped
}

// IsDelivered checks if the order is delivered
func (o *Order) IsDelivered() bool {
	return o.status == StatusDelivered
}

// IsCancelled checks if the order is cancelled
func (o *Order) IsCancelled() bool {
	return o.status == StatusCancelled
}
