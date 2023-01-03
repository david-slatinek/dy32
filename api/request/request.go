package request

import "main/model"

type InvoiceRequest struct {
	InvoiceType  model.InvoiceType `json:"invoiceType"`
	FkCustomer   string            `json:"fkCustomer"`
	PurchaseList []model.Purchase  `json:"purchaseList"`
	TotalSum     float32           `json:"totalSum"`
}
