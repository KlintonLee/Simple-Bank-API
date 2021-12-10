package handler

import (
	"encoding/json"
	"net/http"

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
	}
	customerService := customers.NewCustomerService(c.store)
	customer, err := customerService.Create(customerDTO.Name, customerDTO.Cpf, customerDTO.Birth)
	if err != nil {
		sendResponse(w, Response{Message: "Something went wrong, contact the support"}, http.StatusUnprocessableEntity)
	}

	sendResponse(w, customer, http.StatusCreated)
}
