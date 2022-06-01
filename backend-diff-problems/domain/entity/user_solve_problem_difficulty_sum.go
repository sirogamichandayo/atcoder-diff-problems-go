package entity

import (
	"diff-problems/domain/vo"
	"fmt"
	"strings"
)

type UserSolveProblemDifficultySum struct {
	UserId            string
	ClipDifficultySum float64
	Rank              uint64
}

type UserSolveProblemDifficultySumList []UserSolveProblemDifficultySum

const UserSolveProblemDifficultySumListPluckLen = 333

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
			placeholders = append(placeholders, "(?,?,?)")
		}

		values := make([]interface{}, 0, listSize*3)
		for _, entity := range list[begin:end] {
			values = append(values, entity.UserId, entity.ClipDifficultySum, entity.Rank)
		}

		pv, err := vo.NewPlaceHolderValue(strings.Join(placeholders, ","), values)
		if err != nil {
			return nil, err
		}
		res = append(res, pv)
	}

	return res, nil
}

type SortedUserSolveProblemDifficultySumList struct {
	list UserSolveProblemDifficultySumList
}

func NewSortedUserSolveProblemDifficultySumList(list UserSolveProblemDifficultySumList) (*SortedUserSolveProblemDifficultySumList, error) {
	rnk := uint64(0)
	for _, entity := range list {
		if rnk > entity.Rank {
			return nil, fmt.Errorf("not sorted")
		}
	}
	return &SortedUserSolveProblemDifficultySumList{list}, nil
}

func (list SortedUserSolveProblemDifficultySumList) List() UserSolveProblemDifficultySumList {
	return list.list
}

func (list SortedUserSolveProblemDifficultySumList) Near(userId string, nearCnt int) (*SortedUserSolveProblemDifficultySumList, error) {
	userInd := -1
	for ind, entity := range list.list {
		if entity.UserId == userId {
			userInd = ind
		}
	}
	if userInd == -1 {
		return nil, fmt.Errorf("not found userId %s", userId)
	}

	begin := userInd - nearCnt
	if begin < 0 {
		begin = 0
	}
	end := userInd + nearCnt + 1
	if end > len(list.list) {
		end = len(list.list)
	}

	return NewSortedUserSolveProblemDifficultySumList(list.list[begin:end])
}
