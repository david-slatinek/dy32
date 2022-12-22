package model

import "fmt"

type InvoiceType struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

func (receiver InvoiceType) String() string {
	return fmt.Sprintf("ID: %d", receiver.ID) + fmt.Sprintf("\nType: %s", receiver.Type)
}
