package transactions_test

import (
	"testing"

	"github.com/klintonlee/simple-bank-api/internal/modules/transactions"
	"github.com/stretchr/testify/assert"
)

func TestTransaction_NewTransaction(t *testing.T) {
	transaction := transactions.NewTransaction("abc", "deposit", 10)

	assert.NotNil(t, transaction.ID)
	assert.Equal(t, "abc", transaction.AccountID)
	assert.Equal(t, "deposit", transaction.TransactionType)
	assert.Equal(t, 10.0, transaction.Value)
	assert.NotNil(t, transaction.CreatedAt)
}

func TestTransaction_IsValid(t *testing.T) {
	transaction := transactions.NewTransaction("abc", "wrong type", 0)

	_, err := transaction.IsValid()
	assert.Equal(t, "transaction type must be deposit ou withdrawal", err.Error())

	transaction.TransactionType = "deposit"
	_, err = transaction.IsValid()
	assert.Equal(t, "transaction value must be greater than 10.00", err.Error())

	transaction.TransactionType = "withdrawal"
	transaction.Value = 10.0
	result, err := transaction.IsValid()
	assert.Equal(t, true, result)
	assert.Nil(t, err)
}
