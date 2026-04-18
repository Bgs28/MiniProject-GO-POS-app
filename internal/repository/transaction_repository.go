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



// func detail Transaction by ID
func (r *TransactionRepository) GetTransactionByID(id int) (model.Transaction, error) {
	query := `
		SELECT id, invoice_number, user_id, total_price, created_at 
		FROM transactions 
		WHERE id = ?
	`

	var transaction model.Transaction

	err := r.DB.QueryRow(query, id).Scan(
		&transaction.ID,
		&transaction.InvoiceNumber,
		&transaction.UserID,
		&transaction.TotalPrice,
		&transaction.CreatedAt,
	)

	return transaction, err
}

// func get Transaction by id
func (r *TransactionRepository) GetTransactionItems(transactionID int) ([]model.TransactionItemDetail, error) {

	query := `
	SELECT
		ti.product_id,
		p.name, 
		ti.price, 
		ti.quantity, 
		ti.subtotal
	FROM transaction_items ti
	JOIN products p ON ti.product_id = p.id
	WHERE ti.transaction_id = ?
	`

	rows, err := r.DB.Query(query, transactionID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var items []model.TransactionItemDetail

	for rows.Next() {
		var item model.TransactionItemDetail

		err := rows.Scan(
			&item.ProductID,
			&item.ProductName,
			&item.Price,
			&item.Quantity,
			&item.Subtotal,
		)

		if err != nil {
			return nil , err
		}

		items = append(items, item)
	}

	return items, nil
}

// func get transaction all 
func (r *TransactionRepository) GetAllTransactions() ([]model.Transaction, error) {
	query := `
	SELECT id, invoice_number, user_id, total_price, created_at
	FROM transactions
	ORDER BY created_at DESC
	`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var transactions []model.Transaction

	for rows.Next() {

		var transaction model.Transaction

		err := rows.Scan(
			&transaction.ID,
			&transaction.InvoiceNumber,
			&transaction.UserID,
			&transaction.TotalPrice,
			&transaction.CreatedAt,
		) 

		if err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

// func sales Report
func (r *TransactionRepository) GetSalesReport() (model.SalesReport, error) {
	
	var report model.SalesReport

	query := `
	SELECT COALESCE(SUM(total_price),0) as total_sales,
	COUNT(id) as total_transaction
	FROM transactions
	`

	err := r.DB.QueryRow(query).Scan(
		&report.TotalSales,
		&report.TotalTransactions,
	)

	if err != nil {
		return report, err
	}

	queryToday := `
	SELECT COALESCE(SUM(total_price),0) 
	FROM transactions 
	WHERE DATE(created_at) = CURDATE()
	`

	err = r.DB.QueryRow(queryToday).Scan(&report.TodaySales)
	if err != nil {
		return report, err
	}

	return report, nil
}

// func dashboard stats

func (r *TransactionRepository) GetDashboardStats() (model.DashboardStats, error) {
	
	var stats model.DashboardStats

	// total products

	queryProducts := `
	SELECT COUNT(id)
	FROM products
	`

	err := r.DB.QueryRow(queryProducts).Scan(&stats.TotalProducts)

	if err != nil {
		return stats, err
	}

	// total transactions
	
	queryTransactions := `
	SELECT COUNT(id) FROM 
	transactions
	`

	err = r.DB.QueryRow(queryTransactions).Scan(&stats.TotalTransactions)
	if err != nil {
		return stats, err
	}

	// today sales 
	queryTodaySales := `
	SELECT COALESCE(SUM(total_price),0) 
	FROM transactions 
	WHERE DATE(created_at) = CURDATE()
	`

	err = r.DB.QueryRow(queryTodaySales).Scan(&stats.TodaySales)

	if err != nil {
		return stats, err
	}

	return stats, nil

}