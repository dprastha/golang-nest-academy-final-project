package params

import (
	"final-project/server/model"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Email      string `json:"email" validate:"required,email"`
	Fullname   string `json:"fullname" validate:"required,min=3,max=50"`
	Gender     string `json:"gender" validate:"required,min=1"`
	Contact    string `json:"contact" validate:"required,min=12"`
	Street     string `json:"street" validate:"required,min=3"`
	CityId     string `json:"city_id" validate:"required,min=1"`
	ProvinceId string `json:"province_id" validate:"required,min=1"`
}

type UpdateUser struct {
	Fullname   string `json:"fullname" validate:"required,min=3,max=50"`
	Gender     string `json:"gender" validate:"required,min=1"`
	Contact    string `json:"contact" validate:"required,min=12"`
	Street     string `json:"street" validate:"required,min=3"`
	CityId     string `json:"city_id" validate:"required,min=1"`
	ProvinceId string `json:"province_id" validate:"required,min=1"`
}

type UserRegister struct {
	Fullname string `json:"fullname" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Role     string `json:"role"`
}

type UserLogin struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6"`
}

func Validate(u interface{}) ([]string, error) {
	err := validator.New().Struct(u)

	if err != nil {
		var errString []string
		for _, err := range err.(validator.ValidationErrors) {
			errString = append(errString, err.Field()+" is "+err.Tag())
		}

		return errString, err
	}

	return nil, nil
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
		Email:      u.Email,
		Gender:     u.Gender,
		Contact:    u.Contact,
		Street:     u.Street,
		CityId:     u.CityId,
		ProvinceId: u.ProvinceId,
	}

	return user, nil
}

func (u *UpdateUser) ParseUpdateToModelUser() (*model.User, error) {
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

func ValidateRequestUser(body interface{}) ([]string, error) {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})
	err := validate.Struct(body)
	if err != nil {
		var errString []string
		for _, err := range err.(validator.ValidationErrors) {
			errString = append(errString, err.Field()+" is "+err.Tag())
		}

		return errString, err
	}

	return nil, nil
}
