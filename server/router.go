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
	user := r.router.Group("/users", r.middleware.Auth)
	user.POST("/", r.user.Create)
	user.GET("/", r.user.AllUsers)
	user.GET("/email/:email", r.user.DetailUserByEmail)

	// Product route
	product := r.router.Group("/products", r.middleware.Auth)
	product.GET("/", r.product.GetAllProducts)
	product.GET("/id/:id", r.product.GetProductById)
	product.POST("/", r.product.CreateProduct)
	product.PUT("/id/:id", r.product.UpdateProduct)
	product.DELETE("/id/:id", r.product.DeleteProduct)

	// Transaction route
	transaction := r.router.Group("/transactions", r.middleware.Auth)
	transaction.PUT("/id/:id", r.transaction.UpdateStatTransaction)

	r.router.Run(port)
}
