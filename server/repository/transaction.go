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

func (p *transactionRepo) UpdateStatTransaction(transactionId string, transaction *model.Transaction) error {

	if p.db.Model(transaction).Where("id = ?", transactionId).Updates(transaction).RowsAffected == 0 {
		return errors.New("QUERY_NOT_AFFECTED")
	}

	return nil
}
