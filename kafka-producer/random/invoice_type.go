package random

import "main/model"

var Size = len(Types)

var Types = []model.InvoiceType{
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
