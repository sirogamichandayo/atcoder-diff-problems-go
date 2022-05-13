package database

import (
	"diff-problems/domain/entity"
)

type UserSolveProblemDifficultySumRepository struct {
	SqlHandler
}

func (u UserSolveProblemDifficultySumRepository) Upsert(list entity.UserSolveProblemDifficultySumList) error {
	//TODO implement me
	panic("implement me")
}
