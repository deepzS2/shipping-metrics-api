package mocks

import (
	"context"

	"github.com/deepzS2/shipping-metrics-api/internal/domain"
	"github.com/stretchr/testify/mock"
)

type QuoteServiceMock struct {
	mock.Mock
}

func (m *QuoteServiceMock) CreateQuote(ctx context.Context, input domain.QuoteInput) (*domain.QuoteOutput, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.QuoteOutput), args.Error(1)
}

func (m *QuoteServiceMock) GetMetrics(ctx context.Context, lastN *int) (*domain.MetricsOutput, error) {
	args := m.Called(ctx, lastN)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.MetricsOutput), args.Error(1)
}
