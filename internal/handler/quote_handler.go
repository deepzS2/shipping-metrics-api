package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/deepzS2/shipping-metrics-api/internal/domain"
	"github.com/deepzS2/shipping-metrics-api/internal/service"
	"github.com/deepzS2/shipping-metrics-api/pkg/httputil"
	internalValidator "github.com/deepzS2/shipping-metrics-api/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type QuoteHandler struct {
	service   service.QuoteService
	validator *validator.Validate
}

func NewQuoteHandler(s service.QuoteService) *QuoteHandler {
	return &QuoteHandler{
		service:   s,
		validator: internalValidator.New(),
	}
}

// CreateQuote godoc
//
//	@Summary		Create a shipping quote
//	@Description	Receives shipping details, calculates fictional quotes, and persists the results.
//	@Tags			Quotes
//	@Accept			json
//	@Produce		json
//	@Param			quote	body		domain.QuoteInput	true	"Quote Input"
//	@Success		200		{object}	domain.QuoteOutput
//	@Failure		400		{object}	httputil.ErrorResponse
//	@Failure		500		{object}	httputil.ErrorResponse
//	@Router			/quote [post]
func (h *QuoteHandler) CreateQuote(c *gin.Context) {
	var input domain.QuoteInput

	if err := c.ShouldBindJSON(&input); err != nil {
		httputil.NewErrorResponse(c, http.StatusBadRequest, errors.New("invalid JSON payload"))
		return
	}

	if err := h.validator.Struct(input); err != nil {
		httputil.NewErrorResponse(c, http.StatusBadRequest, errors.New(internalValidator.FormatValidationErrors(err)))
		return
	}

	output, err := h.service.CreateQuote(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, output)
}

// GetMetrics godoc
//
//	@Summary		Get quote metrics
//	@Description	Retrieves and aggregates metrics from previously stored quotes.
//	@Tags			Quotes
//	@Produce		json
//	@Param			last_quotes	query		int	false	"Number of last quotes to include in metrics"
//	@Success		200			{object}	domain.MetricsOutput
//	@Failure		400			{object}	httputil.ErrorResponse
//	@Failure		500			{object}	httputil.ErrorResponse
//	@Router			/metrics [get]
func (h *QuoteHandler) GetMetrics(c *gin.Context) {
	var lastN *int
	lastQuotesParam := c.Query("last_quotes")

	if lastQuotesParam != "" {
		n, err := strconv.Atoi(lastQuotesParam)

		if err != nil || n <= 0 {
			httputil.NewErrorResponse(c, http.StatusBadRequest, errors.New("last_quotes must be a positive integer"))
			return
		}

		lastN = &n
	}

	metrics, err := h.service.GetMetrics(c.Request.Context(), lastN)
	if err != nil {
		httputil.NewErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, metrics)
}
