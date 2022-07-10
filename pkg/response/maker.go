package response

import (
	"net/http"
	"strings"
)

// NewStatusOK is responsible for creating a "ok" response.
func NewStatusOK(data any) Schema {
	return Schema{
		Status: http.StatusOK,
		Error:  nil,
		Data:   data,
	}
}

// NewStatusCreated is responsible for creating a "created" response.
func NewStatusCreated(data any) Schema {
	return Schema{
		Status: http.StatusOK,
		Error:  nil,
		Data:   data,
	}
}

// NewStatusBadRequest is responsible for creating a "bad request" response.
func NewStatusBadRequest(data any, err error) Schema {
	var errMsg string
	if err == nil {
		errMsg = transformToSnakeCase(http.StatusText(http.StatusBadRequest))
	} else {
		errMsg = err.Error()
	}

	return Schema{
		Status: http.StatusBadRequest,
		Error:  &errMsg,
		Data:   data,
	}
}

// NewStatusNotFound is responsible for creating a "not found" response.
func NewStatusNotFound(err error) Schema {
	var errMsg string
	if err == nil {
		errMsg = transformToSnakeCase(http.StatusText(http.StatusNotFound))
	} else {
		errMsg = err.Error()
	}

	return Schema{
		Status: http.StatusNotFound,
		Error:  &errMsg,
		Data:   nil,
	}
}

// StatusUnprocessableEntity is responsible for creating a "unprocessable entity" response.
func NewStatusUnprocessableEntity(data any, err error) Schema {
	var errMsg string
	if err == nil {
		errMsg = transformToSnakeCase(http.StatusText(http.StatusUnprocessableEntity))
	} else {
		errMsg = err.Error()
	}

	return Schema{
		Status: http.StatusUnprocessableEntity,
		Error:  &errMsg,
		Data:   data,
	}
}

// NewStatusInternalServerError is responsible for creating a "internal server error" response.
func NewStatusInternalServerError(err error) Schema {
	var errMsg string
	if err == nil {
		errMsg = transformToSnakeCase(http.StatusText(http.StatusInternalServerError))
	} else {
		errMsg = err.Error()
	}

	return Schema{
		Status: http.StatusInternalServerError,
		Error:  &errMsg,
		Data:   nil,
	}
}

// transformToSnakeCase transforms a given string to lower snake_case string.
func transformToSnakeCase(simple string) string {
	return strings.ToLower(strings.Replace(simple, " ", "_", -1))
}
