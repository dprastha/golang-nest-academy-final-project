package server

import (
	"final-project/server/controller"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router  *gin.Engine
	user    *controller.UserHandler
	product *controller.ProductHandler
}

func NewRouter(router *gin.Engine, user *controller.UserHandler, product *controller.ProductHandler) *Router {
	return &Router{
		router:  router,
		user:    user,
		product: product,
	}
}

func (r *Router) Start(port string) {
	// Auth route
	auth := r.router.Group("/auth")
	auth.POST("/register", r.user.Register)

	user := r.router.Group("/users")
	user.POST("/", r.user.Create)

	product := r.router.Group("/products")
	product.GET("/", r.product.GetAllProducts)
	product.GET("/id/:id", r.product.GetProductById)
	product.POST("/", r.product.CreateProduct)
	product.PUT("/id/:id", r.product.UpdateProduct)
	product.DELETE("/id/:id", r.product.DeleteProduct)
	r.router.Run(port)
}
