package controller

import (
	"final-project/server/params"
	"final-project/server/service"
	"final-project/server/view"
	"log"
	"net/http"
	"strconv"

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

func (u *UserHandler) Create(ctx *gin.Context) {
	var body params.User
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		log.Printf("Error when binding params in create user controller %v\n", err)
		response := view.ErrorResponse("CREATED_USER_FAIL", "BAD_REQUEST", http.StatusBadRequest)
		WriteJsonResponse(ctx, response)
		return
	}

	//TODO : validate body request

	response := u.service.CreateUser(&body)
	WriteJsonResponse(ctx, response)
}

func (u *UserHandler) AllUsers(ctx *gin.Context) {
	pageStr := ctx.Request.URL.Query().Get("page")
	limitStr := ctx.Request.URL.Query().Get("limit")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	response := u.service.GetUsers(page, limit)
	WriteJsonResponse(ctx, response)
}

func (u *UserHandler) DetailUserByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	if email == "" {
		payload := view.ErrorResponse("GET_USER_BY_EMAIL_FAIL", "NOT_FOUND", http.StatusNotFound)
		WriteErrorJsonResponse(ctx, payload)
	}

	payload := u.service.ShowUserByEmail(email)
	WriteJsonResponse(ctx, payload)
}
