package impl

import "diff-problems/domain/entity"

type UserProgressMaker interface {
	Make(entity.ProblemDifficulty, entity.AcUserSubmissionList) entity.UserProgressList
}
