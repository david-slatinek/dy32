package random

import (
	"github.com/google/uuid"
	"main/model"
	"time"
)

func Equipment() model.Info {
	return model.Equipment{
		ID:             uuid.NewString(),
		Description:    String(5, 30),
		Weight:         String(1, 5),
		Size:           String(5, 10),
		RadiationLevel: String(3, 5),
		ProductionDate: time.Now(),
		Quantity:       Int(3, 50),
		EquipmentType:  String(5, 20),
	}
}
