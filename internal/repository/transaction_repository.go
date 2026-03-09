package repository

import (
	"database/sql"
	"go-pos-app/internal/model"
)

type TransactionRepository struct {
	DB *sql.DB
}

// create/insert transaction
func (r *TransactionRepository) CreateTransaction(tx *sql.Tx, transaction model.Transaction) (int, error) {
	query := `
		INSERT INTO transactions (invoice_number, user_id, total_price)
		VALUES (?,?,?)
	`

	result, err := tx.Exec(
		query, 
		transaction.InvoiceNumber,
		transaction.UserID,
		transaction.TotalPrice,
	)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil{
		return 0, err
	}

	return int(id), nil
}

// create/insert transaction item

func ( r *TransactionRepository) CreateTransactionItem(tx *sql.Tx, item model.TransactionItem) error{
	query := `
		INSERT INTO transaction_items
		(transaction_id, product_id, quantity, price, subtotal)
		VALUES (?,?,?,?,?)`

		_, err := tx.Exec(
			query,
			item.TransactionID,
			item.ProductID,
			item.Quantity,
			item.Price,
			item.Subtotal,
		)

		return err
}

// update stock after transaction
func (r *TransactionRepository) UpdateProductStock( 
	tx *sql.Tx, productID int, qty int, 
) error{
	query := `
		UPDATE products SET stock = stock - ?
		WHERE id = ?
	`

	_, err := tx.Exec(query, qty, productID)

	return err
}

// Get Product by ID

func (r *TransactionRepository) GetProductByID(ProductID int) (model.Product, error){
	var product model.Product

	query := `
		SELECT id, name, price, stock
		FROM products WHERE id = ?
	`

	err := r.DB.QueryRow(query, ProductID).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.Stock,
	)

	return product, err
}