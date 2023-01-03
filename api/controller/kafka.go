package controller

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"main/model"
	"main/request"
	"main/response"
	"net/http"
	"time"
)

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

func (receiver InvoiceController) Read(ctx *gin.Context) {
	invoice, err := receiver.Consumer.Read()

	if errors.Is(err, context.DeadlineExceeded) {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{Error: "no new data"})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, invoice)
}
