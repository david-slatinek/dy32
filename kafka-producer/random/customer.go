package random

import (
	"github.com/google/uuid"
	"main/model"
	"strings"
)

func Customer() model.Info {
	return model.Customer{
		ID:           strings.ReplaceAll(uuid.NewString(), "-", ""),
		Name:         String(5, 10),
		Lastname:     String(5, 15),
		StreetName:   String(5, 8),
		StreetNumber: String(1, 3),
		PostNumber:   Int(1000, 9000),
		PostTitle:    String(5, 10),
	}
}
