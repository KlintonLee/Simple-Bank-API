package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/klintonlee/simple-bank-api/internal/adapters/dtos"
	"github.com/klintonlee/simple-bank-api/internal/modules/customers"
)

type CustomerHandler struct {
	store customers.CustomerStoreInterface
}

func (c *CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customerDTO dtos.CustomerDto
	json.NewDecoder(r.Body).Decode(&customerDTO)
	err := customerDTO.CheckFilled()
	if err != nil {
		sendResponse(w, Response{Message: err.Error()}, http.StatusBadRequest)
		return
	}
	customerService := customers.NewCustomerService(c.store)
	customerExists, _ := customerService.FindCustomerByCpf(customerDTO.Cpf)
	if customerExists != nil {
		sendResponse(w, Response{Message: "Customer already exists"}, http.StatusConflict)
		return
	}
	customer, err := customerService.Create(customerDTO.Name, customerDTO.Cpf, customerDTO.Birth)
	if err != nil {
		sendResponse(w, Response{Message: "Something went wrong, contact the support"}, http.StatusUnprocessableEntity)
		return
	}

	sendResponse(w, customer, http.StatusCreated)
}

func (c *CustomerHandler) ShowCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cpf := vars["cpf"]

	customerExists, _ := c.store.FindByCpf(cpf)
	if customerExists != nil {
		sendResponse(w, customerExists, http.StatusOK)
		return
	}
	sendResponse(w, Response{Message: "User does not exists"}, http.StatusNotFound)
}
