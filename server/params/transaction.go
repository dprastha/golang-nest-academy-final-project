package params

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
