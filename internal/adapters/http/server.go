package http

import (
	"net/http"

	"github.com/klintonlee/simple-bank-api/internal/adapters/db"
	"github.com/klintonlee/simple-bank-api/internal/adapters/http/handler"
)

type Response struct {
	Message string `json:"message"`
}

func Run(customerStore *db.CustomerDb) error {
	handler := handler.NewHandler(customerStore)
	handler.SetupRoutes()

	err := http.ListenAndServe(":3333", handler.Router)
	if err != nil {
		return err
	}
	return nil
}
