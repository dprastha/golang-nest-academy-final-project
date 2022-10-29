package server

import (
	"final-project/helper"
	"final-project/server/controller"
	"final-project/server/service"
	"final-project/server/view"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	userService *service.UserService
}

func NewMiddleware(userService *service.UserService) *Middleware {
	return &Middleware{
		userService: userService,
	}
}

func (m *Middleware) Auth(ctx *gin.Context) {
	bearerToken := ctx.GetHeader("Authorization")
	tokenArr := strings.Split(bearerToken, "Bearer ")

	if len(tokenArr) != 2 {
		ctx.Set("ERROR", "no token")
		resp := view.ErrorResponse("No token provided", "UNAUTHORIZED", http.StatusUnauthorized)
		controller.WriteErrorJsonResponse(ctx, resp)
		return
	}

	// Verify token
	token, err := helper.VerifyToken(tokenArr[1])
	if err != nil {
		ctx.Set("ERROR", err.Error())
		resp := view.ErrorResponse("Invalid token", "UNAUTHORIZED", http.StatusUnauthorized)
		controller.WriteErrorJsonResponse(ctx, resp)
		return
	}

	// send to next handler
	ctx.Set("USER_ID", token.UserId)
	ctx.Set("USER_EMAIL", token.Email)

	// process to another handler
	ctx.Next()
}
