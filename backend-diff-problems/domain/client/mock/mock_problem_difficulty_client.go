// Code generated by MockGen. DO NOT EDIT.
// Source: problem_difficulty_client.go

// Package client is a generated GoMock package.
package client

import (
	entity "diff-problems/domain/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockProblemDifficultyClient is a mock of ProblemDifficultyClient interface.
type MockProblemDifficultyClient struct {
	ctrl     *gomock.Controller
	recorder *MockProblemDifficultyClientMockRecorder
}

// MockProblemDifficultyClientMockRecorder is the mock recorder for MockProblemDifficultyClient.
type MockProblemDifficultyClientMockRecorder struct {
	mock *MockProblemDifficultyClient
}

// NewMockProblemDifficultyClient creates a new mock instance.
func NewMockProblemDifficultyClient(ctrl *gomock.Controller) *MockProblemDifficultyClient {
	mock := &MockProblemDifficultyClient{ctrl: ctrl}
	mock.recorder = &MockProblemDifficultyClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProblemDifficultyClient) EXPECT() *MockProblemDifficultyClientMockRecorder {
	return m.recorder
}

// Fetch mocks base method.
func (m *MockProblemDifficultyClient) Fetch() (entity.ProblemDifficultyList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch")
	ret0, _ := ret[0].(entity.ProblemDifficultyList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch.
func (mr *MockProblemDifficultyClientMockRecorder) Fetch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockProblemDifficultyClient)(nil).Fetch))
}