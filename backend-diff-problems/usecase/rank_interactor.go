package usecase

import (
	"diff-problems/domain/entity"
	"diff-problems/domain/repository"
)

type RankInteractor struct {
	UserSolveProblemDifficultySumRepository repository.UserSolveProblemDifficultySumRepository
}

func (interactor RankInteractor) Near(
	userId string,
	nearCnt int,
) (*entity.SortedUserSolveProblemDifficultySumList, error) {
	sortedList, err := interactor.UserSolveProblemDifficultySumRepository.SortedAll()
	if err != nil {
		return nil, err
	}

	return sortedList.Near(userId, nearCnt)
}
