package entity

import (
	"diff-problems/domain/vo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_makeFromJsonBytes(t *testing.T) {
	json := `{"abc138_a":{"difficulty": -848}, "abc138_b": {}}`

	list, err := MakeProblemDifficultyListFromJsonBytes([]byte(json))
	assert.Nil(t, err)

	expected := ProblemDifficultyList{
		ProblemDifficulty{ProblemId: "abc138_a", Difficulty: vo.RawDifficulty{Difficulty: -848, Valid: true}, ClipDifficulty: vo.ClipDifficulty{Difficulty: 17.662867367877144, Valid: true}},
		ProblemDifficulty{ProblemId: "abc138_b", Difficulty: vo.RawDifficulty{Difficulty: 0, Valid: false}, ClipDifficulty: vo.ClipDifficulty{Difficulty: 0, Valid: false}},
	}
	assert.ElementsMatch(t, expected, list)
}
