package params

import (
	"errors"
	"final-project/server/model"
	"time"

	"github.com/go-playground/validator/v10"
)

type InquireTransactions struct {
	ProductId   string `json:"product_id"`
	Quantity    int    `json:"quantity"`
	Destination string `json:"destination"`
	Weight      int    `json:"weight"`
	TotalPrice  int    `json:"total_price"`
	Courier     string `json:"courier"`
}

type ConfirmTransaction struct {
	ProductId   string         `json:"product_id"`
	Quantity    int            `json:"quantity"`
	Destination string         `json:"destination"`
	Weight      int            `json:"weight"`
	Courier     ConfirmCourier `json:"courier"`
}

type ConfirmCourier struct {
	Code    string `json:"code"`
	Service string `json:"service"`
}

func (t *ConfirmTransaction) ParseToModel() *model.Transaction {
	return &model.Transaction{
		ProductId:      t.ProductId,
		Quantity:       t.Quantity,
		Destination:    t.Destination,
		Weight:         t.Weight,
		CourierCode:    t.Courier.Code,
		CourierService: t.Courier.Service,
		Status:         "WAITING",
		BaseModel: model.BaseModel{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
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

func ValidateConfirmTransaction(u interface{}) ([]string, error) {
	err := validator.New().Struct(u)

	if err != nil {
		var errString []string
		for _, err := range err.(validator.ValidationErrors) {
			errString = append(errString, err.Field()+" is "+err.Tag())
		}

		return errString, err
	}

	return nil, nil
}
