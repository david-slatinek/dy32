package random

import (
	"github.com/google/uuid"
	"main/model"
	"strings"
)

func Purchase() model.Info {
	p := model.Purchase{
		ID:          strings.ReplaceAll(uuid.NewString(), "-", ""),
		Quantity:    Int(3, 10),
		Price:       Float(10, 2000),
		FkEquipment: strings.ReplaceAll(uuid.NewString(), "-", ""),
		//FkEquipment: "1",
	}
	p.PartialSum = float64(p.Quantity) * p.Price
	return p
}
