// Code generated by MockGen. DO NOT EDIT.
// Source: input_port.go

// Package mock_application is a generated GoMock package.
package mock_application

import (
	reflect "reflect"

	application "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/application/get_connection_status"
	gomock "github.com/golang/mock/gomock"
)

// MockInputPort is a mock of InputPort interface.
type MockInputPort struct {
	ctrl     *gomock.Controller
	recorder *MockInputPortMockRecorder
}

// MockInputPortMockRecorder is the mock recorder for MockInputPort.
type MockInputPortMockRecorder struct {
	mock *MockInputPort
}

// NewMockInputPort creates a new mock instance.
func NewMockInputPort(ctrl *gomock.Controller) *MockInputPort {
	mock := &MockInputPort{ctrl: ctrl}
	mock.recorder = &MockInputPortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInputPort) EXPECT() *MockInputPortMockRecorder {
	return m.recorder
}

// Handle mocks base method.
func (m *MockInputPort) Handle(arg0 application.InputData) application.OutputData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle", arg0)
	ret0, _ := ret[0].(application.OutputData)
	return ret0
}

// Handle indicates an expected call of Handle.
func (mr *MockInputPortMockRecorder) Handle(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockInputPort)(nil).Handle), arg0)
}
