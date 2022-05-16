package usecase

import (
	"diff-problems/domain/entity"
	"diff-problems/domain/repository"
)

type UserInteractor struct {
	UserSolveProblemDifficultySumRepository          repository.UserSolveProblemDifficultySumRepository
	UserSolveProblemDifficultySumUpdatedAtRepository repository.UserSolveProblemDifficultySumUpdatedAtRepository
}

func (interactor UserInteractor) DiffRankById(userId string) (entity.UserSolveProblemDifficultySumAndUpdatedEpochTime, error) {
	userSum, err := interactor.UserSolveProblemDifficultySumRepository.FindById(userId)
	if err != nil {
		return entity.UserSolveProblemDifficultySumAndUpdatedEpochTime{}, err
	}
	updatedEpochTime, err := interactor.UserSolveProblemDifficultySumUpdatedAtRepository.Get()
	if err != nil {
		return entity.UserSolveProblemDifficultySumAndUpdatedEpochTime{}, err
	}

	sumAndRankAndEpochTime := entity.UserSolveProblemDifficultySumAndUpdatedEpochTime{UserSum: userSum, RankUpdatedEpochTime: updatedEpochTime}

	return sumAndRankAndEpochTime, nil
}
