package model

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Equipment struct {
	ID             string    `json:"id"`
	Description    string    `json:"description"`
	Weight         string    `json:"weight"`
	Size           string    `json:"size"`
	RadiationLevel string    `json:"radiationLevel"`
	ProductionDate time.Time `json:"productionDate"`
	Quantity       int       `json:"quantity"`
	EquipmentType  string    `json:"equipmentType"`
}

func (receiver Equipment) String() string {
	var sb strings.Builder

	sb.WriteString("ID: " + receiver.ID)
	sb.WriteString("\nDescription: " + receiver.Description)
	sb.WriteString("\nWeight: " + receiver.Weight)
	sb.WriteString("\nSize: " + receiver.Size)
	sb.WriteString("\nRadiationLevel: " + receiver.RadiationLevel)
	sb.WriteString("\nProductionDate: " + receiver.ProductionDate.String())
	sb.WriteString("\nQuantity: " + fmt.Sprintf("%d", receiver.Quantity))
	sb.WriteString("\nEquipmentType: " + receiver.EquipmentType)

	return sb.String()
}

func (receiver Equipment) Json() ([]byte, error) {
	return json.Marshal(receiver)
}

func (receiver Equipment) GetID() string {
	return receiver.ID
}
