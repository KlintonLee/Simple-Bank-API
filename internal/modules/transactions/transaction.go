package transactions

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Transaction struct {
	ID              string
	AccountID       string
	TransactionType string
	Value           float64
	CreatedAt       time.Time
}

type TransactionInterface interface {
	IsValid() (bool, error)
}

func NewTransaction(accountID, transactionType string, value float64) *Transaction {
	transaction := &Transaction{
		ID:              uuid.NewV4().String(),
		AccountID:       accountID,
		Value:           value,
		TransactionType: transactionType,
		CreatedAt:       time.Now(),
	}
	return transaction
}

func (t *Transaction) IsValid() (bool, error) {
	if t.TransactionType != "deposit" && t.TransactionType != "withdrawal" {
		return false, errors.New("transaction type must be deposit ou withdrawal")
	}
	if t.Value < 10 {
		return false, errors.New("transaction value must be greater than 10.00")
	}
	return true, nil
}
