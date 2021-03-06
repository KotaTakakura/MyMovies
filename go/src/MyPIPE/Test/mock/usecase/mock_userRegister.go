// Code generated by MockGen. DO NOT EDIT.
// Source: ./usecase/UserRegister.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	model "MyPIPE/domain/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIUserRegister is a mock of IUserRegister interface
type MockIUserRegister struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRegisterMockRecorder
}

// MockIUserRegisterMockRecorder is the mock recorder for MockIUserRegister
type MockIUserRegisterMockRecorder struct {
	mock *MockIUserRegister
}

// NewMockIUserRegister creates a new mock instance
func NewMockIUserRegister(ctrl *gomock.Controller) *MockIUserRegister {
	mock := &MockIUserRegister{ctrl: ctrl}
	mock.recorder = &MockIUserRegisterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIUserRegister) EXPECT() *MockIUserRegisterMockRecorder {
	return m.recorder
}

// RegisterUser mocks base method
func (m *MockIUserRegister) RegisterUser(newUser *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", newUser)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterUser indicates an expected call of RegisterUser
func (mr *MockIUserRegisterMockRecorder) RegisterUser(newUser interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockIUserRegister)(nil).RegisterUser), newUser)
}
