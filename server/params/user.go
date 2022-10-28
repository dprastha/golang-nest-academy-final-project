package params

import (
	"errors"
	"final-project/server/model"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Fullname   string `json:"fullname" validate:"required,min=3,max=50"`
	Gender     string `json:"gender" validate:"required"`
	Contact    string `json:"contact" validate:"required"`
	Street     string `json:"street" validate:"required"`
	CityId     string `json:"city_id" validate:"required"`
	ProvinceId string `json:"province_id" validate:"required"`
}

type UserRegister struct {
	Fullname string `json:"fullname" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role"`
}

type UserLogin struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
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
	user := &model.User{
		Fullname:   u.Fullname,
		Gender:     u.Gender,
		Contact:    u.Contact,
		Street:     u.Street,
		CityId:     u.CityId,
		ProvinceId: u.ProvinceId,
	}

	return user, nil
}
