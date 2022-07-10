package validator

import (
	"github.com/go-playground/validator/v10"
)

// v holds an instance of the validator struct.
var v *validator.Validate

// init function will be executed when this package is used.
func init() {
	v = validator.New()
}

// GetInstance gets the instantiated validator instance.
func GetInstance() *validator.Validate {
	return v
}
