package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/klintonlee/simple-bank-api/internal/adapters/db"
)

type Response struct {
	Message string `json:"message"`
}

type Handler struct {
	Router        *mux.Router
	customerStore *db.CustomerDb
}

func NewHandler(customerStore *db.CustomerDb) *Handler {
	return &Handler{
		customerStore: customerStore,
	}
}

func (h *Handler) SetupRoutes() {
	h.Router = mux.NewRouter()

	customerHandler := CustomerHandler{store: h.customerStore}
	h.Router.HandleFunc("/api/v1/customers", customerHandler.CreateCustomer).Methods("POST")
	h.Router.HandleFunc("/api/v1/customers/{cpf}", customerHandler.ShowCustomer).Methods("GET")

	h.Router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(Response{Message: "I'm alive"})
	}).Methods("GET")
}

func sendResponse(w http.ResponseWriter, resp interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(resp)
}
