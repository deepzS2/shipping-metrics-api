package httputil

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorResponse defines the standard JSON structure for API errors.
type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Error      string `json:"error,omitempty"`
}

func NewErrorResponse(c *gin.Context, statusCode int, err error) {
	resp := ErrorResponse{
		StatusCode: statusCode,
		Message:    http.StatusText(statusCode), // e.g., "Bad Request", "Internal Server Error"
		Error:      err.Error(),
	}

	c.AbortWithStatusJSON(statusCode, resp)
}
