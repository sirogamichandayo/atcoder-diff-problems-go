// Code generated by MockGen. DO NOT EDIT.
// Source: ../domain/client/user_submission_client.go

// Package mock is a generated GoMock package.
package mock

import (
	entity "diff-problems/domain/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserSubmissionClient is a mock of UserSubmissionClient interface.
type MockUserSubmissionClient struct {
	ctrl     *gomock.Controller
	recorder *MockUserSubmissionClientMockRecorder
}

// MockUserSubmissionClientMockRecorder is the mock recorder for MockUserSubmissionClient.
type MockUserSubmissionClientMockRecorder struct {
	mock *MockUserSubmissionClient
}

// NewMockUserSubmissionClient creates a new mock instance.
func NewMockUserSubmissionClient(ctrl *gomock.Controller) *MockUserSubmissionClient {
	mock := &MockUserSubmissionClient{ctrl: ctrl}
	mock.recorder = &MockUserSubmissionClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserSubmissionClient) EXPECT() *MockUserSubmissionClientMockRecorder {
	return m.recorder
}

// FetchSinceByEpochTime mocks base method.
func (m *MockUserSubmissionClient) FetchSinceByEpochTime(arg0 int64) (entity.UserSubmissionList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchSinceByEpochTime", arg0)
	ret0, _ := ret[0].(entity.UserSubmissionList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchSinceByEpochTime indicates an expected call of FetchSinceByEpochTime.
func (mr *MockUserSubmissionClientMockRecorder) FetchSinceByEpochTime(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchSinceByEpochTime", reflect.TypeOf((*MockUserSubmissionClient)(nil).FetchSinceByEpochTime), arg0)
}
