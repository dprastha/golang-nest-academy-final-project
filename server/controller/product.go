package controller

import (
	"final-project/server/params"
	"final-project/server/service"
	"final-project/server/view"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	svc *service.ProductService
}

func NewProductHandler(svc *service.ProductService) *ProductHandler {
	return &ProductHandler{
		svc: svc,
	}
}

func (p *ProductHandler) GetAllProducts(c *gin.Context) {
	resp := p.svc.GetAllProducts()
	WriteJsonResponse(c, resp)
}

func (p *ProductHandler) GetProductById(c *gin.Context) {
	//get productId
	productId, isExist := c.Params.Get("id")

	if !isExist {
		WriteJsonResponse(c, view.ErrorResponse("GET_DETAIL_PRODUCT_FAIL", "BAD_REQUEST", http.StatusBadRequest))
		return
	}

	resp := p.svc.GetProductById(productId)

	WriteJsonResponse(c, resp)
}

func (p *ProductHandler) CreateProduct(c *gin.Context) {
	var req params.ProductReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		WriteJsonResponse(c, view.ErrorResponse("CREATE_PRODUCT_FAIL", "BAD_REQUEST", http.StatusBadRequest))
		return
	}

	resp := p.svc.CreateProduct(&req)

	WriteJsonResponse(c, resp)
}

func (p *ProductHandler) UpdateProduct(c *gin.Context) {
	//get productId
	productId, isExist := c.Params.Get("id")

	if !isExist {
		WriteJsonResponse(c, view.ErrorResponse("UPDATE_PRODUCT_FAIL", "BAD_REQUEST", http.StatusBadRequest))
		return
	}

	//getting and binding data update
	var req params.ProductReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		WriteJsonResponse(c, view.ErrorResponse("UPDATE_PRODUCT_FAIL", "BAD_REQUEST", http.StatusBadRequest))
		return
	}

	resp := p.svc.UpdateProduct(productId, &req)

	WriteJsonResponse(c, resp)
}

func (p *ProductHandler) DeleteProduct(c *gin.Context) {
	productId, isExist := c.Params.Get("id")

	if !isExist {
		WriteJsonResponse(c, view.ErrorResponse("DELETE_PRODUCT_FAIL", "BAD_REQUEST", http.StatusBadRequest))
		return
	}

	resp := p.svc.DeleteProduct(productId)

	WriteJsonResponse(c, resp)
}
