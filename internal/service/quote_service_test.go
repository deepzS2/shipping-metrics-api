package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/deepzS2/shipping-metrics-api/internal/domain"
	"github.com/deepzS2/shipping-metrics-api/internal/service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestQuoteService_CreateQuote(t *testing.T) {
	quoteRepositoryMock := new(mocks.QuoteRepositoryMock)
	freteRapidoServiceMock := new(mocks.FreteRapidoServiceMock)

	quoteService := NewQuoteService(freteRapidoServiceMock, quoteRepositoryMock)

	input := domain.QuoteInput{
		Recipient: struct {
			Address struct {
				Zipcode string `json:"zipcode" validate:"required,len=8,numeric"`
			} `json:"address" validate:"required"`
		}{
			Address: struct {
				Zipcode string `json:"zipcode" validate:"required,len=8,numeric"`
			}{Zipcode: "01311000"},
		},
		Volumes: []domain.QuoteInputVolume{
			{Category: 7, Amount: 1, UnitaryWeight: 5, Price: 349, SKU: "abc-teste-123", Height: 0.2, Width: 0.2, Length: 0.2},
		},
	}

	t.Run("Success", func(t *testing.T) {
		mockResponse := &domain.QuoteSimulationResponse{
			Dispatchers: []domain.ResponseDispatcher{
				{
					Offers: []domain.Offer{
						{
							Carrier:    domain.Carrier{Name: "Correios"},
							Service:    "SEDEX",
							FinalPrice: 25.50,
							DeliveryTime: domain.DeliveryTimeInfo{
								Days:          3,
								EstimatedDate: time.Now().AddDate(0, 0, 3).Format("2006-01-02"),
							},
						},
						{
							Carrier:    domain.Carrier{Name: "Jadlog"},
							Service:    ".Package",
							FinalPrice: 22.75,
							DeliveryTime: domain.DeliveryTimeInfo{
								Days:          2,
								EstimatedDate: time.Now().AddDate(0, 0, 2).Format("2006-01-02"),
							},
						},
					},
				},
			},
		}
		freteRapidoServiceMock.On("SimulateQuote", mock.Anything).Return(mockResponse, nil).Once()
		quoteRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("[]domain.QuoteOutputCarrier")).Return(nil).Once()

		output, err := quoteService.CreateQuote(context.Background(), input)

		assert.NoError(t, err)
		assert.NotNil(t, output)
		assert.Len(t, output.Carrier, 2)
		quoteRepositoryMock.AssertExpectations(t)
		freteRapidoServiceMock.AssertExpectations(t)
	})

	t.Run("Frete Rapido service error", func(t *testing.T) {
		expectedErr := errors.New("API error")
		freteRapidoServiceMock.On("SimulateQuote", mock.Anything).Return(nil, expectedErr).Once()

		output, err := quoteService.CreateQuote(context.Background(), input)

		assert.Error(t, err)
		assert.Nil(t, output)
		assert.Contains(t, err.Error(), "frete rapido error")
		freteRapidoServiceMock.AssertExpectations(t)
	})

	t.Run("Repository save error", func(t *testing.T) {
		mockResponse := &domain.QuoteSimulationResponse{
			Dispatchers: []domain.ResponseDispatcher{
				{
					Offers: []domain.Offer{
						{Carrier: domain.Carrier{Name: "Correios"}, Service: "SEDEX", FinalPrice: 25.50},
					},
				},
			},
		}
		freteRapidoServiceMock.On("SimulateQuote", mock.Anything).Return(mockResponse, nil).Once()
		expectedErr := errors.New("database error")
		quoteRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("[]domain.QuoteOutputCarrier")).Return(expectedErr).Once()

		output, err := quoteService.CreateQuote(context.Background(), input)

		assert.Error(t, err)
		assert.Nil(t, output)
		assert.Contains(t, err.Error(), "failed to save quote data")
		quoteRepositoryMock.AssertExpectations(t)
		freteRapidoServiceMock.AssertExpectations(t)
	})
}
