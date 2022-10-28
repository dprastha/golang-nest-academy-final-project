package repository

import (
	"errors"
	"final-project/server/model"

	"gorm.io/gorm"
)

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) *productRepo {
	return &productRepo{
		db: db,
	}
}

func (p *productRepo) GetAllProductsWithPagi(pagination model.Pagination) (*[]model.Product, error) {
	var product []model.Product

	offset := (pagination.Page - 1) * pagination.Limit

	err := p.db.Limit(pagination.Limit).Offset(offset).Find(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *productRepo) GetAllProducts() (*[]model.Product, error) {
	var product []model.Product

	err := p.db.Find(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *productRepo) GetProductById(productId string) (*model.Product, error) {
	var product model.Product

	err := p.db.Where("id = ?", productId).First(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *productRepo) CreateProduct(product *model.Product) error {
	return p.db.Create(product).Error
}

func (p *productRepo) UpdateProduct(productId string, product *model.Product) error {

	if p.db.Model(product).Where("id = ?", productId).Updates(product).RowsAffected == 0 {
		return errors.New("QUERY_NOT_AFFECTED")
	}

	return nil
}

func (p *productRepo) DeleteProduct(productId string) error {
	var product model.Product

	if p.db.Where("id = ?", productId).Delete(&product).RowsAffected == 0 {
		return errors.New("QUERY_NOT_AFFECTED")
	}

	return nil
}
