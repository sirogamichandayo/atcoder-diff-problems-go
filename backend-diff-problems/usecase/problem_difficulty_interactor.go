package usecase

import (
	"diff-problems/domain/client"
	"diff-problems/domain/repository"
)

type ProblemDifficultyInteractor struct {
	ProblemDifficultyRepository           repository.ProblemDifficultyRepository
	ProblemDifficultyAtCoderProblemClient client.ProblemDifficultyClient
}

func (interactor *ProblemDifficultyInteractor) Update() error {
	fetchedProblemDifficultyList, err := interactor.ProblemDifficultyAtCoderProblemClient.Fetch()
	if err != nil {
		return err
	}

	return interactor.ProblemDifficultyRepository.BulkUpsert(fetchedProblemDifficultyList)
}
