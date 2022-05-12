package usecase

import "diff-problems/domain/entity"

type ProblemDifficultyInteractor struct {
	ProblemDifficultyRepository           ProblemDifficultyRepository
	ProblemDifficultyAtCoderProblemClient ProblemDifficultyAtCoderProblemClient
}

func (interactor *ProblemDifficultyInteractor) Update() (err error) {
	problemDifficultyList, err := interactor.ProblemDifficultyAtCoderProblemClient.Fetch()
	if err != nil {
		return
	}
	err = interactor.ProblemDifficultyRepository.BulkUpsert(problemDifficultyList)

	return
}

type ProblemDifficultyRepository interface {
	BulkUpsert(u entity.ProblemDifficultyList) (err error)
}

type ProblemDifficultyAtCoderProblemClient interface {
	Fetch() (list entity.ProblemDifficultyList, err error)
}
