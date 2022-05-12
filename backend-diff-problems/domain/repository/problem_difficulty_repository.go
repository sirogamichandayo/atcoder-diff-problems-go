package repository

import "diff-problems/domain/entity"

type ProblemDifficultyRepository interface {
	BulkUpsert(entity.ProblemDifficultyList) error
}
