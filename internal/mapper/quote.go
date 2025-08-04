package mapper

import (
	"strconv"

	"github.com/deepzS2/shipping-metrics-api/internal/domain"
)

func MapQuoteToQuoteSimulationRequest(quoteInput *domain.QuoteInput) (*domain.QuoteSimulationRequest, error) {
	const (
		cnpj             = "25438296000158"
		token            = "1d52a9b6b78cf07b08586152459a5c90"
		platformCode     = "5AKVkHqCn"
		dispatherZipcode = 29161376
	)

	recipientZipcode, err := strconv.Atoi(quoteInput.Recipient.Address.Zipcode)
	if err != nil {
		return nil, err
	}

	dispatcher := domain.QuoteSimulationDispatcher{
		RegisteredNumber: cnpj,
		Zipcode:          dispatherZipcode,
		Volumes:          make([]domain.QuoteSimulationVolume, len(quoteInput.Volumes)),
	}

	for index, volume := range quoteInput.Volumes {
		dispatcher.Volumes[index] = domain.QuoteSimulationVolume{
			Category:      strconv.Itoa(volume.Category),
			Amount:        volume.Amount,
			UnitaryWeight: volume.UnitaryWeight,
			UnitaryPrice:  volume.Price,
			Sku:           volume.SKU,
			Height:        volume.Height,
			Width:         volume.Width,
			Length:        volume.Length,
		}
	}

	quoteSimulationRequest := domain.QuoteSimulationRequest{
		Recipient: domain.QuoteSimulationRecipient{
			Zipcode: recipientZipcode,
			Country: "BRA",
			Type:    1,
		},
		Shipper: domain.QuoteSimulationShipper{
			RegisteredNumber: cnpj,
			PlatformCode:     platformCode,
			Token:            token,
		},
		Dispatchers:    []domain.QuoteSimulationDispatcher{dispatcher},
		SimulationType: []int{0},
	}

	return &quoteSimulationRequest, nil
}
