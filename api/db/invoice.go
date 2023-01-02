package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"main/model"
	"main/util"
)

type InvoiceCollection struct {
	Collection *mongo.Collection
	Context    context.Context
}

func (receiver InvoiceCollection) Create(invoice model.Invoice) (string, error) {
	id, err := receiver.Collection.InsertOne(receiver.Context, invoice)
	return util.IDToString(id), err
}

func (receiver InvoiceCollection) PrettyPrint(invoices []model.Invoice) {
	for key, value := range invoices {
		fmt.Println("Index: " + fmt.Sprintf("%d", key))
		fmt.Println(value)
		fmt.Println()
	}
}
