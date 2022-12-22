package random

import (
	"github.com/google/uuid"
	"kafka-producer/model"
	"math/rand"
	"time"
)

func Invoice() model.Invoice {
	invoice := model.Invoice{
		ID:          uuid.NewString(),
		Issued:      time.Now(),
		InvoiceType: Types[rand.Intn(Size)],
		FkCustomer:  uuid.NewString(),
	}

	count := Int(3, 5)

	for i := 0; i < count; i++ {
		p := purchase()
		invoice.PurchaseList = append(invoice.PurchaseList, p)
		invoice.TotalSum += invoice.TotalSum + p.PartialSum
	}

	return invoice
}
