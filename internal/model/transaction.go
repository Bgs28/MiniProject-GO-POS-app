package model

type Transaction struct {
	ID            int               `json:"id"`
	InvoiceNumber string            `json:"invoice_number"`
	UserID        int               `json:"user_id"`
	TotalPrice    int               `json:"total_price"`
	Items         []TransactionItem `json:"items"`
}