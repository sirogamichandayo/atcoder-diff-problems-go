package usecase

import (
	cqrsDto "diff-problems/usecase/cqrs_dto"
	cqrsService "diff-problems/usecase/cqrs_service"
)

type UserProblemInteractor struct {
	UserProblemService cqrsService.UserProblemService
}

func (interactor UserProblemInteractor) FindByUserId(userId string) (cqrsDto.UserProblem, error) {
	return interactor.UserProblemService.FindByUserId(userId)
}
