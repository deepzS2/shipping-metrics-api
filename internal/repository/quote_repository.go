package repository

import (
	"context"
	"database/sql"
	"fmt"
	"math"

	"github.com/deepzS2/shipping-metrics-api/internal/domain"
)

type quoteRepository struct {
	database *sql.DB
}

func NewQuoteRepository(database *sql.DB) domain.QuoteRepository {
	return &quoteRepository{database}
}

func (r *quoteRepository) Save(ctx context.Context, quotes []domain.QuoteOutputCarrier) error {
	tx, err := r.database.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	stmtCarrier, err := tx.PrepareContext(ctx, "INSERT INTO carrier_quotes (carrier_name, carrier_service, deadline_days, price) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}

	defer stmtCarrier.Close()

	for _, quote := range quotes {
		_, err := stmtCarrier.ExecContext(ctx, quote.Name, quote.Service, quote.Deadline, quote.Price)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *quoteRepository) GetMetrics(ctx context.Context, lastN *int) (*domain.MetricsOutput, error) {
	query := `
		SELECT
			carrier_name,
			COUNT(*),
			SUM(price),
			AVG(price)
		FROM carrier_quotes
		GROUP BY carrier_name
		%s
	`

	limitClause := ""

	if lastN != nil && *lastN > 0 {
		limitClause = fmt.Sprintf("LIMIT %d", *lastN)
	}

	query = fmt.Sprintf(query, limitClause)

	rows, err := r.database.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	metrics := &domain.MetricsOutput{
		ResultsByCarrier:     make(map[string]domain.MetricsOutputCarrier),
		CheapestFreight:      math.MaxFloat64,
		MostExpensiveFreight: -1.0,
	}

	for rows.Next() {
		var carrierName string
		var metric domain.MetricsOutputCarrier

		if err := rows.Scan(&carrierName, &metric.Count, &metric.TotalPrice, &metric.AveragePrice); err != nil {
			return nil, err
		}

		metrics.ResultsByCarrier[carrierName] = metric
	}

	minMaxQuery := `
		SELECT
			COALESCE(MIN(price), 0),
			COALESCE(MAX(price), 0)
		FROM carrier_quotes
		%s
	`

	minMaxQuery = fmt.Sprintf(minMaxQuery, limitClause)

	var minPrice, maxPrice float64

	err = r.database.QueryRowContext(ctx, minMaxQuery).Scan(&minPrice, &maxPrice)
	if err != nil {
		return nil, err
	}

	if len(metrics.ResultsByCarrier) > 0 {
		metrics.CheapestFreight = minPrice
		metrics.MostExpensiveFreight = maxPrice
	} else {
		metrics.CheapestFreight = 0
		metrics.MostExpensiveFreight = 0
	}

	return metrics, nil
}
