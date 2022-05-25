// Code generated by MockGen. DO NOT EDIT.
// Source: request_handler.go

// Package api is a generated GoMock package.
package api

import (
	api "diff-problems/interfaces/api"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRequestHandler is a mock of RequestHandler interface.
type MockRequestHandler struct {
	ctrl     *gomock.Controller
	recorder *MockRequestHandlerMockRecorder
}

// MockRequestHandlerMockRecorder is the mock recorder for MockRequestHandler.
type MockRequestHandlerMockRecorder struct {
	mock *MockRequestHandler
}

// NewMockRequestHandler creates a new mock instance.
func NewMockRequestHandler(ctrl *gomock.Controller) *MockRequestHandler {
	mock := &MockRequestHandler{ctrl: ctrl}
	mock.recorder = &MockRequestHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRequestHandler) EXPECT() *MockRequestHandlerMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockRequestHandler) Get(url string, headers map[string]string) (api.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", url, headers)
	ret0, _ := ret[0].(api.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRequestHandlerMockRecorder) Get(url, headers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRequestHandler)(nil).Get), url, headers)
}

// MockResponse is a mock of Response interface.
type MockResponse struct {
	ctrl     *gomock.Controller
	recorder *MockResponseMockRecorder
}

// MockResponseMockRecorder is the mock recorder for MockResponse.
type MockResponseMockRecorder struct {
	mock *MockResponse
}

// NewMockResponse creates a new mock instance.
func NewMockResponse(ctrl *gomock.Controller) *MockResponse {
	mock := &MockResponse{ctrl: ctrl}
	mock.recorder = &MockResponseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResponse) EXPECT() *MockResponseMockRecorder {
	return m.recorder
}

// BodyBytes mocks base method.
func (m *MockResponse) BodyBytes() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BodyBytes")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// BodyBytes indicates an expected call of BodyBytes.
func (mr *MockResponseMockRecorder) BodyBytes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BodyBytes", reflect.TypeOf((*MockResponse)(nil).BodyBytes))
}

// IsSuccess mocks base method.
func (m *MockResponse) IsSuccess() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSuccess")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSuccess indicates an expected call of IsSuccess.
func (mr *MockResponseMockRecorder) IsSuccess() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSuccess", reflect.TypeOf((*MockResponse)(nil).IsSuccess))
}