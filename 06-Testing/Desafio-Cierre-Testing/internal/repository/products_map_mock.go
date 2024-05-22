package repository

import (
	"testing/desafio/internal"

	"github.com/stretchr/testify/mock"
)

// NewProductsMapMock returns a new ProductsMapMock.
func NewProductsMapMock() *ProductsMapMock {
	return &ProductsMapMock{}
}

// ProductsMapMock is a mock for ProductsMap.
type ProductsMapMock struct {
	mock.Mock
}

// SearchProducts returns a list of products that match the query.
func (m *ProductsMapMock) SearchProducts(query internal.ProductQuery) (map[int]internal.Product, error) {
	args := m.Called(query)
	return args.Get(0).(map[int]internal.Product), args.Error(1)
}
