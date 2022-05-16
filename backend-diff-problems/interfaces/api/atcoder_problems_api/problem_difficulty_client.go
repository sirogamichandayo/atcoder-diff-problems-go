package atcoder_problems_api

import (
	"diff-problems/domain/entity"
	"diff-problems/interfaces/api"
)

type ProblemDifficultyClient struct {
	api.RequestHandler
}

func (client *ProblemDifficultyClient) Fetch() (list entity.ProblemDifficultyList, err error) {
	res, err := client.Get("https://kenkoooo.com/atcoder/resources/problem-models.json", nil)
	if err != nil {
		return
	}
	bytes := res.BodyBytes()

	list, err = entity.MakeProblemDifficultyListFromJsonBytes(bytes)
	return
}
