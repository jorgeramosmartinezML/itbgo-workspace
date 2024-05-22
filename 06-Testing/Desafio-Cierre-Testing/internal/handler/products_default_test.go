package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"testing/desafio/internal"
	"testing/desafio/internal/handler"
	"testing/desafio/internal/repository"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestProductDefault_Get(t *testing.T) {
	// Case 1: The query is valid but it doesnt match anything.
	t.Run("The query is valid but it doesnt match anything", func(t *testing.T) {
		// Given
		mockRepo := repository.NewProductsMapMock()
		mockRepo.On("SearchProducts", mock.Anything).Return(map[int]internal.Product{}, nil)
		productsHandler := handler.NewProductsDefault(mockRepo)

		// When
		request := httptest.NewRequest("GET", "/products?id=1", nil)
		response := httptest.NewRecorder()
		productsHandler.Get()(response, request)

		// Then
		expectedCode := http.StatusOK
		expectedBody := `{"message":"success","data":{}}`
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		mockRepo.AssertExpectations(t)
	})

	// Case 2: The query is valid and it matches a product.
	t.Run("The query is valid and it matches a product", func(t *testing.T) {
		// Given
		mockRepo := repository.NewProductsMapMock()
		mockRepo.On("SearchProducts", mock.Anything).Return(map[int]internal.Product{
			1: {
				Id: 1,
				ProductAttributes: internal.ProductAttributes{
					Description: "Product 1",
					Price:       10.0,
					SellerId:    1,
				},
			},
		}, nil)
		productsHandler := handler.NewProductsDefault(mockRepo)

		// When
		request := httptest.NewRequest("GET", "/products?id=1", nil)
		response := httptest.NewRecorder()
		productsHandler.Get()(response, request)

		// Then
		expectedCode := http.StatusOK
		expectedBody := `{"message":"success","data":{"1":{"id":1,"description":"Product 1","price":10.0,"seller_id":1}}}`
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		mockRepo.AssertExpectations(t)
	})

	// Case 3: The query is invalid.
	t.Run("The query is invalid", func(t *testing.T) {
		// Given
		mockRepo := repository.NewProductsMapMock()
		productsHandler := handler.NewProductsDefault(mockRepo)

		// When
		request := httptest.NewRequest("GET", "/products?id=invalid", nil)
		response := httptest.NewRecorder()
		productsHandler.Get()(response, request)

		// Then
		expectedCode := http.StatusBadRequest
		expectedBody := `{"message":"invalid id", "status":"Bad Request"}`
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		mockRepo.AssertExpectations(t)
	})

	// Case 4: The repository returns an error.
	t.Run("The repository returns an error", func(t *testing.T) {
		// Given
		mockRepo := repository.NewProductsMapMock()
		mockRepo.On("SearchProducts", mock.Anything).Return(make(map[int]internal.Product), internal.ErrRepositoryInternalError)
		productsHandler := handler.NewProductsDefault(mockRepo)

		// When
		request := httptest.NewRequest("GET", "/products?id=1", nil)
		response := httptest.NewRecorder()
		productsHandler.Get()(response, request)

		// Then
		expectedCode := http.StatusInternalServerError
		expectedBody := `{"message":"internal error", "status":"Internal Server Error"}`
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		mockRepo.AssertExpectations(t)
	})

	// Case 5: The query is nil
	t.Run("The query is nil", func(t *testing.T) {
		// Given
		mockRepo := repository.NewProductsMapMock()
		mockRepo.On("SearchProducts", mock.Anything).Return(map[int]internal.Product{}, nil)
		productsHandler := handler.NewProductsDefault(mockRepo)

		// When
		request := httptest.NewRequest("GET", "/products", nil)
		response := httptest.NewRecorder()
		productsHandler.Get()(response, request)

		// Then
		expectedCode := http.StatusOK
		expectedBody := `{"message":"success","data":{}}`
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		mockRepo.AssertExpectations(t)
	})

	// Case 6: The query's int is nagative
	t.Run("The query's int is negative and the db has records", func(t *testing.T) {
		// Given
		mockRepo := repository.NewProductsMapMock()
		mockRepo.On("SearchProducts", mock.Anything).Return(map[int]internal.Product{
			1: {
				Id: 1,
				ProductAttributes: internal.ProductAttributes{
					Description: "Product 1",
					Price:       10.0,
					SellerId:    1,
				},
			},
		}, nil)
		productsHandler := handler.NewProductsDefault(mockRepo)

		// When
		request := httptest.NewRequest("GET", "/products?id=-1", nil)
		response := httptest.NewRecorder()
		productsHandler.Get()(response, request)

		// Then
		expectedCode := http.StatusOK
		expectedBody := `{"message":"success","data":{"1":{"id":1,"description":"Product 1","price":10.0,"seller_id":1}}}`
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		mockRepo.AssertExpectations(t)

	})
}
