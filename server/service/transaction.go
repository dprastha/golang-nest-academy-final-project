package service

import (
	"encoding/json"
	"final-project/adaptor"
	"final-project/server/params"
	"final-project/server/repository"
	"final-project/server/view"
	"net/http"
	"strings"
)

type TransactionService struct {
	repo              repository.TransactionRepo
	userRepo          repository.UserRepo
	productRepo       repository.ProductRepo
	rajaongkirAdaptor *adaptor.RajaOngkirAdaptor
}

func NewTransactionServices(repo repository.TransactionRepo, userRepo repository.UserRepo, productRepo repository.ProductRepo, rajaongkirAdaptor *adaptor.RajaOngkirAdaptor) *TransactionService {
	return &TransactionService{
		repo:              repo,
		userRepo:          userRepo,
		productRepo:       productRepo,
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

func (t *TransactionService) ConfirmTransaction(req *params.ConfirmTransaction, userId string) *view.Response {
	transaction := req.ParseToModel()

	// Get product
	product, err := t.repo.GetDetailProduct(transaction.ProductId)
	if err != nil {
		return view.ErrorResponse("GET_DETAIL_PRODUCT_FAIL", "NOT_FOUND", http.StatusNotFound)
	}

	if transaction.Quantity > product.Stock {
		return view.ErrorResponse("INQUIRY_TRANSACTION_FAIL", "UNPROCESSABLE_ENTITY", http.StatusUnprocessableEntity)
	}

	// Get user
	user, err := t.userRepo.DetailUserById(userId)
	if err != nil {
		return view.ErrorResponse("GET_DETAIL_USER_FAIL", "NOT_FOUND", http.StatusNotFound)
	}

	// Get cost
	getCost, err := GetCost(req, user.CityId, t.rajaongkirAdaptor)
	if err != nil {
		return view.ErrorResponse("GET_DETAIL_COST_COURIER_FAIL", "NOT_FOUND", http.StatusNotFound)
	}

	// Get courier service cost
	var servicesCost []int
	var estimationArrived []string
	for _, courierCost := range *getCost {
		if strings.EqualFold(req.Courier.Service, courierCost.Service) {
			for _, serviceCost := range *courierCost.ServiceCost {
				servicesCost = append(servicesCost, serviceCost.Value)
				estimationArrived = append(estimationArrived, serviceCost.Etd)
			}
		}
	}

	if len(servicesCost) < 1 {
		return view.ErrorResponse("GET_DETAIL_COST_COURIER_FAIL", "NOT_FOUND", http.StatusNotFound)
	}

	// Update quantity product
	product.Stock -= req.Quantity
	err = t.productRepo.UpdateProduct(product.Id, product)
	if err != nil {
		return view.ErrorResponse("UPDATE_PRODUCT_STOCK_FAIL", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	// Create transaction
	transaction.UserId = user.Id
	transaction.ProductId = product.Id
	transaction.CourierCost = servicesCost[0]
	transaction.EstArrived = estimationArrived[0]

	err = t.repo.ConfirmTransaction(transaction)
	if err != nil {
		return view.ErrorResponse("CONFIRM_TRANSACTION_FAIL", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

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

func GetCost(req *params.ConfirmTransaction, cityId string, rajaOngkir *adaptor.RajaOngkirAdaptor) (*[]adaptor.CourierCost, error) {
	payload := map[string]interface{}{
		"origin":      cityId,
		"destination": req.Destination,
		"weight":      req.Weight,
		"courier":     req.Courier,
	}

	cost, err := rajaOngkir.PostCost(payload)
	if err != nil {
		return nil, err
	}

	var jsonData *adaptor.RajaOngkirResponse

	err = json.Unmarshal(cost, &jsonData)
	if err != nil {
		return nil, err
	}

	result := jsonData.Rajaongkir.Result.CourierCost

	return result, nil
}
