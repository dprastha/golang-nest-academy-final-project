package model

type Product struct {
	BaseModel
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Weight      int32  `json:"weight"`
	Price       int32  `json:"price"`
	Stock       int32  `json:"stock"`
	ImageUrl    string `json:"image_url"`
}
