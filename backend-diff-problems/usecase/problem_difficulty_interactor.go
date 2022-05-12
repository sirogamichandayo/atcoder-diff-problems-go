package usecase

import (
	"diff-problems/domain/client"
	"diff-problems/domain/repository"
)

type ProblemDifficultyInteractor struct {
	ProblemDifficultyRepository           repository.ProblemDifficultyRepository
	ProblemDifficultyAtCoderProblemClient client.ProblemDifficultyAtCoderProblemClient
}

func (interactor *ProblemDifficultyInteractor) Update() (err error) {
	problemDifficultyList, err := interactor.ProblemDifficultyAtCoderProblemClient.Fetch()
	if err != nil {
		return
	}
	err = interactor.ProblemDifficultyRepository.BulkUpsert(problemDifficultyList)

	return
}
