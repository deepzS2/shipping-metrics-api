package service

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/deepzS2/shipping-metrics-api/internal/domain"
	"github.com/deepzS2/shipping-metrics-api/internal/mapper"
)

type QuoteService interface {
	CreateQuote(ctx context.Context, input domain.QuoteInput) (*domain.QuoteOutput, error)
	GetMetrics(ctx context.Context, lastN *int) (*domain.MetricsOutput, error)
}

type quoteService struct {
	FreteRapido domain.FreteRapidoService
	Repository  domain.QuoteRepository
}

func NewQuoteService(freteRapidoService domain.FreteRapidoService, quoteRepository domain.QuoteRepository) QuoteService {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	return &quoteService{
		freteRapidoService,
		quoteRepository,
	}
}

func (s *quoteService) CreateQuote(ctx context.Context, input domain.QuoteInput) (*domain.QuoteOutput, error) {
	quoteSimulationRequest, err := mapper.MapQuoteToQuoteSimulationRequest(&input)
	if err != nil {
		return nil, fmt.Errorf("failed to map request payload to simulation: %w", err)
	}

	log.Printf("%+v\n", quoteSimulationRequest)

	quoteSimulationResponse, err := s.FreteRapido.SimulateQuote(quoteSimulationRequest)
	if err != nil {
		return nil, fmt.Errorf("frete rapido error: %w", err)
	}

	log.Printf("%+v\n", quoteSimulationResponse)

	var carriers []domain.QuoteOutputCarrier

	for _, offer := range quoteSimulationResponse.Dispatchers[0].Offers {
		carrier := domain.QuoteOutputCarrier{
			Name:     offer.Carrier.Name,
			Service:  offer.Service,
			Price:    offer.FinalPrice,
			Deadline: offer.DeliveryTime.Days,
		}

		carriers = append(carriers, carrier)
	}

	if err := s.Repository.Save(ctx, carriers); err != nil {
		return nil, fmt.Errorf("failed to save quote data: %w", err)
	}

	output := &domain.QuoteOutput{
		Carrier: carriers,
	}

	return output, nil
}

func (s *quoteService) GetMetrics(ctx context.Context, lastN *int) (*domain.MetricsOutput, error) {
	metrics, err := s.Repository.GetMetrics(ctx, lastN)
	if err != nil {
		return nil, fmt.Errorf("failed to get metrics: %w", err)
	}

	return metrics, nil
}
