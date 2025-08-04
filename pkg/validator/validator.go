package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func New() *validator.Validate {
	return validator.New()
}

func FormatValidationErrors(err error) string {
	if errs, ok := err.(validator.ValidationErrors); ok {
		var errorMsgs []string
		for _, e := range errs {
			errorMsgs = append(errorMsgs, fmt.Sprintf("field '%s' failed on the '%s' tag", e.Field(), e.Tag()))
		}
		return strings.Join(errorMsgs, "; ")
	}

	return "Invalid request body"
}
