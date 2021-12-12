package customers

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Customer struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf"`
	Birth     string    `json:"birth"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
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
	_, err := uuid.FromString(c.ID)
	if err != nil {
		return false, errors.New("customer ID is invalid")
	}
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
