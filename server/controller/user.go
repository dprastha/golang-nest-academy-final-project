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

func (u *UserHandler) GinRegister(c *gin.Context) {
	var req params.UserRegister
	err := c.ShouldBindJSON(&req)
	if err != nil {
		resp := view.ErrorResponse("Invalid Request", "BAD_REQUEST", http.StatusBadRequest)
		WriteJsonResponseGin(c, resp)
		return
	}

	err = params.Validate(req)
	if err != nil {
		resp := view.ErrorResponse("Invalid Request", "BAD_REQUEST", http.StatusBadRequest)
		WriteJsonResponseGin(c, resp)
		return
	}

	resp := u.service.Register(&req)
	WriteJsonResponseGin(c, resp)
}
