package cqrsDto

import (
	"diff-problems/domain/entity"
)

type UserProblem struct {
	userId              string
	solvedProblemList   entity.ProblemDifficultyList
	allProblemList      entity.ProblemDifficultyList
	clipDifficultyTotal float64
	updatedEpochTime    int64
}

func NewUserProblem(
	userId string,
	solvedProblemList entity.ProblemDifficultyList,
	allProblemList entity.ProblemDifficultyList,
	updatedEpochTime int64,
) UserProblem {
	return UserProblem{
		userId:            userId,
		solvedProblemList: solvedProblemList,
		allProblemList:    allProblemList,
		updatedEpochTime:  updatedEpochTime,
	}
}

func (p UserProblem) UserId() string                                  { return p.userId }
func (p UserProblem) SolvedProblemList() entity.ProblemDifficultyList { return p.solvedProblemList }
func (p UserProblem) AllProblemList() entity.ProblemDifficultyList    { return p.allProblemList }
func (p UserProblem) UpdatedEpochTime() int64                         { return p.updatedEpochTime }
