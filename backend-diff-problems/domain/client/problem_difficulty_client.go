package client

import "diff-problems/domain/entity"

type ProblemDifficultyAtCoderProblemClient interface {
	Fetch() (list entity.ProblemDifficultyList, err error)
}
