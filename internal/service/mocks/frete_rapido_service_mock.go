package mocks

import (
	"github.com/deepzS2/shipping-metrics-api/internal/domain"
	"github.com/stretchr/testify/mock"
)

type FreteRapidoServiceMock struct {
	mock.Mock
}

func (m *FreteRapidoServiceMock) SimulateQuote(qr *domain.QuoteSimulationRequest) (*domain.QuoteSimulationResponse, error) {
	args := m.Called(qr)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.QuoteSimulationResponse), args.Error(1)
}
