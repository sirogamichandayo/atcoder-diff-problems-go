package usecase

import (
	"diff-problems/domain/entity"
	"diff-problems/domain/repository"
)

type RankInteractor struct {
	userSolveProblemDifficultySumRepository repository.UserSolveProblemDifficultySumRepository
}

func (interactor RankInteractor) Near(
	userId string,
	nearCnt int,
) (*entity.SortedUserSolveProblemDifficultySumList, error) {
	sortedList, err := interactor.userSolveProblemDifficultySumRepository.SortedAll()
	if err != nil {
		return nil, err
	}

	return sortedList.Near(userId, nearCnt)
}
