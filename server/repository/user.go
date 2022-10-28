package repository

import (
	"final-project/server/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	// CreateUser Berfungsi untuk create user
	CreateUser() (*model.User, error)

	// GetAllUsers Akan menampilkan seluruh data users
	GetAllUsers(limit int, page int) (*[]model.User, error)

	// DetailUserById Akan menampilkan detail profile dari user
	DetailUserById(id int32) (*model.User, error)

	// DetailUserByEmail Akan menampilkan detail profile dari user berdasarkan email
	DetailUserByEmail(email string) (*model.User, error)

	// EditUser Berfungsi untuk mengubah data user
	EditUser(id int32) (*model.User, error)
}

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
