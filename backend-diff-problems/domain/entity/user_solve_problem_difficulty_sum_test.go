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
