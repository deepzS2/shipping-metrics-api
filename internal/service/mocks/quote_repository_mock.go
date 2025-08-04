package mocks

import (
	"context"

	"github.com/deepzS2/shipping-metrics-api/internal/domain"
	"github.com/stretchr/testify/mock"
)

type QuoteRepositoryMock struct {
	mock.Mock
}

func (m *QuoteRepositoryMock) Save(ctx context.Context, quotes []domain.QuoteOutputCarrier) error {
	args := m.Called(ctx, quotes)
	return args.Error(0)
}

func (m *QuoteRepositoryMock) GetMetrics(ctx context.Context, lastN *int) (*domain.MetricsOutput, error) {
	args := m.Called(ctx, lastN)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.MetricsOutput), args.Error(1)
}
