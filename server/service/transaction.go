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

func (t *TransactionService) InquireTransaction(req *params.InquireTransactions) *view.Response {
	transaction := req

	product, err := t.repo.GetDetailProduct(transaction.ProductId)
	if err != nil {
		return view.ErrorResponse("GET_DETAIL_PRODUCT_FAIL", "NOT_FOUND", http.StatusNotFound)
	}

	if transaction.Quantity > product.Stock {
		return view.ErrorResponse("INQUIRY_TRANSACTION_FAIL", "UNPROCESSABLE_ENTITY", http.StatusUnprocessableEntity)
	}

	payload, err := view.GetInquireDetailsPayload(transaction, product, t.rajaongkirAdaptor)
	if err != nil {
		return view.ErrorResponse("GET_DETAIL_COST_COURIER_FAIL", "NOT_FOUND", http.StatusNotFound)
	}

	return view.SuccessResponse("INQUIRY_TRANSACTION_SUCCESS", payload, http.StatusCreated)
}

func (t *TransactionService) ConfirmTransaction(req *params.ProductReq) *view.Response {
	transaction := req.ParseToModel()

	//TODO::Get UserId, and CityId

	//TODO::Get Quantity Product and Update Quantity Product

	return view.SuccessResponse("CONFIRM_TRANSACTION_SUCCESS", transaction, http.StatusCreated)
}

func (t *TransactionService) UpdateStatTransaction(transactionId string, req *params.UpdateStatTransaction) *view.Response {
	status := req.ParseToModel()

	err := t.repo.UpdateStatTransaction(transactionId, status)
	if err != nil {
		return view.ErrorResponse("UPDATE_STATUS_TRANSACTION_FAIL", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return view.SuccessResponse("UPDATE_STATUS_TRANSACTION_SUCCESS", nil, http.StatusAccepted)
}
