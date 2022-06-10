package usecase

import (
	"diff-problems/domain/client"
	"diff-problems/domain/repository"
	cqrsDto "diff-problems/usecase/cqrs_dto"
)

type UserProgressInteractor struct {
	ProblemDifficultyRepository repository.ProblemDifficultyRepository
	FirstAcSubmissionRepository repository.UserFirstAcSubmissionRepository
	ContestResultClient         client.ContestResultClient
}

func (interactor UserProgressInteractor) exec(userId string) (cqrsDto.UserRate, error) {
	// 
	return interactor.UserService.FindByUserId(userId)
}
