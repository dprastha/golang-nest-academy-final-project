package model

type Item struct {
	Id            int32 `json:"id"`
	TransactionId int32 `json:"transaction_id"`
	ProductId     int32 `json:"product_id"`
	Quantity      int32 `json:"quantity"`
}
