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

func (t *transactionRepo) UpdateStatTransaction(transactionId string, transaction *model.Transaction) error {

	if t.db.Model(transaction).Where("id = ?", transactionId).Updates(transaction).RowsAffected == 0 {
		return errors.New("QUERY_NOT_AFFECTED")
	}

	return nil
}

func (t *transactionRepo) ConfirmTransaction(transaction *model.Transaction) error {
	var confirmTransaction *model.Transaction
	err := t.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&confirmTransaction).Error; err != nil {
			return err
		}

		return nil
	})

	return err

}
