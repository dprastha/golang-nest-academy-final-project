package params

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type UserRegister struct {
	Fullname string `validate:"required"`
	Email    string `validate:"required"`
	Password string `validate:"required"`
	Role     string
}

func Validate(u interface{}) error {
	err := validator.New().Struct(u)

	if err != nil {
		panic(err)
	}

	validateErrors := err.(*validator.ValidationErrors)
	errString := ""
	for _, err := range *validateErrors {
		errString += err.Field() + " is " + err.Tag()
	}

	return errors.New(errString)
}
