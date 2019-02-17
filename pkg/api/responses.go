package api

import (
	"encoding/json"
	"net/http"
	"zgo/pkg/domain"
)

// CustomError contains the fields of a simple error (non validation error)
type CustomError struct {
	ErrType string `json:"type"`
	Message string `json:"message"`
}

// ErrorResponse contains the fields needed to render a simple error
type ErrorResponse struct {
	Status string      `json:"status"`
	Error  CustomError `json:"error"`
}

// ValidationErrorResponse contains the fields needed to render a http validation error
type ValidationErrorResponse struct {
	Status string            `json:"status"`
	Errors []ValidationError `json:"errors"`
}

// SendJSON will send a json response back to the client
func SendJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(data)
}

// SendCustomError will send a json with a simple error
func SendCustomError(w http.ResponseWriter, err *domain.Error) {
	SendJSON(w, http.StatusConflict, ErrorResponse{
		Status: "error",
		Error: CustomError{
			ErrType: err.Code(),
			Message: err.Error(),
		},
	})
}

// SendInternalError will send a json with an internal server error
func SendInternalError(w http.ResponseWriter) {
	SendJSON(w, http.StatusInternalServerError, ErrorResponse{
		Status: "error",
		Error: CustomError{
			ErrType: "internal_error",
			Message: "Something went wrong",
		},
	})
}

// SendBadRequest will send a json with an error for a malformed body
func SendBadRequest(w http.ResponseWriter) {
	SendJSON(w, http.StatusBadRequest, ErrorResponse{
		Status: "error",
		Error: CustomError{
			ErrType: "bad_request",
			Message: "Bad request data provided",
		},
	})
}

// SendValidationErrors will send back to the client one or more errors in json format
func SendValidationErrors(w http.ResponseWriter, errors []ValidationError) {
	SendJSON(w, http.StatusUnprocessableEntity, ValidationErrorResponse{
		Status: "error",
		Errors: errors,
	})
}
