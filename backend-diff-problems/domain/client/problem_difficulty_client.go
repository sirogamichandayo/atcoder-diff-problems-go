package client

import "diff-problems/domain/entity"

type ProblemDifficultyClient interface {
	Fetch() (list entity.ProblemDifficultyList, err error)
}
