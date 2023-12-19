package handler

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

const (
	STATUS_OK         = `OK`
	STATUS_NOTFOUND   = `NOT FOUND`
	STATUS_BADREQUEST = `BAD REQUEST`
)

func ValidateStruct(myStruct interface{}) error {
	validate := validator.New()
	err := validate.Struct(myStruct)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make([]string, len(validationErrors))
		for i, err := range validationErrors {
			errorMessages[i] = fmt.Sprintf("Error when validation %s", err.Field())
		}
		msg := errorMessages[0]
		return errors.New(msg)
	}
	return nil
}
