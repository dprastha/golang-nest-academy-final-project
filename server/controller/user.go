package controller

import (
	"final-project/server/params"
	"final-project/server/service"
	"final-project/server/view"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (u *UserHandler) Register(ctx *gin.Context) {
	var req params.UserRegister
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		resp := view.ErrorResponse("Invalid Request", "BAD_REQUEST", http.StatusBadRequest)
		WriteJsonResponse(ctx, resp)
		return
	}

	err = params.Validate(req)
	if err != nil {
		resp := view.ErrorResponse("Invalid Request", "BAD_REQUEST", http.StatusBadRequest)
		WriteJsonResponse(ctx, resp)
		return
	}

	resp := u.service.Register(&req)
	WriteJsonResponse(ctx, resp)
}

func (u *UserHandler) Login(ctx *gin.Context) {
	var req params.UserLogin
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		resp := view.ErrorResponse("Invalid Request", "BAD_REQUEST", http.StatusBadRequest)
		WriteJsonResponse(ctx, resp)
		return
	}

	resp := u.service.Login(&req)
	WriteJsonResponse(ctx, resp)
}
