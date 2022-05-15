package usecase

import (
	"diff-problems/domain/entity"
	"diff-problems/domain/repository"
)

type UserInteractor struct {
	UserSolveProblemDifficultySumRepository          repository.UserSolveProblemDifficultySumRepository
	UserSolveProblemDifficultySumUpdatedAtRepository repository.UserSolveProblemDifficultySumUpdatedAtRepository
}

func (interactor *UserInteractor) DiffRankById(userId string) (entity.UserSolveProblemDifficultySumAndRankAndUpdatedEpochTime, error) {
	userSum, err := interactor.UserSolveProblemDifficultySumRepository.FindById(userId)
	if err != nil {
		return entity.UserSolveProblemDifficultySumAndRankAndUpdatedEpochTime{}, err
	}
	diffSumList, err := interactor.UserSolveProblemDifficultySumRepository.All(userId)
	if err != nil {
		return entity.UserSolveProblemDifficultySumAndRankAndUpdatedEpochTime{}, err
	}
	updatedEpochTime, err := interactor.UserSolveProblemDifficultySumUpdatedAtRepository.Get()
	if err != nil {
		return entity.UserSolveProblemDifficultySumAndRankAndUpdatedEpochTime{}, err
	}
	rank := userSum.CalcRank(diffSumList)

	sumAndRank := entity.UserSolveProblemDifficultySumAndRank{Rank: rank, UserDifficultySum: userSum}
	sumAndRankAndEpochTime := entity.UserSolveProblemDifficultySumAndRankAndUpdatedEpochTime{UserSumAndRank: sumAndRank, UpdatedEpochTime: updatedEpochTime}

	return sumAndRankAndEpochTime, nil
}
