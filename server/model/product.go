package model

type Product struct {
	BaseModel
	Name     string `json:"name"`
	Category string `json:"category"`
	Desc     string `json:"desc"`
	Weight   int    `json:"weight"`
	Price    int    `json:"price"`
	Stock    int    `json:"stock"`
	ImgUrl   string `json:"img_url"`
}
