package repository

import (
	"final-project/server/model"
	"fmt"

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
		fmt.Println("error here 1")
		return nil, err
	}

	return &user, nil
}

func (u *userRepo) GetUsers(page int, limit int) (*[]model.User, error) {
	var users []model.User
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 25
	}

	offset := (page - 1) * limit
	err := u.db.Offset(offset).Limit(limit).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return &users, nil
}

func (u *userRepo) DetailUserById(id string) (*model.User, error) {
	//TODO : if JWT include id user
	return nil, nil
}
