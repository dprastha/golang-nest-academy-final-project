package service

import (
	"final-project/adaptor"
	"final-project/server/params"
	"final-project/server/repository"
	"final-project/server/view"
	"net/http"
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

func (t *TransactionService) UpdateStatTransaction(transactionId string, req *params.UpdateStatTransaction) *view.Response {
	status := req.ParseToModel()

	err := t.repo.UpdateStatTransaction(transactionId, status)
	if err != nil {
		return view.ErrorResponse("UPDATE_STATUS_TRANSACTION_FAIL", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return view.SuccessResponse("UPDATE_STATUS_TRANSACTION_SUCCESS", nil, http.StatusAccepted)
}
