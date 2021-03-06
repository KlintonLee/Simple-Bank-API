// Code generated by MockGen. DO NOT EDIT.
// Source: internal/modules/customers/customer.go

// Package mock_customer is a generated GoMock package.
package mock_customer

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCustomerInterface is a mock of CustomerInterface interface.
type MockCustomerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCustomerInterfaceMockRecorder
}

// MockCustomerInterfaceMockRecorder is the mock recorder for MockCustomerInterface.
type MockCustomerInterfaceMockRecorder struct {
	mock *MockCustomerInterface
}

// NewMockCustomerInterface creates a new mock instance.
func NewMockCustomerInterface(ctrl *gomock.Controller) *MockCustomerInterface {
	mock := &MockCustomerInterface{ctrl: ctrl}
	mock.recorder = &MockCustomerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomerInterface) EXPECT() *MockCustomerInterfaceMockRecorder {
	return m.recorder
}

// GetBirth mocks base method.
func (m *MockCustomerInterface) GetBirth() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBirth")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetBirth indicates an expected call of GetBirth.
func (mr *MockCustomerInterfaceMockRecorder) GetBirth() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBirth", reflect.TypeOf((*MockCustomerInterface)(nil).GetBirth))
}

// GetCpf mocks base method.
func (m *MockCustomerInterface) GetCpf() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCpf")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetCpf indicates an expected call of GetCpf.
func (mr *MockCustomerInterfaceMockRecorder) GetCpf() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCpf", reflect.TypeOf((*MockCustomerInterface)(nil).GetCpf))
}

// GetID mocks base method.
func (m *MockCustomerInterface) GetID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetID indicates an expected call of GetID.
func (mr *MockCustomerInterfaceMockRecorder) GetID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockCustomerInterface)(nil).GetID))
}

// GetName mocks base method.
func (m *MockCustomerInterface) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockCustomerInterfaceMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockCustomerInterface)(nil).GetName))
}

// IsValid mocks base method.
func (m *MockCustomerInterface) IsValid() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValid")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsValid indicates an expected call of IsValid.
func (mr *MockCustomerInterfaceMockRecorder) IsValid() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValid", reflect.TypeOf((*MockCustomerInterface)(nil).IsValid))
}
