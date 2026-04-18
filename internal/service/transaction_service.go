package service

import (
	"database/sql"
	"fmt"
	"go-pos-app/internal/model"
	"go-pos-app/internal/repository"
	"time"
)

type TransactionService struct {
	DB *sql.DB
	Repo *repository.TransactionRepository
}

// create transaction service
func (s *TransactionService) CreateTransaction(UserID int, items []model.TransactionItem)error{
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	var totalPrice int

	// count subtotal/item
	for i := range items {

		product, err := s.Repo.GetProductByID(items[i].ProductID)
		if err != nil {
			tx.Rollback()
			return err
		}

		if product.Stock < items[i].Quantity {
			tx.Rollback()
			return fmt.Errorf("Stock tidak Cukup untuk Produk %s", product.Name)
		}

		items[i].Price = product.Price
		items[i].Subtotal = product.Price * items[i].Quantity

		totalPrice += items[i].Subtotal
	}

	invoice := fmt.Sprintf("INV-%d", time.Now().Unix())

	transaction := model.Transaction{
		InvoiceNumber: invoice,
		UserID: UserID,
		TotalPrice: totalPrice,
	}

	transactionID, err := s.Repo.CreateTransaction(tx, transaction)
	if err != nil {
		tx.Rollback()
		return err
	}

	// insert items

	for _, item := range items {

		item.TransactionID = transactionID

		err := s.Repo.CreateTransactionItem(tx, item)
		if err != nil {
			tx.Rollback()
			return err
		}

		err = s.Repo.UpdateProductStock(tx, item.ProductID, item.Quantity)
		if err != nil {
			tx.Rollback()
			return err
		}

	}

	return tx.Commit()
}


// get detail transaction service
func (s *TransactionService) GetTransactionDetail(id int) (*model.TransactionDetail, error) {

	transaction, err := s.Repo.GetTransactionByID(id)
	if err != nil {
		return nil, err
	}

	items, err := s.Repo.GetTransactionItems(id)
	if err != nil {
		return nil, err
	}

	result := model.TransactionDetail{
		ID: transaction.ID,
		InvoiceNumber: transaction.InvoiceNumber,
		UserID: transaction.UserID,
		TotalPrice: transaction.TotalPrice,
		CreatedAt: transaction.CreatedAt,
		Items: items,
	}

	return &result, nil
}

// get all transaction service
func (s *TransactionService) GetAllTransactions() ([]model.Transaction, error) {

	transactions, err := s.Repo.GetAllTransactions()
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

// get sales report service
func (s *TransactionService) GetSalesReport() (model.SalesReport, error){
	report, err := s.Repo.GetSalesReport()

	if err != nil {
		return report, err
	}

	return report, nil
}

// dashboard stats service

func (s *TransactionService) GetDashboardStats() (model.DashboardStats, error) {
	stats, err := s.Repo.GetDashboardStats()

	if err != nil {
		return stats, err
	}

	return stats, nil
}