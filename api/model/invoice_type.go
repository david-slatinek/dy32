package model

import "fmt"

type InvoiceType struct {
	ID   int    `bson:"id" json:"id"`
	Type string `bson:"type" json:"type"`
}

func (receiver InvoiceType) String() string {
	return fmt.Sprintf("ID: %d", receiver.ID) + fmt.Sprintf("\nType: %s", receiver.Type)
}
