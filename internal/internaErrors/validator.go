package internaerrors

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(obj interface{}) error {

	validate := validator.New()
	err := validate.Struct(obj)

	if err == nil {
		return nil
	}

	validatonErrors := err.(validator.ValidationErrors)
	validatonError := validatonErrors[0]

	switch validatonError.Tag() {
	case "required":
		return errors.New(validatonError.StructField() + " é obrigatório")

	case "min":
		return errors.New(
			validatonError.StructField() +
				" precisa ter no mínimo " +
				validatonError.Param(),
		)

	case "max":
		return errors.New(
			validatonError.StructField() +
				" precisa ter no máximo " +
				validatonError.Param(),
		)

	case "email":
		return errors.New(
			validatonError.StructField() +
				" precisa ser um email válido",
		)
	}
	return nil
}
