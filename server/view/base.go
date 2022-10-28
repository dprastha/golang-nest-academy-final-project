package view

import (
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
	"BAD_REQUEST":           gin.H{"message": "invalid request payload"},
	"FORBIDDEN_ACCESS":      gin.H{"message": "you dont have access for this resources"},
	"UNAUTHORIZED":          gin.H{"message": "you need to login for access this resources"},
	"NOT_FOUND":             gin.H{"message": "data not found in this resources"},
	"UNPROCESSABLE_ENTITY":  gin.H{"message": "unprocessable entity"},
	"INTERNAL_SERVER_ERROR": gin.H{"message": "internal server error"},
}

func SuccessResponse(msg string, payload interface{}, statusCode int) *Response {
	return &Response{
		Status:      statusCode,
		Message:     msg,
		Payload:     payload,
		GeneralInfo: "Golang-4-Shop",
	}
}

func ErrorResponse(msg string, err string, statusCode int) *Response {
	return &Response{
		Status:      statusCode,
		Message:     msg,
		Error:       err,
		AddInfo:     ErrMap[err],
		GeneralInfo: "Golang-4-Shop",
	}
}
