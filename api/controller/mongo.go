package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"main/consumer"
	"main/db"
	"main/model"
	"main/producer"
	"main/request"
	"main/response"
	"net/http"
	"time"
)

type InvoiceController struct {
	Producer   producer.KafkaProducer
	Consumer   consumer.KafkaConsumer
	Collection db.InvoiceCollection
}

func stringToObjectID(id string) (primitive.ObjectID, error) {
	if len(id) == 36 {
		return primitive.ObjectIDFromHex(id[10:34])
	} else if len(id) == 24 {
		return primitive.ObjectIDFromHex(id)
	}
	return primitive.ObjectID{}, errors.New("invalid id")
}

func (receiver InvoiceController) CreateMongo(ctx *gin.Context) {
	var req request.InvoiceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Error: err.Error()})
		return
	}

	invoice := model.Invoice{
		Issued:       time.Now(),
		InvoiceType:  req.InvoiceType,
		FkCustomer:   req.FkCustomer,
		PurchaseList: req.PurchaseList,
		TotalSum:     req.TotalSum,
	}

	id, err := receiver.Collection.Create(invoice)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]string{"id": id.Hex()})
}

func (receiver InvoiceController) GetById(ctx *gin.Context) {
	var req request.IDRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Error: err.Error()})
		return
	}

	id, err := stringToObjectID(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: err.Error()})
		return
	}

	invoice, err := receiver.Collection.GetById(id)
	if errors.Is(err, mongo.ErrNoDocuments) {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{Error: "document with id=" + id.Hex() + " not found"})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, invoice)
}

func (receiver InvoiceController) DeleteById(ctx *gin.Context) {
	hId := ctx.Param("id")
	if hId == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Error: "id not specified"})
		return
	}

	id, err := stringToObjectID(hId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: err.Error()})
		return
	}

	err = receiver.Collection.DeleteById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (receiver InvoiceController) GetAll(ctx *gin.Context) {
	inv, err := receiver.Collection.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, inv)
}

func (receiver InvoiceController) Update(ctx *gin.Context) {
	var invoice model.Invoice
	if err := ctx.ShouldBindJSON(&invoice); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Error: err.Error()})
		return
	}

	count, err := receiver.Collection.Replace(invoice)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, map[string]int{"modified": count})
}
