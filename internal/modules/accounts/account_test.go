package accounts_test

import (
	"testing"

	"github.com/klintonlee/simple-bank-api/internal/modules/accounts"
	"github.com/stretchr/testify/assert"
)

func TestAccount_NewAccount(t *testing.T) {
	account := accounts.NewAccount("abc", 1000.0, 1)

	assert.Equal(t, account.ID, account.GetID())
	assert.Equal(t, "abc", account.GetCustomerID())
	assert.Equal(t, 0.0, account.GetBalance())
	assert.Equal(t, 1000.0, account.GetWithdrawalDailyLimit())
	assert.Equal(t, true, account.GetIsActive())
	assert.Equal(t, int8(1), account.GetAccountType())
	assert.NotNil(t, account.CreatedAt)
	assert.NotNil(t, account.UpdatedAt)
}
