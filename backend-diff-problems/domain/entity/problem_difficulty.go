package entity

import (
	"diff-problems/domain/vo"
	"encoding/json"
	"strings"
)

type ProblemDifficulty struct {
	ProblemId      string
	Difficulty     vo.RawDifficulty
	ClipDifficulty vo.ClipDifficulty
}

type ProblemDifficultyList []ProblemDifficulty

func (list ProblemDifficultyList) ClipDifficultyTotal() float64 {
	total := float64(0)
	for _, entity := range list {
		cd := entity.ClipDifficulty
		if !cd.Valid {
			continue
		}
		total += cd.Difficulty
	}
	return total
}

func (list ProblemDifficultyList) CountByColor(color vo.ProblemColor) int {
	cnt := 0
	for _, entity := range list {
		if entity.ClipDifficulty.EqualColor(color) {
			cnt++
		}
	}
	return cnt
}

func (list ProblemDifficultyList) MakeValueForUpsertMySql() (string, []interface{}) {
	listSize := len(list)
	placeholders := make([]string, 0, listSize)
	for i := 0; i < listSize; i++ {
		placeholders = append(placeholders, "(?,?,?)")
	}

	values := make([]interface{}, 0, listSize*3)
	for _, problem := range list {
		values = append(
			values,
			problem.ProblemId,
			problem.Difficulty,
			problem.ClipDifficulty,
		)
	}

	return strings.Join(placeholders, ","), values
}

func MakeProblemDifficultyListFromJsonBytes(bytes []byte) (list ProblemDifficultyList, err error) {
	var rawProblemDifficultyMap map[string]map[string]interface{}
	err = json.Unmarshal(bytes, &rawProblemDifficultyMap)
	if err != nil {
		return
	}

	for problemId, problem := range rawProblemDifficultyMap {
		rawDifficulty, hasKey := problem["difficulty"]

		difficultyVo, err := vo.NewRawDifficulty(rawDifficulty, hasKey)
		if err != nil {
			return nil, err
		}
		list = append(list, ProblemDifficulty{
			ProblemId: problemId, Difficulty: difficultyVo, ClipDifficulty: difficultyVo.MakeClipDifficulty(),
		})
	}

	return
}
