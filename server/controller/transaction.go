package controller

import (
	"final-project/server/params"
	"final-project/server/service"
	"final-project/server/view"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	svc *service.TransactionService
}

func NewTranscationHandler(svc *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		svc: svc,
	}
}

func (t *TransactionHandler) UpdateStatTransaction(c *gin.Context) {
	//get productId
	transactionId, isExist := c.Params.Get("id")

	if !isExist {
		WriteJsonResponse(c, view.ErrorResponse("UPDATE_STATUS_TRANSACTION_FAIL", "BAD_REQUEST", http.StatusBadRequest))
		return
	}

	//getting and binding data update
	var req params.UpdateStatTransaction

	err := c.ShouldBindJSON(&req)
	if err != nil {
		WriteJsonResponse(c, view.ErrorResponse("UPDATE_STATUS_TRANSACTION__FAIL", "BAD_REQUEST", http.StatusBadRequest))
		return
	}

	err = params.ValidateStatus(req)
	if err != nil {
		resp := view.ErrorResponse("STATUS_NOT_VALID", "BAD_REQUEST", http.StatusBadRequest)
		WriteJsonResponse(c, resp)
		return
	}

	resp := t.svc.UpdateStatTransaction(transactionId, &req)

	WriteJsonResponse(c, resp)
}
