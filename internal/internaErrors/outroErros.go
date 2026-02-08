package internaerrors

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func validated(obj interface{}) error {
	validated := validator.New()

	err := validated.Struct(obj)

	errorsValidated := err.(validator.ValidationErrors)

	errorPresent := errorsValidated[0]

	switch errorPresent.Tag() {
	case "required":
		return errors.New("aksjdaksjd")
	}
	return nil
}
