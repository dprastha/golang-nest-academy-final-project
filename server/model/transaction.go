package model

type Transaction struct {
	BaseModel
	UserId      int32  `json:"user_id"`
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	Weight      int32  `json:"weight"`
	Courier     string `json:"courier"`
	Status      string `json:"status"`
}
