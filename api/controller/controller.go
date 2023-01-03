package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"main/model"
	"main/producer"
	"main/request"
	"main/response"
	"net/http"
	"time"
)

type InvoiceController struct {
	Producer producer.KafkaProducer
}

func (receiver InvoiceController) Create(ctx *gin.Context) {
	var req request.InvoiceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Error: err.Error()})
		return
	}

	invoice := model.Invoice{
		ID:           uuid.NewString(),
		Issued:       time.Now(),
		InvoiceType:  req.InvoiceType,
		FkCustomer:   req.FkCustomer,
		PurchaseList: req.PurchaseList,
		TotalSum:     req.TotalSum,
	}

	err := receiver.Producer.Write(invoice)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, map[string]string{"id": invoice.ID})
}
