package model

type SalesReport struct {
	TotalSales        int `json:"total_sales"`
	TotalTransactions int `json:"total_transactions"`
	TodaySales        int `json:"today_sales"`
}