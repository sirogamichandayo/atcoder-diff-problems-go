package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MakeValueForUpsertMySql(t *testing.T) {
	userId1 := "user1"
	diffSum1 := 11111.1
	userId2 := "user2"
	diffSum2 := 22222.2
	list := UserSolveProblemDifficultySumList{
		{userId1, diffSum1},
		{userId2, diffSum2},
	}

	placeholderValueList, err := list.MakeValuesForUpsertMySql()
	assert.Nil(t, err)
	placeholderValue := placeholderValueList[0]
	placeholder, values := placeholderValue.Placeholder(), placeholderValue.Values()

	assert.Equal(
		t,
		"(?,?),(?,?)",
		placeholder,
		[]interface{}{userId1, diffSum1, userId2, diffSum2},
	)
	assert.Equal(
		t,
		[]interface{}{userId1, diffSum1, userId2, diffSum2},
		values,
	)
}
