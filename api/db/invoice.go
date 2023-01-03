package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"main/model"
	"time"
)

type InvoiceCollection struct {
	Collection *mongo.Collection
}

func StringToID(id string) primitive.ObjectID {
	var result primitive.ObjectID

	if len(id) == 36 {
		result, _ = primitive.ObjectIDFromHex(id[10:34])
	} else if len(id) == 24 {
		result, _ = primitive.ObjectIDFromHex(id)
	}
	return result
}

func (receiver InvoiceCollection) Create(invoice model.Invoice) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, err := receiver.Collection.InsertOne(ctx, invoice)
	return id.InsertedID.(primitive.ObjectID), err
}

func (receiver InvoiceCollection) GetById(id primitive.ObjectID) (model.Invoice, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result model.Invoice
	err := receiver.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	return result, err
}

func (receiver InvoiceCollection) DeleteById(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := receiver.Collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (receiver InvoiceCollection) GetAll() ([]model.Invoice, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cur, err := receiver.Collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	defer cur.Close(ctx)

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var invoices []model.Invoice
	for cur.Next(ctx) {
		var inv model.Invoice
		if err := cur.Decode(&inv); err != nil {
			continue
		}
		invoices = append(invoices, inv)
	}
	return invoices, cur.Err()
}
