package util

import (
	"github.com/google/uuid"
	"kafka-producer/model"
	"math/rand"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var types = [3]model.InvoiceType{
	{
		ID:   1,
		Type: randomString(),
	},
	{
		ID:   2,
		Type: randomString(),
	},
	{
		ID:   3,
		Type: randomString(),
	},
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func randomString() string {
	return random(rand.Intn(40-30) + 30)
}

func random(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func randomPurchase() model.Purchase {
	purchase := model.Purchase{
		ID:          uuid.NewString(),
		Quantity:    rand.Intn(5-1) + 5,
		Price:       rand.Float64()*(2000-10) + 10,
		FkEquipment: uuid.NewString(),
	}
	purchase.PartialSum = float64(purchase.Quantity) * purchase.Price
	return purchase
}

func RandomInvoice() model.Invoice {
	invoice := model.Invoice{
		ID:          uuid.NewString(),
		Issued:      time.Now(),
		InvoiceType: types[rand.Intn(len(types))],
		FkCustomer:  uuid.NewString(),
	}

	count := rand.Intn(5-1) + 5

	for i := 0; i < count; i++ {
		p := randomPurchase()
		invoice.PurchaseList = append(invoice.PurchaseList, p)
		invoice.TotalSum += invoice.TotalSum + p.PartialSum
	}

	return invoice
}
