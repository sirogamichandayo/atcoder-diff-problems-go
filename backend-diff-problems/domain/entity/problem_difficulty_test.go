package entity

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_makeFromJsonBytes(t *testing.T) {
	json := `{"abc138_a":{"difficulty": -848}, "abc138_b": {}}`
	list, err := MakeProblemDifficultyListFromJsonBytes([]byte(json))
	assert.Nil(t, err)

	assert.Len(t, list, 2)
	entity1 := list[0]
	entity2 := list[1]
	fmt.Println(entity2)
	assert.Equal(t, "abc138_a", entity1.ProblemId)
	assert.Equal(t, float64(-848), *entity1.Difficulty)
	assert.Equal(t, "abc138_b", entity2.ProblemId)
	assert.Nil(t, entity2.Difficulty)
}

func Test_makeMySqlValues(t *testing.T) {
	problemId1 := "problem_id1"
	problemId2 := "problem_id2"
	difficulty1 := -1000.12345

	list := ProblemDifficultyList{
		ProblemDifficulty{problemId1, &difficulty1},
		ProblemDifficulty{problemId2, nil},
	}

	expected := "(problem_id1,-1000.12345),(problem_id2,NULL)"

	assert.Equal(t, expected, list.MakeMySqlValues())
}
