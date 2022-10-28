package repository

import "final-project/server/model"

type UserRepo interface {
	Register(user *model.User) error
	FindUserByEmail(email string) (*model.User, error)
}

type ProductRepo interface {
	GetAllProductsWithPagi(pagination model.Pagination) (*[]model.Product, error)
	GetAllProducts() (*[]model.Product, error)
	GetProductById(productId string) (*model.Product, error)
	CreateProduct(product *model.Product) error
	UpdateProduct(productId string, product *model.Product) error
	DeleteProduct(productId string) error
}
