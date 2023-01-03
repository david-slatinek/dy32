package model

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Invoice struct {
	ID           string      `bson:"_id,omitempty" json:"id,omitempty"`
	Issued       time.Time   `bson:"issued" json:"issued"`
	InvoiceType  InvoiceType `bson:"invoiceType" json:"invoiceType"`
	FkCustomer   string      `bson:"fkCustomer" json:"fkCustomer"`
	PurchaseList []Purchase  `bson:"purchase" json:"purchaseList"`
	TotalSum     float32     `bson:"totalSum" json:"totalSum"`
}

func (receiver Invoice) String() string {
	var sb strings.Builder

	sb.WriteString("ID: " + receiver.ID)
	sb.WriteString("\nIssued: " + receiver.Issued.String())
	sb.WriteString("\nInvoiceType: " + receiver.InvoiceType.String())
	sb.WriteString("\nFkCustomer: " + receiver.FkCustomer)
	sb.WriteString("\nTotalSum: " + fmt.Sprintf("%f", receiver.TotalSum))

	sb.WriteString("\n\nPurchaseList\n")
	for key, value := range receiver.PurchaseList {
		sb.WriteString("Index: " + fmt.Sprintf("%d\n", key))
		sb.WriteString(value.String() + "\n\n")
	}

	return sb.String()
}

func (receiver Invoice) ToJson() ([]byte, error) {
	return json.Marshal(receiver)
}

func FromJson(data []byte) (Invoice, error) {
	var inv Invoice
	err := json.Unmarshal(data, &inv)
	return inv, err
}
