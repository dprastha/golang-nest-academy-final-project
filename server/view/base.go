package view

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	Payload     interface{} `json:"payload,omitempty"`
	Query       interface{} `json:"query,omitempty"`
	Error       string      `json:"error,omitempty"`
	AddInfo     interface{} `json:"additional_info,omitempty"`
	GeneralInfo interface{} `json:"general_info,omitempty"`
}

var ErrMap = map[string]interface{}{
	"BAD_REQUEST":          gin.H{"message": "invalid request payload"},
	"FORBIDDEN_ACCESS":     gin.H{"message": "you dont have access for this resources"},
	"UNAUTHORIZED":         gin.H{"message": "you need to login for access this resources"},
	"NOT_FOUND":            gin.H{"message": "data not found in this resources"},
	"UNPROCESSABLE_ENTITY": gin.H{"message": "stock prodcut not enough"},
}

func SuccessRequest(msg string) *Response {
	return &Response{
		Status:      http.StatusOK,
		Message:     msg,
		GeneralInfo: "Golang-4-Shop",
	}
}

func ErrorRequest(msg string, err string) *Response {
	return &Response{
		Status:      http.StatusBadRequest,
		Message:     msg,
		Error:       err,
		AddInfo:     ErrMap[err],
		GeneralInfo: "Golang-4-Shop",
	}
}
