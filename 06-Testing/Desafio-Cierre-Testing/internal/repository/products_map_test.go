package repository_test

import (
	"testing"
	"testing/desafio/internal"
	"testing/desafio/internal/repository"

	"github.com/stretchr/testify/require"
)

// Tests for SearchProducts including the cases:
// Case 1: The db is nil and the query is valid.
// Case 2: The db is nil and the query is nil.
// Case 3: The db is empty and the query is invalid.
// Case 4: The db is not empty and the query is valid but it doesnt match anything.
// Case 5: The db is not empty and the query is valid and it matches something.

func TestProductsMap_SearchProducts(t *testing.T) {
	// Case 1: The db is nil and the query is valid.
	t.Run("The db is nil and the query is valid", func(t *testing.T) {
		// Given
		var db map[int]internal.Product
		repo := repository.NewProductsMap(db)
		query := internal.ProductQuery{Id: 1}

		// When
		result, err := repo.SearchProducts(query)

		// Then
		require.NoError(t, err)
		require.Empty(t, result)
	})

	// Case 2: The db is nil and the query is nil.
	t.Run("The db is nil and the query is nil", func(t *testing.T) {
		// Given
		var db map[int]internal.Product
		repo := repository.NewProductsMap(db)
		//query := internal.ProductQuery{Id: -1}
		// query is nil
		var query internal.ProductQuery

		// When
		result, err := repo.SearchProducts(query)

		// Then
		require.NoError(t, err)
		require.Empty(t, result)
	})

	t.Run("The db is empty and the query is invalid", func(t *testing.T) {
		// Given
		db := make(map[int]internal.Product)
		repo := repository.NewProductsMap(db)
		query := internal.ProductQuery{Id: -1}

		// When
		result, err := repo.SearchProducts(query)

		// Then
		require.NoError(t, err)
		require.Empty(t, result)
	})

	t.Run("The db is not empty and the query is valid but it doesnt match anything", func(t *testing.T) {
		// Given
		db := map[int]internal.Product{
			1: {
				Id: 1,
				ProductAttributes: internal.ProductAttributes{
					Description: "Product 1",
					Price:       10.0,
					SellerId:    1,
				},
			},
		}
		repo := repository.NewProductsMap(db)
		query := internal.ProductQuery{Id: 2}

		// When
		result, err := repo.SearchProducts(query)

		// Then
		require.NoError(t, err)
		require.Empty(t, result)
	})

	t.Run("The db is not empty and the query is valid and it matches something", func(t *testing.T) {
		// Given
		db := map[int]internal.Product{
			1: {
				Id: 1,
				ProductAttributes: internal.ProductAttributes{
					Description: "Product 1",
					Price:       10.0,
					SellerId:    1,
				},
			},
		}
		repo := repository.NewProductsMap(db)
		query := internal.ProductQuery{Id: 1}
		expectedMap := map[int]internal.Product{
			1: {
				Id: 1,
				ProductAttributes: internal.ProductAttributes{
					Description: "Product 1",
					Price:       10.0,
					SellerId:    1,
				},
			},
		}

		// When
		result, err := repo.SearchProducts(query)

		// Then
		require.NoError(t, err)
		require.Equal(t, expectedMap, result)
	})

}
