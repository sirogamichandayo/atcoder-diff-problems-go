package usecase

import (
	"diff-problems/domain/client"
	"diff-problems/domain/vo"
)

type UserRateInteractor struct {
	ContestResultClient client.ContestRequestClient
}

func (interactor UserRateInteractor) ShowLatest(userId string) (vo.ContestResult, error) {
	list, err := interactor.ContestResultClient.All(userId)
	if err != nil {
		return vo.ContestResult{}, err
	}
	return list.Last()
}

func (interactor UserRateInteractor) Index(userId string) (vo.ContestResultList, error) {
	list, err := interactor.ContestResultClient.All(userId)
	return list, err
}
