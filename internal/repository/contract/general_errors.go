package contract

import (
	"errors"
)

// This block holds common errors that might happen within user source repository.
var (
	ErrRecordNotFound = errors.New("record_not_found")
	ErrDuplicateEntry = errors.New("duplicate_entry")
	ErrUncaught       = errors.New("uncought_error")
	ErrEmptyResult    = errors.New("empty_result")
)
