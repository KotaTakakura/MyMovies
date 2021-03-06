// Code generated by MockGen. DO NOT EDIT.
// Source: ./domain/queryService/GetComment.go

// Package mock_queryService is a generated GoMock package.
package mock_queryService

import (
	model "MyPIPE/domain/model"
	queryService "MyPIPE/domain/queryService"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCommentQueryService is a mock of CommentQueryService interface
type MockCommentQueryService struct {
	ctrl     *gomock.Controller
	recorder *MockCommentQueryServiceMockRecorder
}

// MockCommentQueryServiceMockRecorder is the mock recorder for MockCommentQueryService
type MockCommentQueryServiceMockRecorder struct {
	mock *MockCommentQueryService
}

// NewMockCommentQueryService creates a new mock instance
func NewMockCommentQueryService(ctrl *gomock.Controller) *MockCommentQueryService {
	mock := &MockCommentQueryService{ctrl: ctrl}
	mock.recorder = &MockCommentQueryServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCommentQueryService) EXPECT() *MockCommentQueryServiceMockRecorder {
	return m.recorder
}

// FindByMovieId mocks base method
func (m *MockCommentQueryService) FindByMovieId(movieId model.MovieID) queryService.FindByMovieIdDTO {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByMovieId", movieId)
	ret0, _ := ret[0].(queryService.FindByMovieIdDTO)
	return ret0
}

// FindByMovieId indicates an expected call of FindByMovieId
func (mr *MockCommentQueryServiceMockRecorder) FindByMovieId(movieId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByMovieId", reflect.TypeOf((*MockCommentQueryService)(nil).FindByMovieId), movieId)
}
