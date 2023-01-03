package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"main/model"
	"time"
)

type InvoiceCollection struct {
	Collection *mongo.Collection
}

func (receiver InvoiceCollection) Create(invoice model.Invoice) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, err := receiver.Collection.InsertOne(ctx, invoice)
	return id.InsertedID.(string), err
}

func (receiver InvoiceCollection) PrettyPrint(invoices []model.Invoice) {
	for key, value := range invoices {
		fmt.Println("Index: " + fmt.Sprintf("%d", key))
		fmt.Println(value)
		fmt.Println()
	}
}
