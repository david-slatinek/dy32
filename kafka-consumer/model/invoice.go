package model

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Invoice struct {
	ID           string      `json:"id"`
	Issued       string      `json:"issued" `
	InvoiceType  InvoiceType `json:"invoiceType"`
	FkCustomer   string      `json:"fkCustomer"`
	PurchaseList []Purchase  `json:"purchaseList"`
	TotalSum     float64     `json:"totalSum"`
}

func (receiver Invoice) String() string {
	var sb strings.Builder

	sb.WriteString("ID: " + receiver.ID)
	sb.WriteString("\nIssued: " + receiver.Issued)
	sb.WriteString("\nInvoiceType:\n" + receiver.InvoiceType.String())
	sb.WriteString("\nFkCustomer: " + receiver.FkCustomer)
	sb.WriteString("\nTotalSum: " + fmt.Sprintf("%.2f", receiver.TotalSum))

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
