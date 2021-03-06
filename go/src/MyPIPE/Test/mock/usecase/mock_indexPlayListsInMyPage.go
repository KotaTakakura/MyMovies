// Code generated by MockGen. DO NOT EDIT.
// Source: ./usecase/IndexPlayListsInMyPage.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	queryService "MyPIPE/domain/queryService"
	usecase "MyPIPE/usecase"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIIndexPlayListsInMyPage is a mock of IIndexPlayListsInMyPage interface
type MockIIndexPlayListsInMyPage struct {
	ctrl     *gomock.Controller
	recorder *MockIIndexPlayListsInMyPageMockRecorder
}

// MockIIndexPlayListsInMyPageMockRecorder is the mock recorder for MockIIndexPlayListsInMyPage
type MockIIndexPlayListsInMyPageMockRecorder struct {
	mock *MockIIndexPlayListsInMyPage
}

// NewMockIIndexPlayListsInMyPage creates a new mock instance
func NewMockIIndexPlayListsInMyPage(ctrl *gomock.Controller) *MockIIndexPlayListsInMyPage {
	mock := &MockIIndexPlayListsInMyPage{ctrl: ctrl}
	mock.recorder = &MockIIndexPlayListsInMyPageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIIndexPlayListsInMyPage) EXPECT() *MockIIndexPlayListsInMyPageMockRecorder {
	return m.recorder
}

// All mocks base method
func (m *MockIIndexPlayListsInMyPage) All(indexPlayListsInMyPageDTO *usecase.IndexPlayListsInMyPageDTO) *queryService.IndexPlayListsInMyPageDTO {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", indexPlayListsInMyPageDTO)
	ret0, _ := ret[0].(*queryService.IndexPlayListsInMyPageDTO)
	return ret0
}

// All indicates an expected call of All
func (mr *MockIIndexPlayListsInMyPageMockRecorder) All(indexPlayListsInMyPageDTO interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockIIndexPlayListsInMyPage)(nil).All), indexPlayListsInMyPageDTO)
}
