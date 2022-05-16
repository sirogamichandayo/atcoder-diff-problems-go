package atcoder_problems_api

import (
	"diff-problems/domain/entity"
	"diff-problems/interfaces/api"
	"fmt"
)

type UserSubmissionClient struct {
	api.RequestHandler
}

func (client *UserSubmissionClient) FetchSinceByEpochTime(fromEpochTime int64) (list entity.UserSubmissionList, err error) {
	url := fmt.Sprintf("https://kenkoooo.com/atcoder/atcoder-api/v3/from/%d", fromEpochTime)
	res, err := client.Get(url, nil)
	if err != nil {
		return
	}
	if !res.IsSuccess() {
		bytes := res.BodyBytes()
		if err != nil {
			return nil, fmt.Errorf("リクエスト失敗しました。エラーメッセージの生成にも失敗しました")
		}
		return nil, fmt.Errorf(string(bytes))
	}

	bytes := res.BodyBytes()
	if err != nil {
		return
	}

	list, err = entity.MakeUserSubmissionListFromJsonBytes(bytes)
	return
}
