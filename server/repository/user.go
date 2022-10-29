package repository

import (
	"final-project/server/model"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) Register(user *model.User) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}

func (u *userRepo) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := u.db.Where("email=?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepo) GetUsers(page int, limit int) (*[]model.User, error) {
	var users []model.User

	offset := (page - 1) * limit
	err := u.db.Offset(offset).Limit(limit).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return &users, nil
}

func (u *userRepo) DetailUserById(id string) (*model.User, error) {
	var user model.User
	err := u.db.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepo) EditUser(user *model.User, id string) error {
	var updateUser *model.User
	err := u.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Find(&updateUser, "id = ?", id).Error; err != nil {
			return err
		}

		updateUser.Fullname = user.Fullname
		updateUser.Gender = user.Gender
		updateUser.Contact = user.Contact
		updateUser.Street = user.Street
		updateUser.CityId = user.CityId
		updateUser.ProvinceId = user.ProvinceId

		if err := tx.Save(&updateUser).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}
