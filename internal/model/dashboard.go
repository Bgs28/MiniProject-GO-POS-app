package model

type DashboardStats struct {
	TotalProducts     int `json:"total_product"`
	TotalTransactions int `json:"total_transactions"`
	TodaySales        int `json:"today_sales"`
}