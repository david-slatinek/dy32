package model

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Customer struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Lastname     string `json:"lastname"`
	StreetName   string `json:"streetName"`
	StreetNumber string `json:"streetNumber"`
	PostNumber   int    `json:"postNumber"`
	PostTitle    string `json:"postTitle"`
}

func (receiver Customer) String() string {
	var sb strings.Builder

	sb.WriteString("ID: " + receiver.ID)
	sb.WriteString("\nName: " + receiver.Name)
	sb.WriteString("\nLastname: " + receiver.Lastname)
	sb.WriteString("\nStreetName: " + receiver.StreetName)
	sb.WriteString("\nStreetNumber: " + receiver.StreetNumber)
	sb.WriteString("\nPostNumber: " + fmt.Sprintf("%d", receiver.PostNumber))
	sb.WriteString("\nPostTitle: " + receiver.PostTitle)

	return sb.String()
}

func (receiver Customer) Json() ([]byte, error) {
	return json.Marshal(receiver)
}

func (receiver Customer) GetID() string {
	return receiver.ID
}
