package customers_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/klintonlee/simple-bank-api/internal/modules/customers"
	mock_customer "github.com/klintonlee/simple-bank-api/internal/modules/customers/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCustomerService_FindCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	customer := mock_customer.NewMockCustomerInterface(ctrl)
	store := mock_customer.NewMockCustomerStoreInterface(ctrl)
	service := customers.NewCustomerService(store)

	store.EXPECT().FindByCpf("abc").Return(customer, nil)
	result, err := service.FindCustomerByCpf("abc")
	assert.Nil(t, err)
	assert.Equal(t, customer, result)

	store.EXPECT().FindByCpf("abc").Return(nil, errors.New("Some error"))
	result, err = service.FindCustomerByCpf("abc")
	assert.Nil(t, result)
	assert.Equal(t, "Some error", err.Error())
}

func TestCustomerService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	customer := mock_customer.NewMockCustomerInterface(ctrl)
	store := mock_customer.NewMockCustomerStoreInterface(ctrl)
	service := customers.NewCustomerService(store)

	store.EXPECT().Save(gomock.Any()).Return(customer, nil)
	result, err := service.Create("a", "b", "c")
	assert.Nil(t, err)
	assert.Equal(t, customer, result)

	store.EXPECT().Save(gomock.Any()).Return(nil, errors.New("Store error"))
	_, err = service.Create("a", "b", "c")
	assert.Equal(t, "Store error", err.Error())
}
