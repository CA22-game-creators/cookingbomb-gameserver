// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	reflect "reflect"

	domain "github.com/CA22-game-creators/cookingbomb-gameserver/domain/model/account"
	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Find mocks base method.
func (m *MockRepository) Find(sessionToken string) (domain.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", sessionToken)
	ret0, _ := ret[0].(domain.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockRepositoryMockRecorder) Find(sessionToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockRepository)(nil).Find), sessionToken)
}