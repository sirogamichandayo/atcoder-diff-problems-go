package usecase

import (
	"diff-problems/domain/client"
	"diff-problems/domain/repository"
)

type ProblemDifficultyInteractor struct {
	ProblemDifficultyRepository           repository.ProblemDifficultyRepository
	ProblemDifficultyAtCoderProblemClient client.ProblemDifficultyRepository
}

func (interactor *ProblemDifficultyInteractor) Update() (err error) {
	problemDifficultyList, err := interactor.ProblemDifficultyAtCoderProblemClient.Get()
	if err != nil {
		return
	}
	err = interactor.ProblemDifficultyRepository.BulkUpsert(problemDifficultyList)

	return
}
