package entity

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type ProblemDifficulty struct {
	ProblemId  string
	Difficulty *float64
}

type ProblemDifficultyList []ProblemDifficulty

func (list *ProblemDifficultyList) MakeMySqlValues() string {
	res := make([]string, 0)
	for _, difficulty := range *list {
		diffStr := "NULL"
		if difficulty.Difficulty != nil {
			diffStr = strconv.FormatFloat(*difficulty.Difficulty, 'f', -1, 64)
		}
		res = append(res, fmt.Sprintf("(%s,%s)", difficulty.ProblemId, diffStr))
	}
	return strings.Join(res, ",")
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
			fmt.Println("hoge")
			difficulty = nil
		} else {
			tmp := rawDifficulty.(float64)
			difficulty = &tmp
		}

		list = append(list, ProblemDifficulty{ProblemId: problemId, Difficulty: difficulty})
	}

	return
}
