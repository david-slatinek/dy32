package random

import (
	"github.com/google/uuid"
	"main/model"
	"math/rand"
	"strings"
	"time"
)

func Invoice() model.Info {
	invoice := model.Invoice{
		ID:          strings.ReplaceAll(uuid.NewString(), "-", ""),
		Issued:      time.Now().Format("2006-01-02 15:04:05"),
		InvoiceType: Types[rand.Intn(Size)],
		FkCustomer:  strings.ReplaceAll(uuid.NewString(), "-", ""),
	}

	count := Int(3, 5)

	for i := 0; i < count; i++ {
		p := Purchase().(model.Purchase)
		invoice.PurchaseList = append(invoice.PurchaseList, p)
		invoice.TotalSum += invoice.TotalSum + p.PartialSum
	}

	return invoice
}
