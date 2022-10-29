package middleware

import (
	"final-project/helper"
	"final-project/server/service"
	"final-project/server/view"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	userService *service.UserService
}

func NewMiddleware(userService *service.UserService) *AuthMiddleware {
	return &AuthMiddleware{
		userService: userService,
	}
}

func (m *AuthMiddleware) Auth(ctx *gin.Context) {
	bearerToken := ctx.GetHeader("Authorization")
	tokenArr := strings.Split(bearerToken, "Bearer")

	if len(tokenArr) != 2 {
		ctx.Set("ERROR", "no token")
		view.ErrorResponse("No token provided", "UNAUTHORIZED", http.StatusUnauthorized)
		return
	}

	// Verify token
	token, err := helper.VerifyToken(tokenArr[1])
	if err != nil {
		ctx.Set("ERROR", err.Error())
		view.ErrorResponse("No token provided", "UNAUTHORIZED", http.StatusUnauthorized)
		return
	}

	// send to next handler
	ctx.Set("USER_ID", token.UserId)
	ctx.Set("USER_EMAIL", token.Email)

	// process to another handler
	ctx.Next()
}
