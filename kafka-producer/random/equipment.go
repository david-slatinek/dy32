package random

import (
	"github.com/google/uuid"
	"main/model"
	"strings"
	"time"
)

func Equipment() model.Info {
	return model.Equipment{
		ID:             strings.ReplaceAll(uuid.NewString(), "-", ""),
		Description:    String(5, 30),
		Weight:         String(1, 5),
		Size:           String(5, 10),
		RadiationLevel: String(3, 5),
		ProductionDate: time.Now().Format("2006-01-02 15:04:05"),
		Quantity:       Int(3, 50),
		EquipmentType:  String(5, 20),
	}
}
