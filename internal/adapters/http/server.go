package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/klintonlee/simple-bank-api/internal/adapters/db"
	"github.com/klintonlee/simple-bank-api/internal/adapters/dtos"
	"github.com/klintonlee/simple-bank-api/internal/modules/customers"
)

type Response struct {
	Message string `json:"message"`
}

func Run(store *db.CustomerDb) error {
	handler := mux.NewRouter()
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode(Response{Message: "I am Alive"})
		if err != nil {
			panic(err)
		}
	})

	handler.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		var customerDto dtos.CustomerDto
		err := json.NewDecoder(r.Body).Decode(&customerDto)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{Message: "Failure on creating new Customer"})
			return
		}

		customerService := customers.NewCustomerService(store)
		customer, err := customerService.Create(customerDto.ID, customerDto.Name, customerDto.Cpf, customerDto.Birth)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{Message: "Failure on creating new Customer"})
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(customer)
	}).Methods("POST")
	err := http.ListenAndServe(":3333", handler)
	if err != nil {
		return err
	}
	return nil
}
