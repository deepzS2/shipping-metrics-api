package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/deepzS2/shipping-metrics-api/internal/domain"
)

type freteRapidoService struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewFreteRapidoService(url string) domain.FreteRapidoService {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	return &freteRapidoService{
		url,
		client,
	}
}

func (s *freteRapidoService) SimulateQuote(quoteRequest *domain.QuoteSimulationRequest) (*domain.QuoteSimulationResponse, error) {
	url := fmt.Sprintf("%s/quote/simulate", s.BaseURL)

	jsonBody, err := json.Marshal(quoteRequest)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBody))

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	if err != nil {
		return nil, err
	}

	response, err := s.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer request.Body.Close()

	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		return nil, fmt.Errorf("received non-OK HTTP status: %s, body: %s", response.Status, string(body))
	}

	var jsonResponse domain.QuoteSimulationResponse
	if err := json.NewDecoder(response.Body).Decode(&jsonResponse); err != nil {
		return nil, err
	}

	return &jsonResponse, nil
}
