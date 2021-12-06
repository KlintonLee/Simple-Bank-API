package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Message string `json:"message"`
}

func Run() error {
	handler := mux.NewRouter()
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode(Response{Message: "I am Alive"})
		if err != nil {
			panic(err)
		}
	})
	err := http.ListenAndServe(":3333", handler)
	if err != nil {
		return err
	}
	return nil
}
