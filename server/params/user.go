package params

import (
	"final-project/server/model"
	"github.com/go-playground/validator/v10"
)

type User struct {
	FullName   string `validate:"required"`
	Gender     string `validate:"required"`
	Contact    string `validate:"required"`
	Street     string `validate:"required"`
	CityId     int32  `validate:"required"`
	ProvinceId int32  `validate:"required"`
}

func (u *User) ParseToModel() (*model.User, error) {
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
