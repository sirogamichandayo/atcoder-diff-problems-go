package usecase

import (
	cqrsDto "diff-problems/usecase/cqrs_dto"
	cqrsService "diff-problems/usecase/cqrs_service"
)

type UserInteractor struct {
	UserService cqrsService.UserService
}

func (interactor UserInteractor) FindByUserId(userId string) (cqrsDto.User, error) {
	return interactor.UserService.FindByUserId(userId)
}
