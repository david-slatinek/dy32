package model

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Purchase struct {
	ID          string  `bson:"id" json:"id"`
	Quantity    int     `bson:"quantity" json:"quantity"`
	Price       float64 `bson:"price" json:"price"`
	PartialSum  float64 `bson:"partialSum" json:"partialSum"`
	FkEquipment string  `bson:"fkEquipment" json:"fkEquipment"`
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

func (receiver Purchase) Json() ([]byte, error) {
	return json.Marshal(receiver)
}
