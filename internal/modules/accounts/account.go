package accounts

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Account struct {
	ID                   string
	CustomerID           string
	Balance              float64
	WithdrawalDailyLimit float64
	IsActive             bool
	AccountType          int8
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

type AccountInterface interface {
	GetID() string
	GetCustomerID() string
	GetBalance() float64
	GetWithdrawalDailyLimit() float64
	GetIsActive() bool
	GetAccountType() int8
}

func NewAccount(customerID string, withdrawalDailyLimit float64, accountType int8) *Account {
	account := Account{
		ID:                   uuid.NewV4().String(),
		CustomerID:           customerID,
		Balance:              0.0,
		WithdrawalDailyLimit: withdrawalDailyLimit,
		IsActive:             true,
		AccountType:          accountType,
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}
	return &account
}

func (a *Account) GetID() string {
	return a.ID
}

func (a *Account) GetCustomerID() string {
	return a.CustomerID
}

func (a *Account) GetBalance() float64 {
	return a.Balance
}

func (a *Account) GetWithdrawalDailyLimit() float64 {
	return a.WithdrawalDailyLimit
}

func (a *Account) GetIsActive() bool {
	return a.IsActive
}

func (a *Account) GetAccountType() int8 {
	return a.AccountType
}
