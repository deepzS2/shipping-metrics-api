package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/deepzS2/shipping-metrics-api/internal/domain"
	"github.com/deepzS2/shipping-metrics-api/internal/handler/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	return router
}

func TestQuoteHandler_CreateQuote(t *testing.T) {
	mockService := new(mocks.QuoteServiceMock)
	handler := NewQuoteHandler(mockService)

	router := setupTestRouter()
	router.POST("/quote", handler.CreateQuote)

	t.Run("Success", func(t *testing.T) {
		input := domain.QuoteInput{
			Recipient: struct {
				Address struct {
					Zipcode string `json:"zipcode" validate:"required,len=8,numeric"`
				} `json:"address" validate:"required"`
			}{Address: struct {
				Zipcode string `json:"zipcode" validate:"required,len=8,numeric"`
			}{Zipcode: "01311000"}},
			Volumes: []domain.QuoteInputVolume{{Category: 7, Amount: 1, UnitaryWeight: 5, Price: 349, Height: 0.2, Width: 0.2, Length: 0.2}},
		}
		expectedOutput := &domain.QuoteOutput{Carrier: []domain.QuoteOutputCarrier{{Name: "Test Carrier", Price: 123.45}}}

		mockService.On("CreateQuote", mock.Anything, input).Return(expectedOutput, nil).Once()

		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/quote", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("Invalid JSON", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/quote", bytes.NewBuffer([]byte("{invalid")))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Contains(t, rr.Body.String(), "Invalid JSON payload")
	})

	t.Run("Validation error", func(t *testing.T) {
		input := domain.QuoteInput{} // Empty input to trigger validation
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/quote", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Contains(t, rr.Body.String(), "failed on the 'required' tag")
	})

	t.Run("Service error", func(t *testing.T) {
		input := domain.QuoteInput{
			Recipient: struct {
				Address struct {
					Zipcode string `json:"zipcode" validate:"required,len=8,numeric"`
				} `json:"address" validate:"required"`
			}{Address: struct {
				Zipcode string `json:"zipcode" validate:"required,len=8,numeric"`
			}{Zipcode: "01311000"}},
			Volumes: []domain.QuoteInputVolume{{Category: 7, Amount: 1, UnitaryWeight: 5, Price: 349, Height: 0.2, Width: 0.2, Length: 0.2}},
		}
		expectedErr := errors.New("service failed")

		mockService.On("CreateQuote", mock.Anything, input).Return(nil, expectedErr).Once()

		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/quote", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		mockService.AssertExpectations(t)
	})
}

func TestQuoteHandler_GetMetrics(t *testing.T) {
	mockService := new(mocks.QuoteServiceMock)
	handler := NewQuoteHandler(mockService)

	router := setupTestRouter()
	router.GET("/metrics", handler.GetMetrics)

	t.Run("Success without params", func(t *testing.T) {
		expectedMetrics := &domain.MetricsOutput{}
		mockService.On("GetMetrics", mock.Anything, (*int)(nil)).Return(expectedMetrics, nil).Once()

		req, _ := http.NewRequest(http.MethodGet, "/metrics", nil)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("Success with last_quotes param", func(t *testing.T) {
		lastN := 10
		expectedMetrics := &domain.MetricsOutput{}
		mockService.On("GetMetrics", mock.Anything, &lastN).Return(expectedMetrics, nil).Once()

		req, _ := http.NewRequest(http.MethodGet, "/metrics?last_quotes=10", nil)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("Invalid last_quotes param", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/metrics?last_quotes=abc", nil)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Contains(t, rr.Body.String(), "last_quotes must be a positive integer")
	})
}
