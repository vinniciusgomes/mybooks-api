package utils

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(obj interface{}) error {
	validate := validator.New()
	err := validate.Struct(obj)
	if err == nil {
		return nil
	}
	validationErrors := err.(validator.ValidationErrors)
	validationError := validationErrors[0]

	field := strings.ToLower(validationError.StructField())
	switch validationError.Tag() {
	case "required":
		return errors.New(field + " is required")
	case "url":
		return errors.New(field + " is invalid URL")
	case "max":
		return errors.New(field + " must be less than or equal to " + validationError.Param())
	case "min":
		return errors.New(field + " must be greater than or equal to " + validationError.Param())
	case "email":
		return errors.New(field + " is invalid email")
	case "uuid4":
		return errors.New(field + " must be a valid UUIDv4")
	case "oneof":
		return errors.New(field + " must be one of: " + validationError.Param())
	default:
		return errors.New("Validation error for field: " + field)
	}
}
