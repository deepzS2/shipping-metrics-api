package domain

import "time"

type QuoteSimulationRequest struct {
	Recipient      QuoteSimulationRecipient    `json:"recipient"`
	Shipper        QuoteSimulationShipper      `json:"shipper"`
	Dispatchers    []QuoteSimulationDispatcher `json:"dispatchers"`
	SimulationType []int                       `json:"simulation_type"`
}

type QuoteSimulationRecipient struct {
	Zipcode int    `json:"zipcode"`
	Country string `json:"country"`
	Type    int    `json:"type"`
}

type QuoteSimulationShipper struct {
	RegisteredNumber string `json:"registered_number"`
	PlatformCode     string `json:"platform_code"`
	Token            string `json:"token"`
}

type QuoteSimulationDispatcher struct {
	RegisteredNumber string                  `json:"registered_number"`
	Zipcode          int                     `json:"zipcode"`
	Volumes          []QuoteSimulationVolume `json:"volumes"`
}

type QuoteSimulationVolume struct {
	Category      string  `json:"category"`
	Amount        int     `json:"amount"`
	UnitaryWeight float64 `json:"unitary_weight"`
	UnitaryPrice  float64 `json:"unitary_price"`
	Sku           string  `json:"sku"`
	Height        float64 `json:"height"`
	Width         float64 `json:"width"`
	Length        float64 `json:"length"`
}

type QuoteSimulationResponse struct {
	Dispatchers []ResponseDispatcher `json:"dispatchers"`
}

type ResponseDispatcher struct {
	ID                         string  `json:"id"`
	Offers                     []Offer `json:"offers"`
	RegisteredNumberDispatcher string  `json:"registered_number_dispatcher"`
	RegisteredNumberShipper    string  `json:"registered_number_shipper"`
	RequestID                  string  `json:"request_id"`
	ZipcodeOrigin              int     `json:"zipcode_origin"`
}

type Offer struct {
	Carrier                     Carrier          `json:"carrier"`
	CarrierOriginalDeliveryTime DeliveryTimeInfo `json:"carrier_original_delivery_time"`
	CostPrice                   float64          `json:"cost_price"`
	DeliveryTime                DeliveryTimeInfo `json:"delivery_time"`
	Esg                         EsgInfo          `json:"esg"`
	Expiration                  time.Time        `json:"expiration"`
	FinalPrice                  float64          `json:"final_price"`
	HomeDelivery                bool             `json:"home_delivery"`
	Modal                       string           `json:"modal"`
	Offer                       int              `json:"offer"`
	OriginalDeliveryTime        DeliveryTimeInfo `json:"original_delivery_time"`
	Service                     string           `json:"service"`
	SimulationType              int              `json:"simulation_type"`
	TableReference              string           `json:"table_reference"`
	Weights                     WeightInfo       `json:"weights"`
}

type Carrier struct {
	CompanyName      string `json:"company_name"`
	Logo             string `json:"logo"`
	Name             string `json:"name"`
	Reference        int    `json:"reference"`
	RegisteredNumber string `json:"registered_number"`
	StateInscription string `json:"state_inscription"`
}

type DeliveryTimeInfo struct {
	Days          int    `json:"days"`
	EstimatedDate string `json:"estimated_date"`
}

type EsgInfo struct {
	Co2EmissionEstimate float64 `json:"co2_emission_estimate"`
}

type WeightInfo struct {
	Cubed float64 `json:"cubed"`
	Real  int     `json:"real"`
	Used  float64 `json:"used"`
}

type FreteRapidoService interface {
	SimulateQuote(qr *QuoteSimulationRequest) (*QuoteSimulationResponse, error)
}
