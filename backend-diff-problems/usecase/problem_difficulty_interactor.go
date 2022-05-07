package usecase

import "diff-problems/domain/entity"

type ProblemDifficultyInteractor struct {
	ProblemDifficultyRepository           ProblemDifficultyRepository
	ProblemDifficultyAtcoderProblemClient ProblemDifficultyAtcoderProblemClient
}

func (interactor *ProblemDifficultyInteractor) FetchAndStore() (err error) {
	problemDifficultyList, err := interactor.ProblemDifficultyAtcoderProblemClient.Fetch()
	if err != nil {
		return
	}
	err = interactor.ProblemDifficultyRepository.BulkUpsert(problemDifficultyList)

	return
}

type ProblemDifficultyRepository interface {
	BulkUpsert(u entity.ProblemDifficultyList) (err error)
}

type ProblemDifficultyAtcoderProblemClient interface {
	Fetch() (list entity.ProblemDifficultyList, err error)
}
