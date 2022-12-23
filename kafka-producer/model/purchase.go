package model

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Purchase struct {
	ID          string  `json:"id"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	PartialSum  float64 `json:"partialSum"`
	FkEquipment string  `json:"fkEquipment"`
}

func (receiver Purchase) Json() ([]byte, error) {
	return json.Marshal(receiver)
}

func (receiver Purchase) GetID() string {
	return receiver.ID
}

func (receiver Purchase) String() string {
	var sb strings.Builder

	sb.WriteString("ID: " + receiver.ID)
	sb.WriteString("\nQuantity: " + fmt.Sprintf("%d", receiver.Quantity))
	sb.WriteString("\nPrice: " + fmt.Sprintf("%.2f", receiver.Price))
	sb.WriteString("\nPartialSum: " + fmt.Sprintf("%.2f", receiver.PartialSum))
	sb.WriteString("\nFkEquipment: " + receiver.FkEquipment)

	return sb.String()
}
