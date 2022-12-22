package random

import "kafka-producer/model"

const Size = 3

var Types = [3]model.InvoiceType{
	{
		ID:   1,
		Type: String(5, 10),
	},
	{
		ID:   2,
		Type: String(5, 10),
	},
	{
		ID:   3,
		Type: String(5, 10),
	},
}
