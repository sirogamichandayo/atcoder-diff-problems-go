// Code generated by MockGen. DO NOT EDIT.
// Source: context.go

// Package controllers is a generated GoMock package.
package controllers

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockContext is a mock of Context interface.
type MockContext struct {
	ctrl     *gomock.Controller
	recorder *MockContextMockRecorder
}

// MockContextMockRecorder is the mock recorder for MockContext.
type MockContextMockRecorder struct {
	mock *MockContext
}

// NewMockContext creates a new mock instance.
func NewMockContext(ctrl *gomock.Controller) *MockContext {
	mock := &MockContext{ctrl: ctrl}
	mock.recorder = &MockContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContext) EXPECT() *MockContextMockRecorder {
	return m.recorder
}

// Bind mocks base method.
func (m *MockContext) Bind(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bind", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Bind indicates an expected call of Bind.
func (mr *MockContextMockRecorder) Bind(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bind", reflect.TypeOf((*MockContext)(nil).Bind), arg0)
}

// DefaultQuery mocks base method.
func (m *MockContext) DefaultQuery(arg0, arg1 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultQuery", arg0, arg1)
	ret0, _ := ret[0].(string)
	return ret0
}

// DefaultQuery indicates an expected call of DefaultQuery.
func (mr *MockContextMockRecorder) DefaultQuery(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultQuery", reflect.TypeOf((*MockContext)(nil).DefaultQuery), arg0, arg1)
}

// JSON mocks base method.
func (m *MockContext) JSON(arg0 int, arg1 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "JSON", arg0, arg1)
}

// JSON indicates an expected call of JSON.
func (mr *MockContextMockRecorder) JSON(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JSON", reflect.TypeOf((*MockContext)(nil).JSON), arg0, arg1)
}

// Param mocks base method.
func (m *MockContext) Param(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Param", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Param indicates an expected call of Param.
func (mr *MockContextMockRecorder) Param(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Param", reflect.TypeOf((*MockContext)(nil).Param), arg0)
}

// Status mocks base method.
func (m *MockContext) Status(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Status", arg0)
}

// Status indicates an expected call of Status.
func (mr *MockContextMockRecorder) Status(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockContext)(nil).Status), arg0)
}
