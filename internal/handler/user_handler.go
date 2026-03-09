package handler

import (
	"encoding/json"
	"go-pos-app/internal/model"
	"go-pos-app/internal/service"
	"net/http"
	"strconv"
)
 
type UserHandler struct {
	Service *service.UserService
}

func (h *UserHandler) HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetUsers(w,r)
	case http.MethodPost:
		h.CreateUsers(w,r)
	case http.MethodPut:
		h.UpdateUser(w,r)
	case http.MethodDelete:
		h.DeleteUser(w,r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-Type", "application/json")

	users, err := h.Service.GetUsers()

	if err != nil {
		http.Error(w, "Failed fetch Users", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) CreateUsers(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-Type", "application/json")

	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	err = h.Service.CreateUsers(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "User Created",
	})
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "invalid Request", http.StatusBadRequest)
		return
	}

	err = h.Service.UpdateUser(user)

	if err != nil {
		http.Error(w, "Failed update User", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message" : "User Updated",
	})
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteUser(id)
	if err != nil{
		http.Error(w, "Failed delete user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "User Delete",
	})
}