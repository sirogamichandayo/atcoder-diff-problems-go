package atcoder_api

import (
	"diff-problems/domain/vo"
	"diff-problems/interfaces/api"
	"diff-problems/interfaces/api/converter"
	"fmt"
)

type ContestResultClient struct {
	api.RequestHandler
}

const ContestResultFormat = "https://atcoder.jp/users/%s/history/json"

func (client ContestResultClient) FetchList(userId string) (vo.ContestResultList, error) {
	url := fmt.Sprintf(ContestResultFormat, userId)
	res, err := client.Get(url, nil)
	if err != nil {
		return nil, err
	}

	return converter.ConvertContestResultList(res.BodyBytes())
}
