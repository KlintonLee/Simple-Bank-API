package customers_test

import (
	"testing"

	"github.com/klintonlee/simple-bank-api/internal/modules/customers"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestCustomer_IsValid(t *testing.T) {
	customer := customers.Customer{}

	_, err := customer.IsValid()
	assert.Equal(t, "customer ID is invalid", err.Error())

	customer.ID = uuid.NewV4().String()
	_, err = customer.IsValid()
	assert.Equal(t, "customer name is required", err.Error())

	customer.Name = "John Doe"
	_, err = customer.IsValid()
	assert.Equal(t, "customer CPF is required", err.Error())

	customer.Cpf = "12345678900"
	_, err = customer.IsValid()
	assert.Equal(t, "customer birth is required", err.Error())

	customer.Birth = "2001-01-20"
	result, err := customer.IsValid()
	assert.Nil(t, err)
	assert.Equal(t, true, result)
}

func TestCustomer_GetID(t *testing.T) {
	customer := customers.NewCustomer()
	id := customer.GetID()

	_, err := uuid.FromString(id)
	assert.Nil(t, err)
}

func TestCustomer_GetName(t *testing.T) {
	customer := customers.Customer{}
	customer.Name = "John Doe"

	name := customer.GetName()
	assert.Equal(t, customer.Name, name)
}

func TestCustomer_GetCpf(t *testing.T) {
	customer := customers.Customer{}
	customer.Cpf = "12345678900"

	cpf := customer.GetCpf()
	assert.Equal(t, customer.Cpf, cpf)
}

func TestCustomer_GetBirth(t *testing.T) {
	customer := customers.Customer{}
	customer.Birth = "2021-10-27"

	birth := customer.GetBirth()
	assert.Equal(t, customer.Birth, birth)
}
