package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MakeValueForUpsertMySql(t *testing.T) {
	userId1 := "user1"
	diffSum1 := 11111.1
	rnk1 := uint64(1)
	userId2 := "user2"
	diffSum2 := 22222.2
	rnk2 := uint64(2)
	list := UserSolveProblemDifficultySumList{
		{userId1, diffSum1, rnk1},
		{userId2, diffSum2, rnk2},
	}

	placeholderValueList, err := list.MakeValuesForUpsertMySql()
	assert.Nil(t, err)
	placeholderValue := placeholderValueList[0]
	placeholder, values := placeholderValue.Placeholder(), placeholderValue.Values()

	assert.Equal(
		t,
		"(?,?,?),(?,?,?)",
		placeholder,
		[]interface{}{userId1, diffSum1, rnk1, userId2, diffSum2, rnk2},
	)
	assert.Equal(
		t,
		[]interface{}{userId1, diffSum1, rnk1, userId2, diffSum2, rnk2},
		values,
	)
}

func Test_Near(t *testing.T) {
	sum1 := UserSolveProblemDifficultySum{UserId: "user_id1", ClipDifficultySum: 55555.0, Rank: 1}
	sum2 := UserSolveProblemDifficultySum{UserId: "user_id2", ClipDifficultySum: 44444.0, Rank: 2}
	sum3 := UserSolveProblemDifficultySum{UserId: "user_id3", ClipDifficultySum: 33333.0, Rank: 3}
	sum4 := UserSolveProblemDifficultySum{UserId: "user_id4", ClipDifficultySum: 22222.0, Rank: 4}
	sum5 := UserSolveProblemDifficultySum{UserId: "user_id5", ClipDifficultySum: 11111.0, Rank: 5}

	list := UserSolveProblemDifficultySumList{sum1, sum2, sum3, sum4, sum5}
	sortedList, err := NewSortedUserSolveProblemDifficultySumList(list)
	assert.Nil(t, err)

	t.Run("元配列の範囲内", func(t *testing.T) {
		nearList, err := sortedList.Near("user_id3", 1)
		assert.Nil(t, err)
		assert.Equal(t, UserSolveProblemDifficultySumList{sum2, sum3, sum4}, nearList.List())
	})

	t.Run("元配列を下方向に超える", func(t *testing.T) {
		nearList, err := sortedList.Near("user_id1", 1)
		assert.Nil(t, err)
		assert.Equal(t, UserSolveProblemDifficultySumList{sum1, sum2}, nearList.List())
	})

	t.Run("元配列を上方向に超える", func(t *testing.T) {
		nearList, err := sortedList.Near("user_id5", 1)
		assert.Nil(t, err)
		assert.Equal(t, UserSolveProblemDifficultySumList{sum4, sum5}, nearList.List())
	})

	t.Run("元配列を上下方向に超える", func(t *testing.T) {
		nearList, err := sortedList.Near("user_id5", 100)
		assert.Nil(t, err)
		assert.Equal(t, UserSolveProblemDifficultySumList{sum1, sum2, sum3, sum4, sum5}, nearList.List())
	})

	t.Run("指定されたuserIdが存在しない", func(t *testing.T) {
		nearList, err := sortedList.Near("non user", 100)
		assert.Nil(t, nearList)
		assert.NotNil(t, err)
	})
}
