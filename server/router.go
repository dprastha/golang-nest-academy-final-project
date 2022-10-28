package server

import (
	"final-project/server/controller"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router      *gin.Engine
	user        *controller.UserHandler
	product     *controller.ProductHandler
	transaction *controller.TranscationHandler
}

func NewRouter(router *gin.Engine, user *controller.UserHandler, product *controller.ProductHandler, transaction *controller.TranscationHandler) *Router {
	return &Router{
		router:      router,
		user:        user,
		product:     product,
		transaction: transaction,
	}
}

func (r *Router) Start(port string) {
	// Auth route
	auth := r.router.Group("/auth")
	auth.POST("/register", r.user.Register)
	auth.POST("/login", r.user.Login)

	product := r.router.Group("/products")
	product.GET("/", r.product.GetAllProducts)
	product.GET("/id/:id", r.product.GetProductById)
	product.POST("/", r.product.CreateProduct)
	product.PUT("/id/:id", r.product.UpdateProduct)
	product.DELETE("/id/:id", r.product.DeleteProduct)
	r.router.Run(port)
}
