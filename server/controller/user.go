package controller

import (
	"final-project/server/params"
	"final-project/server/service"
	"final-project/server/view"
	"log"
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

func (u *UserHandler) Create(c *gin.Context) {
	var body params.User
	err := c.ShouldBindJSON(&body)
	if err != nil {
		log.Printf("Error when binding params in create user controller %v\n", err)
		response := view.ErrorResponse("CREATED_USER_FAIL", "BAD_REQUEST", http.StatusBadRequest)
		WriteJsonResponse(c, response)
		return
	}

	response := u.service.CreateUser(&body)
	WriteJsonResponse(c, response)
}
