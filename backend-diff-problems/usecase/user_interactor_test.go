package usecase

import (
	"diff-problems/domain/entity"
	repository "diff-problems/domain/repository/mock"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

func Test_正常系_DiffRankById(t *testing.T) {
	ctrl := gomock.NewController(t)

	sumRepositoryMock := repository.NewMockUserSolveProblemDifficultySumRepository(ctrl)
	sumUpdatedAtRepositoryMock := repository.NewMockUserFirstAcSubmissionUpdatedAtRepository(ctrl)
	userId := "userId1"
	diffSum := 100.0
	rank := uint64(1)
	userSum := entity.UserSolveProblemDifficultySum{UserId: userId, ClipDifficultySum: diffSum, Rank: rank}
	updatedEpochTime := int64(12345)

	gomock.InOrder(
		sumRepositoryMock.EXPECT().
			FindById(userId).
			Times(1).
			Return(userSum, nil),
		sumUpdatedAtRepositoryMock.EXPECT().
			Get().
			Times(1).
			Return(updatedEpochTime, nil),
	)

	interactor := UserInteractor{
		UserSolveProblemDifficultySumRepository:          sumRepositoryMock,
		UserSolveProblemDifficultySumUpdatedAtRepository: sumUpdatedAtRepositoryMock,
	}

	output, err := interactor.DiffRankById(userId)
	assert.Nil(t, err)
	assert.Equal(t, updatedEpochTime, output.RankUpdatedEpochTime)
	assert.Equal(t, userId, output.UserSum.UserId)
	assert.Equal(t, diffSum, output.UserSum.ClipDifficultySum)
	assert.Equal(t, rank, output.UserSum.Rank)
}
