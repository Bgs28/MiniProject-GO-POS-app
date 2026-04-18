package model

type TransactionDetail struct {
	ID            int                     `json:"id"`
	InvoiceNumber string                  `json:"invoice_number"`
	UserID        int                     `json:"user_id"`
	TotalPrice    int                     `json:"total_price"`
	CreatedAt     string                  `json:"created_at"`
	Items         []TransactionItemDetail `json:"items"`
}

type TransactionItemDetail struct {
	ProductID   int    `json:"product_id"`
	ProductName string `json:"product_name"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
	Subtotal    int    `json:"subtotal"`
}