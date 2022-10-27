package view

import (
	"final-project/server/model"
	"net/http"
)

func GetProductByIdResponse(msg string, payload interface{}) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: msg,
		Payload: payload,
	}
}

type GetAllProducts struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Desc     string `json:"desc"`
	Weight   int32  `json:"weight"`
	Price    int32  `json:"price"`
	Stock    int32  `json:"stock"`
	ImgUrl   string `json:"img_url"`
}

func GetAllProductsPayload(products *[]model.Product) *[]GetAllProducts {
	var productsResp []GetAllProducts

	for _, product := range *products {
		productsResp = append(productsResp, GetAllProducts{
			ID:       product.Id,
			Name:     product.Name,
			Category: product.Category,
			Desc:     product.Desc,
			Weight:   product.Weight,
			Price:    product.Price,
			Stock:    product.Stock,
			ImgUrl:   product.ImgUrl,
		})
	}

	return &productsResp
}

func GetAllProductsResponse(msg string, payload interface{}, query interface{}) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: msg,
		Payload: payload,
		Query:   query,
	}
}
