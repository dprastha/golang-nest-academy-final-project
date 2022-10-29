package server

import (
	"final-project/server/controller"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router      *gin.Engine
	user        *controller.UserHandler
	product     *controller.ProductHandler
	transaction *controller.TransactionHandler
	middleware  *Middleware
}

func NewRouter(router *gin.Engine, user *controller.UserHandler, product *controller.ProductHandler, transaction *controller.TransactionHandler, middleware *Middleware) *Router {
	return &Router{
		router:      router,
		user:        user,
		product:     product,
		transaction: transaction,
		middleware:  middleware,
	}
}

func (r *Router) Start(port string) {
	// Auth route
	auth := r.router.Group("/auth")
	auth.POST("/register", r.user.Register)
	auth.POST("/login", r.user.Login)

	// User route
	user := r.router.Group("/users")
	user.POST("/", r.user.Create, r.middleware.Auth)
	user.GET("/", r.user.AllUsers, r.middleware.Auth)

	// Product route
	product := r.router.Group("/products")
	product.GET("/", r.product.GetAllProducts, r.middleware.Auth)
	product.GET("/id/:id", r.product.GetProductById, r.middleware.Auth)
	product.POST("/", r.product.CreateProduct, r.middleware.Auth)
	product.PUT("/id/:id", r.product.UpdateProduct, r.middleware.Auth)
	product.DELETE("/id/:id", r.product.DeleteProduct, r.middleware.Auth)

	// Transaction route
	transaction := r.router.Group("/transactions")
	transaction.PUT("/id/:id", r.transaction.UpdateStatTransaction, r.middleware.Auth)

	r.router.Run(port)
}
