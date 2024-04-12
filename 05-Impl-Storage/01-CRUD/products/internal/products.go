package internal

import (
	"errors"
	"time"
)

var (
	// ErrProductNotFound is the error that is returned when the product is not found
	ErrProductNotFound = errors.New("product not found")
)

// Product is the struct that represents a product
type Product struct {
	// ID is the unique identifier of the product
	ID int

	// Name is the name of the product
	Name string

	// Description is the description of the product
	Quantity int

	// CodeValue is the code value of the product
	CodeValue string

	// IsPublished is the flag that indicates if the product is published
	IsPublished string

	// Expiration is the expiration date of the product
	Expiration time.Time

	// Price is the price of the product
	Price float64
}

// ProductRepository is the interface that defines the methods that a product repository should implement
type ProductRepository interface {
	// FindAll returns all the products
	FindAll() ([]Product, error)

	// FindByID returns a product by its ID
	FindByID(id int) (Product, error)

	// Save saves a product
	Save(product *Product) error

	// Update updates a product
	Update(product *Product) (Product, error)

	// Delete deletes a product by its ID
	Delete(id int) error
}
