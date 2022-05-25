// Code generated by MockGen. DO NOT EDIT.
// Source: user_solve_problem_difficulty_sum_updated_at_repository.go

// Package repository is a generated GoMock package.
package repository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserSolveProblemDifficultySumUpdatedAtRepository is a mock of UserSolveProblemDifficultySumUpdatedAtRepository interface.
type MockUserSolveProblemDifficultySumUpdatedAtRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserSolveProblemDifficultySumUpdatedAtRepositoryMockRecorder
}

// MockUserSolveProblemDifficultySumUpdatedAtRepositoryMockRecorder is the mock recorder for MockUserSolveProblemDifficultySumUpdatedAtRepository.
type MockUserSolveProblemDifficultySumUpdatedAtRepositoryMockRecorder struct {
	mock *MockUserSolveProblemDifficultySumUpdatedAtRepository
}

// NewMockUserSolveProblemDifficultySumUpdatedAtRepository creates a new mock instance.
func NewMockUserSolveProblemDifficultySumUpdatedAtRepository(ctrl *gomock.Controller) *MockUserSolveProblemDifficultySumUpdatedAtRepository {
	mock := &MockUserSolveProblemDifficultySumUpdatedAtRepository{ctrl: ctrl}
	mock.recorder = &MockUserSolveProblemDifficultySumUpdatedAtRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserSolveProblemDifficultySumUpdatedAtRepository) EXPECT() *MockUserSolveProblemDifficultySumUpdatedAtRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockUserSolveProblemDifficultySumUpdatedAtRepository) Get() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockUserSolveProblemDifficultySumUpdatedAtRepositoryMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserSolveProblemDifficultySumUpdatedAtRepository)(nil).Get))
}

// Update mocks base method.
func (m *MockUserSolveProblemDifficultySumUpdatedAtRepository) Update(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserSolveProblemDifficultySumUpdatedAtRepositoryMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserSolveProblemDifficultySumUpdatedAtRepository)(nil).Update), arg0)
}