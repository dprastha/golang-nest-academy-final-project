package repository

import "final-project/server/model"

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
