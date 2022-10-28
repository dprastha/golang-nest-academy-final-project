package service

import (
	"final-project/adaptor"
	"final-project/server/repository"
)

type TransactionService struct {
	repo              repository.TransactionRepo
	rajaongkirAdaptor *adaptor.RajaOngkirAdaptor
}

func NewTransactionServices(repo repository.TransactionRepo, rajaongkirAdaptor *adaptor.RajaOngkirAdaptor) *TransactionService {
	return &TransactionService{
		repo:              repo,
		rajaongkirAdaptor: rajaongkirAdaptor,
	}
}
