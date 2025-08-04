package domain

import (
	"context"
)

// --- Quote Models ---

type QuoteInput struct {
	Recipient struct {
		Address struct {
			Zipcode string `json:"zipcode" validate:"required,len=8,numeric"`
		} `json:"address" validate:"required"`
	} `json:"recipient" validate:"required"`
	Volumes []QuoteInputVolume `json:"volumes" validate:"required,dive"`
}

type QuoteInputVolume struct {
	Category      int     `json:"category" validate:"required"`
	Amount        int     `json:"amount" validate:"required,gt=0"`
	UnitaryWeight float64 `json:"unitary_weight" validate:"required,gt=0"`
	Price         float64 `json:"price" validate:"required,gt=0"`
	SKU           string  `json:"sku"`
	Height        float64 `json:"height" validate:"required,gt=0"`
	Width         float64 `json:"width" validate:"required,gt=0"`
	Length        float64 `json:"length" validate:"required,gt=0"`
}

type QuoteOutput struct {
	Carrier []QuoteOutputCarrier `json:"carrier"`
}

type QuoteOutputCarrier struct {
	Name     string  `json:"name"`
	Service  string  `json:"service"`
	Deadline int     `json:"deadline"` // in days
	Price    float64 `json:"price"`
}

// --- Metrics Models ---

type MetricsOutput struct {
	ResultsByCarrier     map[string]MetricsOutputCarrier `json:"results_by_carrier"`
	CheapestFreight      float64                         `json:"cheapest_freight"`
	MostExpensiveFreight float64                         `json:"most_expensive_freight"`
}

type MetricsOutputCarrier struct {
	Count        int     `json:"count"`
	TotalPrice   float64 `json:"total_price"`
	AveragePrice float64 `json:"average_price"`
}

// --- Database models and repositories ---

type QuoteRepository interface {
	Save(ctx context.Context, quotes []QuoteOutputCarrier) error

	GetMetrics(ctx context.Context, lastN *int) (*MetricsOutput, error)
}
