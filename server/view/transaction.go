package view

import (
	"encoding/json"
	"final-project/adaptor"
	"final-project/server/model"
	"final-project/server/params"
)

type InquireTransaction struct {
	Product         GetProductDetail `json:"product"`
	Quantity        int              `json:"quantity"`
	Destination     string           `json:"destination"`
	Weight          int              `json:"weight"`
	TotalPrice      int              `json:"total_price"`
	CourierServices interface{}      `json:"courier_services"`
}

type GetProductDetail struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
	ImgUrl string `json:"img_url"`
}

func GetInquireDetailsPayload(req *params.InquireTransactions, product *model.Product, rajaOngkir *adaptor.RajaOngkirAdaptor) (*InquireTransaction, error) {

	productDetails := GetProductDetail{
		ID:     product.Id,
		Name:   product.Name,
		Price:  product.Price,
		ImgUrl: product.ImgUrl,
	}

	results, err := getCostDetails(req, rajaOngkir)
	if err != nil {
		return nil, err
	}

	return &InquireTransaction{
		Product:         productDetails,
		Quantity:        req.Quantity,
		Destination:     req.Destination,
		Weight:          req.Weight,
		TotalPrice:      req.TotalPrice,
		CourierServices: results,
	}, nil
}

func getCostDetails(req *params.InquireTransactions, rajaOngkir *adaptor.RajaOngkirAdaptor) (interface{}, error) {

	payload := map[string]interface{}{
		"origin":      "10",
		"destination": req.Destination,
		"weight":      req.Weight,
		"courier":     req.Courier,
	}

	cost, err := rajaOngkir.PostCost(payload)
	if err != nil {
		return nil, err
	}

	var jsonData map[string]interface{}

	err = json.Unmarshal(cost, &jsonData)
	if err != nil {
		return nil, err
	}

	resp := jsonData["rajaongkir"].(map[string]interface{})
	result := resp["results"]

	return result, nil
}
