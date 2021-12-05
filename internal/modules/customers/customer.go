package customers

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

type Customer struct {
	ID    string
	Name  string
	Cpf   string
	Birth string
}

type CustomerInterface interface {
	IsValid() (bool, error)
	GetID() string
	GetName() string
	GetCpf() string
	GetBirth() string
}

func NewCustomer() *Customer {
	customer := Customer{
		ID: uuid.NewV4().String(),
	}
	return &customer
}

func (c *Customer) IsValid() (bool, error) {
	if c.Name == "" {
		return false, errors.New("customer name is required")
	}
	if c.Cpf == "" {
		return false, errors.New("customer CPF is required")
	}
	if c.Birth == "" {
		return false, errors.New("customer birth is required")
	}
	return true, nil
}

func (c *Customer) GetID() string {
	return c.ID
}

func (c *Customer) GetName() string {
	return c.Name
}

func (c *Customer) GetCpf() string {
	return c.Cpf
}

func (c *Customer) GetBirth() string {
	return c.Birth
}
