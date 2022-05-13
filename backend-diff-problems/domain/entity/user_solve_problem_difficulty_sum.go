package entity

import "strings"

type UserSolveProblemDifficultySum struct {
	UserId            string
	ClipDifficultySum float64
}

type UserSolveProblemDifficultySumList []UserSolveProblemDifficultySum

func (list UserSolveProblemDifficultySumList) MakeValueForUpsertMySql() (string, []interface{}) {
	listSize := len(list)
	placeholders := make([]string, 0, listSize)
	for i := 0; i < listSize; i++ {
		placeholders = append(placeholders, "(?,?)")
	}

	values := make([]interface{}, 0, listSize*2)
	for _, entity := range list {
		values = append(values, entity.UserId, entity.ClipDifficultySum)
	}
	return strings.Join(placeholders, ","), values
}
