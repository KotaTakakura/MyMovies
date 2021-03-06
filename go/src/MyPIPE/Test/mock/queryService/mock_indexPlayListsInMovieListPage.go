// Code generated by MockGen. DO NOT EDIT.
// Source: ./domain/queryService/IndexPlayListsInMovieListPage.go

// Package mock_queryService is a generated GoMock package.
package mock_queryService

import (
	model "MyPIPE/domain/model"
	queryService "MyPIPE/domain/queryService"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIndexPlayListInMovieListPageQueryService is a mock of IndexPlayListInMovieListPageQueryService interface
type MockIndexPlayListInMovieListPageQueryService struct {
	ctrl     *gomock.Controller
	recorder *MockIndexPlayListInMovieListPageQueryServiceMockRecorder
}

// MockIndexPlayListInMovieListPageQueryServiceMockRecorder is the mock recorder for MockIndexPlayListInMovieListPageQueryService
type MockIndexPlayListInMovieListPageQueryServiceMockRecorder struct {
	mock *MockIndexPlayListInMovieListPageQueryService
}

// NewMockIndexPlayListInMovieListPageQueryService creates a new mock instance
func NewMockIndexPlayListInMovieListPageQueryService(ctrl *gomock.Controller) *MockIndexPlayListInMovieListPageQueryService {
	mock := &MockIndexPlayListInMovieListPageQueryService{ctrl: ctrl}
	mock.recorder = &MockIndexPlayListInMovieListPageQueryServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIndexPlayListInMovieListPageQueryService) EXPECT() *MockIndexPlayListInMovieListPageQueryServiceMockRecorder {
	return m.recorder
}

// Find mocks base method
func (m *MockIndexPlayListInMovieListPageQueryService) Find(userId model.UserID, movieId model.MovieID) *queryService.IndexPlayListInMovieListPageDTO {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", userId, movieId)
	ret0, _ := ret[0].(*queryService.IndexPlayListInMovieListPageDTO)
	return ret0
}

// Find indicates an expected call of Find
func (mr *MockIndexPlayListInMovieListPageQueryServiceMockRecorder) Find(userId, movieId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockIndexPlayListInMovieListPageQueryService)(nil).Find), userId, movieId)
}
