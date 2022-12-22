package random

import (
	"github.com/google/uuid"
	"kafka-producer/model"
)

func purchase() model.Purchase {
	p := model.Purchase{
		ID:          uuid.NewString(),
		Quantity:    Int(3, 5),
		Price:       Float(10, 2000),
		FkEquipment: uuid.NewString(),
	}
	p.PartialSum = float64(p.Quantity) * p.Price
	return p
}