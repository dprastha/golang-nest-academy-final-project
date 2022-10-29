package params

import (
	"errors"
	"final-project/server/model"

	"github.com/go-playground/validator/v10"
)

type InquireTransactions struct {
	ProductId   string `json:"product_id"`
	Quantity    int    `json:"quantity"`
	Destination string `json:"destination"`
	Weight      int    `json:"weight"`
	TotalPrice  int    `json:"total_price"`
	Courier     string `json:"courier"`
	Status      string `json:"status"`
	EstArrived  string `json:"estimation_arrived"`
}

type UpdateStatTransaction struct {
	Status string `json:"status" validate:"isValidStatus"`
}

func ValidateStatus(u interface{}) error {
	validate := validator.New()
	validate.RegisterValidation("isValidStatus", isValidStatus)
	err := validate.Struct(u)

	if err == nil {
		return nil
	}
	myErr := err.(validator.ValidationErrors)
	errString := ""
	for _, e := range myErr {
		errString += e.Field() + " is " + e.Tag()
	}
	return errors.New(errString)
}

func isValidStatus(fl validator.FieldLevel) bool {
	v, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	elems := []string{"WAITING", "PICKUP", "ON_THE_WAY", "ARRIVED"}

	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func (t *UpdateStatTransaction) ParseToModel() *model.Transaction {
	return &model.Transaction{
		Status: t.Status,
	}
}
