// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	models "github.com/giicoo/golang_CRUD/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockAuthorServices is a mock of AuthorServices interface.
type MockAuthorServices struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorServicesMockRecorder
}

// MockAuthorServicesMockRecorder is the mock recorder for MockAuthorServices.
type MockAuthorServicesMockRecorder struct {
	mock *MockAuthorServices
}

// NewMockAuthorServices creates a new mock instance.
func NewMockAuthorServices(ctrl *gomock.Controller) *MockAuthorServices {
	mock := &MockAuthorServices{ctrl: ctrl}
	mock.recorder = &MockAuthorServicesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorServices) EXPECT() *MockAuthorServicesMockRecorder {
	return m.recorder
}

// DeleteAuthor mocks base method.
func (m *MockAuthorServices) DeleteAuthor(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAuthor", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAuthor indicates an expected call of DeleteAuthor.
func (mr *MockAuthorServicesMockRecorder) DeleteAuthor(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAuthor", reflect.TypeOf((*MockAuthorServices)(nil).DeleteAuthor), id)
}

// GetAuthor mocks base method.
func (m *MockAuthorServices) GetAuthor(au models.Author) ([]models.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthor", au)
	ret0, _ := ret[0].([]models.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthor indicates an expected call of GetAuthor.
func (mr *MockAuthorServicesMockRecorder) GetAuthor(au interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthor", reflect.TypeOf((*MockAuthorServices)(nil).GetAuthor), au)
}

// NewAuthor mocks base method.
func (m *MockAuthorServices) NewAuthor(name string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAuthor", name)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewAuthor indicates an expected call of NewAuthor.
func (mr *MockAuthorServicesMockRecorder) NewAuthor(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAuthor", reflect.TypeOf((*MockAuthorServices)(nil).NewAuthor), name)
}
