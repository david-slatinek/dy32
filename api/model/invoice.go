package model

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"time"
)

type Invoice struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Issued       time.Time          `bson:"issued" json:"issued"`
	InvoiceType  InvoiceType        `bson:"invoiceType" json:"invoiceType"`
	FkCustomer   primitive.ObjectID `bson:"fkCustomer" json:"fkCustomer"`
	PurchaseList []Purchase         `bson:"purchase" json:"purchaseList"`
	TotalSum     float32            `bson:"totalSum" json:"totalSum"`
}

func (receiver Invoice) String() string {
	var sb strings.Builder

	sb.WriteString("ID: " + receiver.ID.String())
	sb.WriteString("\nIssued: " + receiver.Issued.String())
	sb.WriteString("\nInvoiceType: " + receiver.InvoiceType.String())
	sb.WriteString("\nFkCustomer: " + receiver.FkCustomer.String())
	sb.WriteString("\nTotalSum: " + fmt.Sprintf("%f", receiver.TotalSum))

	sb.WriteString("\n\nPurchaseList\n")
	for key, value := range receiver.PurchaseList {
		sb.WriteString("Index: " + fmt.Sprintf("%d\n", key))
		sb.WriteString(value.String() + "\n\n")
	}

	return sb.String()
}

func FromJson(data []byte) (Invoice, error) {
	var inv Invoice
	err := json.Unmarshal(data, &inv)
	return inv, err
}
