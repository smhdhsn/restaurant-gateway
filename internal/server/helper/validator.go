package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// ParseValidationErr is responsible for parsing validation errors into a proper "key -> value" format.
func ParseValidationErr(err error) map[string]string {
	data := make(map[string]string, 0)
	for _, err := range err.(validator.ValidationErrors) {
		data[err.Field()] = fmt.Sprintf("validation error on '%s' rule", err.Tag())
	}

	return data
}
