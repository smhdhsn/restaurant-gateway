package response

import (
	"net/http"
)

// NewStatusOK is responsible for creating a 'ok' response.
func NewStatusOK(data any) (int, dataResp) {
	resp := dataResp{
		Status: http.StatusOK,
		Data:   data,
	}

	return resp.Status, resp
}

// NewStatusCreated is responsible for creating a 'created' response.
func NewStatusCreated(data any) (int, dataResp) {
	resp := dataResp{
		Status: http.StatusCreated,
		Data:   data,
	}

	return resp.Status, resp
}

// NewStatusBadRequest is responsible for creating a 'bad request' response.
func NewStatusBadRequest(msg string) (int, messageResp) {
	resp := messageResp{
		Status:  http.StatusBadRequest,
		Message: msg,
	}

	return resp.Status, resp
}

// NewStatusBadRequest is responsible for creating a 'not found' response.
func NewStatusNotFound() (int, messageResp) {
	resp := messageResp{
		Status:  http.StatusNotFound,
		Message: http.StatusText(http.StatusNotFound),
	}

	return resp.Status, resp
}

// StatusUnprocessableEntity is responsible for creating a 'unprocessable entity' response.
func NewStatusUnprocessableEntity(data any) (int, dataResp) {
	resp := dataResp{
		Status: http.StatusUnprocessableEntity,
		Data:   data,
	}

	return resp.Status, resp
}

// NewStatusNotImplemented is responsible for creating a 'not implemented' response.
func NewStatusNotImplemented() (int, messageResp) {
	resp := messageResp{
		Status:  http.StatusNotImplemented,
		Message: http.StatusText(http.StatusNotImplemented),
	}

	return resp.Status, resp
}

// NewStatusInternalServerError is responsible for creating a 'internal server error' response.
func NewStatusInternalServerError() (int, messageResp) {
	resp := messageResp{
		Status:  http.StatusInternalServerError,
		Message: http.StatusText(http.StatusInternalServerError),
	}

	return resp.Status, resp
}
