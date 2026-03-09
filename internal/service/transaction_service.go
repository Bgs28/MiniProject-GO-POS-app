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