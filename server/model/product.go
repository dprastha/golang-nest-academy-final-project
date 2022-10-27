package model

type Product struct {
	BaseModel
	Name     string `json:"name"`
	Category string `json:"category"`
	Desc     string `json:"desc"`
	Weight   int32  `json:"weight"`
	Price    int32  `json:"price"`
	Stock    int32  `json:"stock"`
	ImageUrl string `json:"image_url"`
}
