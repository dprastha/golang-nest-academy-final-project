package controller

import "final-project/server/service"

type TranscationHandler struct {
	svc *service.TransactionService
}

func NewTranscationHandler(svc *service.TransactionService) *TranscationHandler {
	return &TranscationHandler{
		svc: svc,
	}
}
