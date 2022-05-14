package entity

import (
	"diff-problems/domain/vo"
	"strings"
)

type UserSolveProblemDifficultySum struct {
	UserId            string
	ClipDifficultySum float64
}

type UserSolveProblemDifficultySumList []UserSolveProblemDifficultySum

const UserSolveProblemDifficultySumListPluckLen = 500

func (list UserSolveProblemDifficultySumList) MakeValuesForUpsertMySql() (vo.PlaceholderValueList, error) {
	lLen := len(list)
	pLen := (lLen + (UserSolveProblemDifficultySumListPluckLen - 1)) / UserSolveProblemDifficultySumListPluckLen

	res := make(vo.PlaceholderValueList, 0, pLen)
	for i := 0; i < pLen; i++ {
		begin := i * UserSolveProblemDifficultySumListPluckLen
		end := (i + 1) * UserSolveProblemDifficultySumListPluckLen
		if end > lLen {
			end = lLen
		}

		listSize := end - begin
		placeholders := make([]string, 0, listSize)
		for i := 0; i < listSize; i++ {
			placeholders = append(placeholders, "(?,?)")
		}

		values := make([]interface{}, 0, listSize*2)
		for _, entity := range list[begin:end] {
			values = append(values, entity.UserId, entity.ClipDifficultySum)
		}

		pv, err := vo.NewPlaceHolderValue(strings.Join(placeholders, ","), values)
		if err != nil {
			return nil, err
		}
		res = append(res, pv)
	}

	return res, nil
}
