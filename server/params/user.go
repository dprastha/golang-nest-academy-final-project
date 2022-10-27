package params

import (
	"errors"
	"final-project/server/model"

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

func (u *UserRegister) ParseToModel() *model.User {
	return &model.User{
		Fullname: u.Fullname,
		Email:    u.Email,
		Password: u.Password,
	}
}
