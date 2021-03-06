// Code generated by MockGen. DO NOT EDIT.
// Source: ./domain/repository/FileUpload.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	model "MyPIPE/domain/model"
	gomock "github.com/golang/mock/gomock"
	multipart "mime/multipart"
	reflect "reflect"
)

// MockFileUpload is a mock of FileUpload interface
type MockFileUpload struct {
	ctrl     *gomock.Controller
	recorder *MockFileUploadMockRecorder
}

// MockFileUploadMockRecorder is the mock recorder for MockFileUpload
type MockFileUploadMockRecorder struct {
	mock *MockFileUpload
}

// NewMockFileUpload creates a new mock instance
func NewMockFileUpload(ctrl *gomock.Controller) *MockFileUpload {
	mock := &MockFileUpload{ctrl: ctrl}
	mock.recorder = &MockFileUploadMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFileUpload) EXPECT() *MockFileUploadMockRecorder {
	return m.recorder
}

// Upload mocks base method
func (m *MockFileUpload) Upload(file multipart.File, movieFileHeader multipart.FileHeader, movieID model.MovieID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upload", file, movieFileHeader, movieID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upload indicates an expected call of Upload
func (mr *MockFileUploadMockRecorder) Upload(file, movieFileHeader, movieID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockFileUpload)(nil).Upload), file, movieFileHeader, movieID)
}
