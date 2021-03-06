// Code generated by MockGen. DO NOT EDIT.
// Source: ./usecase/ChangePassword.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	usecase "MyPIPE/usecase"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIChangePassword is a mock of IChangePassword interface
type MockIChangePassword struct {
	ctrl     *gomock.Controller
	recorder *MockIChangePasswordMockRecorder
}

// MockIChangePasswordMockRecorder is the mock recorder for MockIChangePassword
type MockIChangePasswordMockRecorder struct {
	mock *MockIChangePassword
}

// NewMockIChangePassword creates a new mock instance
func NewMockIChangePassword(ctrl *gomock.Controller) *MockIChangePassword {
	mock := &MockIChangePassword{ctrl: ctrl}
	mock.recorder = &MockIChangePasswordMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIChangePassword) EXPECT() *MockIChangePasswordMockRecorder {
	return m.recorder
}

// ChangePassword mocks base method
func (m *MockIChangePassword) ChangePassword(changePasswordDTO *usecase.ChangePasswordDTO) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", changePasswordDTO)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangePassword indicates an expected call of ChangePassword
func (mr *MockIChangePasswordMockRecorder) ChangePassword(changePasswordDTO interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockIChangePassword)(nil).ChangePassword), changePasswordDTO)
}
