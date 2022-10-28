package repository

import "final-project/server/model"

type UserRepo interface {
	// Auth
	Register(user *model.User) error
	FindUserByEmail(email string) (*model.User, error)

	// Users
	//GetUsers(page int, limit int) (*[]model.User, error)
	//DetailUserById(id int32) (*model.User, error)
	//DetailUserByEmail(email string) (*model.User, error)
	//EditUser(id int32) (*model.User, error)
}

type ProductRepo interface {
	GetAllProductsWithPagi(pagination model.Pagination) (*[]model.Product, error)
	GetAllProducts() (*[]model.Product, error)
	GetProductById(productId string) (*model.Product, error)
	CreateProduct(product *model.Product) error
	UpdateProduct(productId string, product *model.Product) error
	DeleteProduct(productId string) error
}
