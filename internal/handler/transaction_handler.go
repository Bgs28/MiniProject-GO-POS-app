package handler

import (
	"encoding/json"
	"go-pos-app/internal/model"
	"go-pos-app/internal/service"
	"net/http"
	"strconv"
)

type TransactionHandler struct {
	Service *service.TransactionService
}

type TransactionRequest struct{
	UserID int `json:"user_id"`
	Items []model.TransactionItem `json:"items"`
}

// create transaction handler
func (h *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		http.Error(w, "Methode not Allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var req TransactionRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Service.CreateTransaction(req.UserID, req.Items)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Transaction Created"))
}

// transaction detail handler
func (h *TransactionHandler) GetTransactionDetail(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")

	if idStr == "" {
		http.Error(w, "transaction id required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	result, err := h.Service.GetTransactionDetail(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// get all transaction handler
func (h *TransactionHandler) GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	 
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	transactions, err := h.Service.GetAllTransactions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}

// sales report handler
func (h *TransactionHandler) GetSalesReport(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "method Not allowed", http.StatusMethodNotAllowed)
		return
	}

	report, err := h.Service.GetSalesReport()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}

// dashboard stats handler
func (h *TransactionHandler) GetDashboardStats(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	stats, err := h.Service.GetDashboardStats()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}