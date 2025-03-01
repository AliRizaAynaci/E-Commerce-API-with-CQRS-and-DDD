package product

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Product errors
var (
	ErrInvalidID          = errors.New("invalid product ID")
	ErrInvalidName        = errors.New("invalid product name")
	ErrInvalidDescription = errors.New("invalid product description")
	ErrInvalidPrice       = errors.New("invalid product price")
	ErrInvalidStock       = errors.New("invalid product stock")
)

// Product represents the product aggregate root
type Product struct {
	id          ID
	name        Name
	description Description
	price       Price
	stock       Stock
	createdAt   time.Time
	updatedAt   time.Time
}

// NewProduct creates a new product
func NewProduct(name string, description string, price float64, stock int) (*Product, error) {
	id, err := NewID(uuid.New().String())
	if err != nil {
		return nil, err
	}

	nameVO, err := NewName(name)
	if err != nil {
		return nil, err
	}

	descriptionVO, err := NewDescription(description)
	if err != nil {
		return nil, err
	}

	priceVO, err := NewPrice(price)
	if err != nil {
		return nil, err
	}

	stockVO, err := NewStock(stock)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	return &Product{
		id:          id,
		name:        nameVO,
		description: descriptionVO,
		price:       priceVO,
		stock:       stockVO,
		createdAt:   now,
		updatedAt:   now,
	}, nil
}

// ID returns the product ID
func (p *Product) ID() ID {
	return p.id
}

// Name returns the product name
func (p *Product) Name() Name {
	return p.name
}

// Description returns the product description
func (p *Product) Description() Description {
	return p.description
}

// Price returns the product price
func (p *Product) Price() Price {
	return p.price
}

// Stock returns the product stock
func (p *Product) Stock() Stock {
	return p.stock
}

// CreatedAt returns the product creation time
func (p *Product) CreatedAt() time.Time {
	return p.createdAt
}

// UpdatedAt returns the product last update time
func (p *Product) UpdatedAt() time.Time {
	return p.updatedAt
}

// ChangeName changes the product name
func (p *Product) ChangeName(name string) error {
	nameVO, err := NewName(name)
	if err != nil {
		return err
	}

	p.name = nameVO
	p.updatedAt = time.Now()
	return nil
}

// ChangeDescription changes the product description
func (p *Product) ChangeDescription(description string) error {
	descriptionVO, err := NewDescription(description)
	if err != nil {
		return err
	}

	p.description = descriptionVO
	p.updatedAt = time.Now()
	return nil
}

// ChangePrice changes the product price
func (p *Product) ChangePrice(price float64) error {
	priceVO, err := NewPrice(price)
	if err != nil {
		return err
	}

	p.price = priceVO
	p.updatedAt = time.Now()
	return nil
}

// ChangeStock changes the product stock
func (p *Product) ChangeStock(stock int) error {
	stockVO, err := NewStock(stock)
	if err != nil {
		return err
	}

	p.stock = stockVO
	p.updatedAt = time.Now()
	return nil
}

// IncreaseStock increases the product stock
func (p *Product) IncreaseStock(quantity int) error {
	if quantity <= 0 {
		return ErrInvalidStock
	}

	newStock := p.stock.Value() + quantity
	return p.ChangeStock(newStock)
}

// DecreaseStock decreases the product stock
func (p *Product) DecreaseStock(quantity int) error {
	if quantity <= 0 {
		return ErrInvalidStock
	}

	newStock := p.stock.Value() - quantity
	if newStock < 0 {
		return ErrInvalidStock
	}

	return p.ChangeStock(newStock)
}

// IsInStock checks if the product is in stock
func (p *Product) IsInStock() bool {
	return p.stock.Value() > 0
}

// HasSufficientStock checks if the product has sufficient stock
func (p *Product) HasSufficientStock(quantity int) bool {
	return p.stock.Value() >= quantity
}
