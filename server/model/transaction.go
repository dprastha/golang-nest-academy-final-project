package model

type Transaction struct {
	BaseModel
	UserId         string `json:"user_id"`
	ProductId      string `json:"product_id"`
	Origin         string `json:"origin"`
	Destination    string `json:"destination"`
	Quantity       int    `json:"quantity"`
	Weight         int    `json:"weight"`
	TotalPrice     int    `json:"total_price"`
	CourierCode    string `json:"courier_code"`
	CourierService string `json:"courier_service"`
	CourierCost    int    `json:"courier_cost"`
	CourierEst     string `json:"courier_estimation"`
	EstArrived     string `json:"estimation_arrived"`
	Status         string `json:"status"`
}
