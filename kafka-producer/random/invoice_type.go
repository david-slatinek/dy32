package random

import "main/model"

var Size = len(Types)

var Types = []model.InvoiceType{
	{
		ID:   1,
		Type: String(5, 10),
		//Type: "final",
	},
	{
		ID:   2,
		Type: String(5, 10),
		//Type: "final",
	},
	{
		ID:   3,
		Type: String(5, 10),
		//Type: "final",
	},
}
