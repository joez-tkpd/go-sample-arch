// Code generated by MockGen. DO NOT EDIT.
// Source: ./user_handler.go

// Package httphandler is a generated GoMock package.
package httphandler

import (
	gomock "github.com/golang/mock/gomock"
	entity "github.com/joez-tkpd/go-sample-arch/entity"
	reflect "reflect"
)

// MockUserUsecase is a mock of UserUsecase interface
type MockUserUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUserUsecaseMockRecorder
}

// MockUserUsecaseMockRecorder is the mock recorder for MockUserUsecase
type MockUserUsecaseMockRecorder struct {
	mock *MockUserUsecase
}

// NewMockUserUsecase creates a new mock instance
func NewMockUserUsecase(ctrl *gomock.Controller) *MockUserUsecase {
	mock := &MockUserUsecase{ctrl: ctrl}
	mock.recorder = &MockUserUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserUsecase) EXPECT() *MockUserUsecaseMockRecorder {
	return m.recorder
}

// GetUser mocks base method
func (m *MockUserUsecase) GetUser(id int64) entity.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", id)
	ret0, _ := ret[0].(entity.User)
	return ret0
}

// GetUser indicates an expected call of GetUser
func (mr *MockUserUsecaseMockRecorder) GetUser(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserUsecase)(nil).GetUser), id)
}
