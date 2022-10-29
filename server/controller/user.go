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

	errString, err := params.Validate(req)
	if err != nil {
		resp := view.ErrorValidationUserResponse("Invalid Request", gin.H{"message": errString}, http.StatusBadRequest)
		WriteErrorJsonResponse(ctx, resp)
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
		WriteErrorJsonResponse(ctx, response)
		return
	}

	errString, err := params.ValidateRequestUser(&body)
	if err != nil {
		payload := view.ErrorValidationUserResponse("CREATED_USER_FAIL", gin.H{"message": errString}, http.StatusUnprocessableEntity)
		WriteErrorJsonResponse(ctx, payload)
		return
	}

	payload := u.service.CreateUser(&body)
	WriteJsonResponse(ctx, payload)
}

func (u *UserHandler) AllUsers(ctx *gin.Context) {
	pageStr := ctx.Request.URL.Query().Get("page")
	limitStr := ctx.Request.URL.Query().Get("limit")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	payload := u.service.GetUsers(page, limit)
	WriteJsonResponse(ctx, payload)
}

func (u *UserHandler) DetailUserByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	if email == "" {
		payload := view.ErrorResponse("GET_USER_BY_EMAIL_FAIL", "NOT_FOUND", http.StatusNotFound)
		WriteErrorJsonResponse(ctx, payload)
		return
	}

	payload := u.service.ShowUserByEmail(email)
	WriteJsonResponse(ctx, payload)
}

func (u *UserHandler) DetailUserById(ctx *gin.Context) {
	id := ctx.GetString("USER_ID")
	if id == "" {
		payload := view.ErrorResponse("GET_USER_PROFILE_FAIL", "UNAUTHORIZED", http.StatusUnauthorized)
		WriteErrorJsonResponse(ctx, payload)
		return
	}

	payload := u.service.ShowUserById(id)
	WriteJsonResponse(ctx, payload)
}

func (u *UserHandler) UpdateUserById(ctx *gin.Context) {
	id := ctx.GetString("USER_ID")
	if id == "" {
		payload := view.ErrorResponse("UPDATE_USER_FAIL", "UNAUTHORIZED", http.StatusUnauthorized)
		WriteErrorJsonResponse(ctx, payload)
		return
	}

	var body params.UpdateUser
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		log.Printf("Error when binding params in update user controller %v\n", err)
		payload := view.ErrorResponse("UPDATE_USER_FAIL", "BAD_REQUEST", http.StatusBadRequest)
		WriteErrorJsonResponse(ctx, payload)
		return
	}

	errString, err := params.ValidateRequestUser(&body)
	if err != nil {
		payload := view.ErrorValidationUserResponse("UPDATE_USER_FAIL", gin.H{"message": errString}, http.StatusUnprocessableEntity)
		WriteErrorJsonResponse(ctx, payload)
		return
	}

	payload := u.service.UpdateUser(&body, id)
	WriteJsonResponse(ctx, payload)
}
