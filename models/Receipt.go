package models

type Receipt struct {
	Retailer     string `json:"retailer" validate:"required,alphanum"`
	PurchaseDate string `json:"purchaseDate" validate:"required,date"`
	PurchaseTime string `json:"purchaseTime" validate:"required,time"`
	Items        []Item `json:"items" validate:"required,dive"`
	Total        string `json:"total" validate:"required,price"`
}
