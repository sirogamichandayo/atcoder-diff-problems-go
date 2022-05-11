package usecase

import (
	"diff-problems/domain/entity"
	"diff-problems/domain/vo"
	mockUsecase "diff-problems/usecase/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_UpdateAll_正常系(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userFirstAcSubmissionRepositoryMock :=
		mockUsecase.NewMockUserFirstAcSubmissionRepository(ctrl)
	UserFirstAcSubmissionUpdatedAtRepositoryMock :=
		mockUsecase.NewMockUserFirstAcSubmissionUpdatedAtRepository(ctrl)
	UserSubmissionAtCoderProblemClientMock :=
		mockUsecase.NewMockUserSubmissionAtCoderProblemClient(ctrl)

	userId1 := "user_id1"
	problemId1 := "problem_id1"
	result1 := vo.Ac
	epochTime1 := int64(11111)
	userId2 := "user_id2"
	problemId2 := "problem_id2"
	result2 := vo.Wa
	epochTime2 := int64(22222)

	submissionList := entity.UserSubmissionList{
		entity.UserSubmission{UserId: userId1, ProblemId: problemId1, Result: result1, EpochTime: epochTime1},
		entity.UserSubmission{UserId: userId2, ProblemId: problemId2, Result: result2, EpochTime: epochTime2},
	}
	expectAcList := entity.AcUserSubmissionList{
		entity.AcUserSubmission{UserId: userId1, ProblemId: problemId1, EpochTime: epochTime1},
	}
	emptyUserSubmissionList := entity.UserSubmissionList{}

	gomock.InOrder(
		UserSubmissionAtCoderProblemClientMock.
			EXPECT().
			Fetch(gomock.Eq(int64(1))).
			Times(1).
			Return(submissionList, nil),
		userFirstAcSubmissionRepositoryMock.
			EXPECT().
			BulkUpsert(gomock.Eq(expectAcList)).
			Times(1).
			Return(nil),
		UserFirstAcSubmissionUpdatedAtRepositoryMock.
			EXPECT().
			Update(gomock.Eq(epochTime2)).
			Times(1).
			Return(nil),
		UserSubmissionAtCoderProblemClientMock.
			EXPECT().
			Fetch(gomock.Eq(epochTime2+1)).
			Times(1).
			Return(emptyUserSubmissionList, nil),
	)

	interactor := UserFirstAcSubmissionInteractor{
		UserFirstAcSubmissionRepository:          userFirstAcSubmissionRepositoryMock,
		UserFirstAcSubmissionUpdatedAtRepository: UserFirstAcSubmissionUpdatedAtRepositoryMock,
		UserSubmissionAtCoderProblemClient:       UserSubmissionAtCoderProblemClientMock,
	}

	err := interactor.UpdateAll()
	assert.Nil(t, err)
}

func Test_UpdateFromUpdatedAt_正常系(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userFirstAcSubmissionRepositoryMock :=
		mockUsecase.NewMockUserFirstAcSubmissionRepository(ctrl)
	UserFirstAcSubmissionUpdatedAtRepositoryMock :=
		mockUsecase.NewMockUserFirstAcSubmissionUpdatedAtRepository(ctrl)
	UserSubmissionAtCoderProblemClientMock :=
		mockUsecase.NewMockUserSubmissionAtCoderProblemClient(ctrl)

	updatedTime := int64(5000)
	userId1 := "user_id1"
	problemId1 := "problem_id1"
	result1 := vo.Ac
	epochTime1 := int64(11111)
	userId2 := "user_id2"
	problemId2 := "problem_id2"
	result2 := vo.Wa
	epochTime2 := int64(22222)

	submissionList := entity.UserSubmissionList{
		entity.UserSubmission{UserId: userId1, ProblemId: problemId1, Result: result1, EpochTime: epochTime1},
		entity.UserSubmission{UserId: userId2, ProblemId: problemId2, Result: result2, EpochTime: epochTime2},
	}
	expectAcList := entity.AcUserSubmissionList{
		entity.AcUserSubmission{UserId: userId1, ProblemId: problemId1, EpochTime: epochTime1},
	}
	emptyUserSubmissionList := entity.UserSubmissionList{}

	gomock.InOrder(
		UserFirstAcSubmissionUpdatedAtRepositoryMock.
			EXPECT().
			Get().
			Times(1).
			Return(updatedTime, nil),
		UserSubmissionAtCoderProblemClientMock.
			EXPECT().
			Fetch(gomock.Eq(updatedTime+1)).
			Times(1).
			Return(submissionList, nil),
		userFirstAcSubmissionRepositoryMock.
			EXPECT().
			BulkUpsert(gomock.Eq(expectAcList)).
			Times(1).
			Return(nil),
		UserFirstAcSubmissionUpdatedAtRepositoryMock.
			EXPECT().
			Update(gomock.Eq(epochTime2)).
			Times(1).
			Return(nil),
		UserSubmissionAtCoderProblemClientMock.
			EXPECT().
			Fetch(gomock.Eq(epochTime2+1)).
			Times(1).
			Return(emptyUserSubmissionList, nil),
	)

	interactor := UserFirstAcSubmissionInteractor{
		UserFirstAcSubmissionRepository:          userFirstAcSubmissionRepositoryMock,
		UserFirstAcSubmissionUpdatedAtRepository: UserFirstAcSubmissionUpdatedAtRepositoryMock,
		UserSubmissionAtCoderProblemClient:       UserSubmissionAtCoderProblemClientMock,
	}

	err := interactor.UpdateFromUpdatedAt()
	assert.Nil(t, err)
}
