package repository

import "diff-problems/domain/entity"

type ProblemDifficultyRepository interface {
	BulkUpsert(u entity.ProblemDifficultyList) (err error)
}
