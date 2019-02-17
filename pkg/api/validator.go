package api

import (
	"encoding/json"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

var validate = validator.New()

// ValidationError contains the fields needed to indicate a validation error
// Todo: Provide more usable info for the client, like the rules that are failing
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func validateStruct(dst interface{}) ([]ValidationError, error) {
	err := validate.Struct(dst)

	if err == nil {
		return nil, nil
	}

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return nil, err
	}

	customErrors := []ValidationError{}

	for _, err := range validationErrors {
		customErrors = append(customErrors, ValidationError{
			Field:   err.Field(),
			Message: err.(error).Error(),
		})
	}

	return customErrors, nil
}

// BindJSON will parse a json request body into a struct then validate it if needed
func BindJSON(req *http.Request, dst interface{}) ([]ValidationError, error) {
	if err := json.NewDecoder(req.Body).Decode(dst); err != nil {
		return nil, err
	}

	validationErrors, err := validateStruct(dst)
	if err != nil {
		return nil, err
	}

	return validationErrors, nil
}
