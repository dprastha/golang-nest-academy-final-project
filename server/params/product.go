package params

import (
	"final-project/server/model"
	"time"
)

type ProductReq struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Desc     string `json:"desc"`
	Weight   int    `json:"weight"`
	Price    int    `json:"price"`
	Stock    int    `json:"stock"`
	ImgUrl   string `json:"img_url"`
}

func (p *ProductReq) ParseToModel() *model.Product {
	return &model.Product{
		Name:     p.Name,
		Category: p.Category,
		Desc:     p.Desc,
		Weight:   p.Weight,
		Price:    p.Price,
		Stock:    p.Stock,
		ImgUrl:   p.ImgUrl,
		BaseModel: model.BaseModel{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}
