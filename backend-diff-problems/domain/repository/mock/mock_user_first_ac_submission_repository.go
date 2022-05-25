// Code generated by MockGen. DO NOT EDIT.
// Source: user_first_ac_submission_repository.go

// Package repository is a generated GoMock package.
package repository

import (
	entity "diff-problems/domain/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserFirstAcSubmissionRepository is a mock of UserFirstAcSubmissionRepository interface.
type MockUserFirstAcSubmissionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserFirstAcSubmissionRepositoryMockRecorder
}

// MockUserFirstAcSubmissionRepositoryMockRecorder is the mock recorder for MockUserFirstAcSubmissionRepository.
type MockUserFirstAcSubmissionRepositoryMockRecorder struct {
	mock *MockUserFirstAcSubmissionRepository
}

// NewMockUserFirstAcSubmissionRepository creates a new mock instance.
func NewMockUserFirstAcSubmissionRepository(ctrl *gomock.Controller) *MockUserFirstAcSubmissionRepository {
	mock := &MockUserFirstAcSubmissionRepository{ctrl: ctrl}
	mock.recorder = &MockUserFirstAcSubmissionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserFirstAcSubmissionRepository) EXPECT() *MockUserFirstAcSubmissionRepositoryMockRecorder {
	return m.recorder
}

// BulkUpsert mocks base method.
func (m *MockUserFirstAcSubmissionRepository) BulkUpsert(arg0 entity.AcUserSubmissionList) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkUpsert", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkUpsert indicates an expected call of BulkUpsert.
func (mr *MockUserFirstAcSubmissionRepositoryMockRecorder) BulkUpsert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkUpsert", reflect.TypeOf((*MockUserFirstAcSubmissionRepository)(nil).BulkUpsert), arg0)
}