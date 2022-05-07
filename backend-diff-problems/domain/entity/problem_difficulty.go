package entity

import (
	"encoding/json"
	"strings"
)

type ProblemDifficulty struct {
	ProblemId  string
	Difficulty *float64
}

type ProblemDifficultyList []ProblemDifficulty

func (list *ProblemDifficultyList) MakeValueForUpsertMySql() (string, []interface{}) {
	listSize := len(*list)
	placeholders := make([]string, 0, listSize)
	for i := 0; i < listSize; i++ {
		placeholders = append(placeholders, "(?,?)")
	}

	values := make([]interface{}, 0, listSize*2)
	for _, problem := range *list {
		values = append(values, problem.ProblemId, problem.Difficulty)
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

		var difficulty *float64
		if !hasKey {
			difficulty = nil
		} else {
			tmp := rawDifficulty.(float64)
			difficulty = &tmp
		}

		list = append(list, ProblemDifficulty{ProblemId: problemId, Difficulty: difficulty})
	}

	return
}
