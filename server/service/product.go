package service

import (
	"database/sql"
	"final-project/server/model"
	"final-project/server/params"
	"final-project/server/repository"
	"final-project/server/view"
	"fmt"
	"net/http"
)

type ProductService struct {
	repo repository.ProductRepo
}

func NewProductServices(repo repository.ProductRepo) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (p *ProductService) GetAllProducts() *view.Response {
	pagination := model.Pagination{
		Limit: 5,
		Page:  1,
		Total: 0,
	}

	products, err := p.repo.GetAllProductsWithPagi(pagination)
	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrorResponse("GET_ALL_PRODUCTS_FAIL", "NOT_FOUND", http.StatusNotFound)
		}
		view.ErrorResponse("GET_ALL_PRODUCTS_FAIL", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	pagination.Total = len(*products)

	return view.GetAllProductsResponse("GET_ALL_PRODUCTS_SUCCESS", view.GetAllProductsPayload(products), pagination)

}

func (p *ProductService) GetProductById(productId string) *view.Response {
	product, err := p.repo.GetProductById(productId)
	if err != nil {
		return view.ErrorResponse("GET_DETAIL_PRODUCT_FAIL", "NOT_FOUND", http.StatusNotFound)
	}

	return view.GetProductByIdResponse("GET_DETAIL_PRODUCT_SUCCESS", product)
}

func (p *ProductService) CreateProduct(req *params.ProductReq) *view.Response {
	product := req.ParseToModel()

	//get total produt to create product_id
	totalProduct, err := p.repo.GetAllProducts()
	if err != nil {
		return view.ErrorResponse("CREATE_PRODUCT_FAIL", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}
	productId := fmt.Sprintf("KP%v", len(*totalProduct)+1)
	product.BaseModel.Id = productId

	err = p.repo.CreateProduct(product)
	if err != nil {
		return view.ErrorResponse("CREATE_PRODUCT_FAIL", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return view.SuccessResponse("CREATE_PRODUCT_SUCCESS", product, http.StatusCreated)
}

func (p *ProductService) UpdateProduct(productId string, req *params.ProductReq) *view.Response {
	product := req.ParseToModel()

	err := p.repo.UpdateProduct(productId, product)
	if err != nil {
		return view.ErrorResponse("UPDATE_PRODUCT_FAIL", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return view.SuccessResponse("UPDATE_PRODUCT_SUCCESS", product, http.StatusAccepted)
}

func (p *ProductService) DeleteProduct(produtId string) *view.Response {

	err := p.repo.DeleteProduct(produtId)
	if err != nil {
		return view.ErrorResponse("DELETE_PRODUCT_FAIL", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return view.SuccessResponse("DELETE_PRODUCT_SUCCESS", nil, http.StatusNoContent)
}
