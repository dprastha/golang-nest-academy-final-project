package view

import "net/http"

func UserCreated(msg string) *Response {
	return &Response{
		Status:      http.StatusCreated,
		Message:     msg,
		GeneralInfo: "NooBee-Shop",
	}
}
