package usecase

import (
	"diff-problems/domain/entity"
	repository "diff-problems/domain/repository/mock"
	service "diff-problems/domain/repository/service/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Update_正常系(t *testing.T) {
	ctrl := gomock.NewController(t)

	userFirstAcSubmissionUpdatedAtRepositoryMock := repository.NewMockUserFirstAcSubmissionUpdatedAtRepository(ctrl)
	calcUserSolveProblemDifficultySumServiceMock := service.NewMockCalcUserSolveProblemDifficultySumService(ctrl)
	userSolveProblemDifficultySumRepositoryMock := repository.NewMockUserSolveProblemDifficultySumRepository(ctrl)
	userSolveProblemDifficultySumRepositoryUpdatedAtMock := repository.NewMockUserSolveProblemDifficultySumUpdatedAtRepository(ctrl)

	submissionUpdatedAt := int64(12345)
	diffSumList := entity.UserSolveProblemDifficultySumList{}

	gomock.InOrder(
		userFirstAcSubmissionUpdatedAtRepositoryMock.
			EXPECT().
			Get().
			Return(submissionUpdatedAt, nil).
			Times(1),
		calcUserSolveProblemDifficultySumServiceMock.
			EXPECT().
			Calc().
			Return(diffSumList, nil).Times(1),
		userSolveProblemDifficultySumRepositoryMock.
			EXPECT().
			Upsert(diffSumList).
			Return(nil).Times(1),
		userSolveProblemDifficultySumRepositoryUpdatedAtMock.
			EXPECT().
			Update(submissionUpdatedAt).
			Return(nil).
			Times(1),
	)

	interactor := UserSolveProblemDifficultySumInteractor{
		userFirstAcSubmissionUpdatedAtRepositoryMock,
		calcUserSolveProblemDifficultySumServiceMock,
		userSolveProblemDifficultySumRepositoryMock,
		userSolveProblemDifficultySumRepositoryUpdatedAtMock,
	}
	err := interactor.Update()
	assert.Nil(t, err)
}
