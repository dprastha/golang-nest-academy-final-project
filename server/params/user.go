package params

import (
	"errors"
	"final-project/server/model"

	"github.com/go-playground/validator/v10"
)

type User struct {
	FullName   string `validate:"required"`
	Gender     string `validate:"required"`
	Contact    string `validate:"required"`
	Street     string `validate:"required"`
	CityId     string `validate:"required"`
	ProvinceId string `validate:"required"`
}

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

func (u *User) ParseToModelUser() (*model.User, error) {
	//TODO: validate province and city id is available in raja ongkir

	var validate *validator.Validate

	user := &model.User{
		FullName:   u.FullName,
		Gender:     u.Gender,
		Contact:    u.Contact,
		Street:     u.Street,
		CityId:     u.CityId,
		ProvinceId: u.ProvinceId,
	}

	err := validate.Struct(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
