package repository

import (
	"database/sql"
	"errors"
	"impl-storage/crud/products/internal"
)

var (
	// ErrProductNotSaved is the error that is returned when the product is not saved
	ErrProductRowsAffected = errors.New("number of rows affected is not expected")
)

// ProductMySQL is a struct that represents a MySQL implementation of the ProductRepository
type ProductMySQL struct {
	// db is a pool of database connections
	db *sql.DB
}

func NewProductMySQL(db *sql.DB) *ProductMySQL {
	return &ProductMySQL{db: db}
}

// FindAll returns all the products
func (p *ProductMySQL) FindAll() (products []internal.Product, err error) {
	query := "SELECT id, name, quantity, code_value, is_published, expiration, price FROM products"
	rows, err := p.db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var product internal.Product

		// scan the result
		err = rows.Scan(&product.ID, &product.Name, &product.Quantity, &product.CodeValue, &product.IsPublished, &product.Expiration, &product.Price)
		if err != nil {
			return
		}

		// append the product to the products slice
		products = append(products, product)
	}

	// check for errors during the iteration
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// FindByID returns a product by its ID
func (p *ProductMySQL) FindByID(id int) (product internal.Product, err error) {
	query := "SELECT id, name, quantity, code_value, is_published, expiration, price FROM products WHERE id = ?"
	row := p.db.QueryRow(query, id)
	if row.Err() != nil {
		err = row.Err()
		return
	}

	// scan the result
	err = row.Scan(&product.ID, &product.Name, &product.Quantity, &product.CodeValue, &product.IsPublished, &product.Expiration, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			err = internal.ErrProductNotFound
			return
		}
		return
	}
	return

}

// Save saves a product
func (p *ProductMySQL) Save(product *internal.Product) (err error) {
	query := "INSERT INTO products (name, quantity, code_value, is_published, expiration, price) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := p.db.Exec(query, product.Name, product.Quantity, product.CodeValue, product.IsPublished, product.Expiration, product.Price)
	if err != nil {
		return
	}

	// check result
	// - rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return
	}
	if rowsAffected != 1 {
		err = ErrProductRowsAffected
		return
	}

	// get the last inserted ID
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return
	}

	// update the ID
	product.ID = int(lastInsertId)

	return
}

// Update updates the given product
func (p *ProductMySQL) Update(product internal.Product) (err error) {
	// execute the query
	query := "UPDATE products SET name = ?, quantity = ?, code_value = ?, is_published = ?, expiration = ?, price = ? WHERE id = ?"
	result, err := p.db.Exec(query, product.Name, product.Quantity, product.CodeValue, product.IsPublished, product.Expiration, product.Price, product.ID)
	if err != nil {
		return
	}

	// check result
	// - rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return
	}
	if rowsAffected != 1 {
		err = ErrProductRowsAffected
		return
	}
	return
}

// Delete deletes a product by its ID
func (p *ProductMySQL) Delete(id int) (err error) {
	query := "DELETE FROM products WHERE id = ?"
	result, err := p.db.Exec(query, id)
	if err != nil {
		return
	}

	// check result
	// - rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return
	}
	if rowsAffected != 1 {
		err = ErrProductRowsAffected
		return
	}
	return
}
