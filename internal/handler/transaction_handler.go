package handler

import (
	"encoding/json"
	"fmt"
	"go-pos-app/internal/model"
	"go-pos-app/internal/service"
	"net/http"
)

type TransactionHandler struct {
	Service *service.TransactionService
}

type TransactionRequest struct{
	UserID int `json:"user_id"`
	Items []model.TransactionItem `json:"items"`
}


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

	fmt.Println("REQUEST:", req)

	err = h.Service.CreateTransaction(req.UserID, req.Items)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Transaction Created"))
}