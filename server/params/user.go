package params

import (
	"errors"
	"final-project/server/model"

	"github.com/go-playground/validator/v10"
)

type UserRegister struct {
	Fullname string `json:"fullname" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role"`
}

func Validate(u interface{}) error {
	err := validator.New().Struct(u)

	if err != nil {
		panic(err)

		validateErrors := err.(validator.ValidationErrors)
		errString := ""
		for _, err := range validateErrors {
			errString += err.Field() + " is " + err.Tag()
		}
		return errors.New(errString)
	}

	return nil
}

func (u *UserRegister) ParseToModel() *model.User {
	return &model.User{
		Fullname: u.Fullname,
		Email:    u.Email,
		Password: u.Password,
	}
}
