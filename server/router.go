package server

import (
	"final-project/server/controller"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
	user   *controller.UserHandler
}

func NewRouter(router *gin.Engine, user *controller.UserHandler) *Router {
	return &Router{
		router: router,
		user:   user,
	}
}

func (r *Router) Start(port string) {
	// Auth route
	auth := r.router.Group("/auth")
	auth.POST("/register", r.user.GinRegister)

	r.router.Run(port)
}
