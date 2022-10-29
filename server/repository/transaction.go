package repository

import (
	"errors"
	"final-project/server/model"

	"gorm.io/gorm"
)

type transactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) *transactionRepo {
	return &transactionRepo{
		db: db,
	}
}

func (t *transactionRepo) GetDetailProduct(productId string) (*model.Product, error) {
	var product model.Product

	err := t.db.Where("id = ?", productId).First(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (t *transactionRepo) ConfirmTransaction(transaction *model.Transaction) error {
	return t.db.Create(transaction).Error
}

func (t *transactionRepo) UpdateStatTransaction(transactionId string, transaction *model.Transaction) error {

	if t.db.Model(transaction).Where("id = ?", transactionId).Updates(transaction).RowsAffected == 0 {
		return errors.New("QUERY_NOT_AFFECTED")
	}

	return nil
}
