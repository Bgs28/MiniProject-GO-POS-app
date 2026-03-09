package model

type TransactionItem struct {
	ID            int `json:"id"`
	TransactionID int `json:"transaction_id"`
	ProductID     int `json:"product_id"`
	Quantity      int `json:"quantity"`
	Price         int `json:"price"`
	Subtotal      int `json:"subtotal"`
}