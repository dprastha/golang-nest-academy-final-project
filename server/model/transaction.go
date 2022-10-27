package model

type Transaction struct {
	BaseModel
	UserId      string  `json:"user_id"`
	ProductId   string  `json:"product_id"`
	Origin      string  `json:"origin"`
	Destination string  `json:"destination"`
	Quantity    int32   `json:"quantity"`
	Weight      int32   `json:"weight"`
	TotalPrice  int32   `json:"total_price"`
	Courier     string  `json:"courier"`
	CourierCost float32 `json:"courier_cost"`
	EstArrived  string  `json:"estimation_arrived"`
	Status      string  `json:"status"`
}
