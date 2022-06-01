package entity

import (
	"diff-problems/domain/vo"
	"encoding/json"
)

type UserSubmission struct {
	UserId    string
	ProblemId string
	Result    vo.Result
	EpochTime int64
}

func (s UserSubmission) IsAc() bool {
	return s.Result.IsAc()
}

type UserSubmissionList []UserSubmission

func (list *UserSubmissionList) ExactByAc() (AcUserSubmissionList, error) {
	acList := make(AcUserSubmissionList, 0)
	for _, rawSubmission := range *list {
		if rawSubmission.IsAc() {
			newSubmission, err := MakeAcUserSubmissionFromUserSubmission(rawSubmission)
			if err != nil {
				return nil, err
			}
			acList = append(acList, newSubmission)
		}
	}
	return acList, nil
}

func (list *UserSubmissionList) LastEpochTime() (res int64) {
	res = 0
	for _, rawSubmission := range *list {
		if rawSubmission.EpochTime > res {
			res = rawSubmission.EpochTime
		}
	}
	return
}

func (list *UserSubmissionList) IsEmpty() bool {
	return len(*list) == 0
}

func MakeUserSubmissionListFromJsonBytes(bytes []byte) (UserSubmissionList, error) {
	var rawUserSubmissionList []map[string]interface{}
	err := json.Unmarshal(bytes, &rawUserSubmissionList)
	if err != nil {
		return nil, err
	}

	list := make(UserSubmissionList, 0)
	for _, submission := range rawUserSubmissionList {
		userId := submission["user_id"].(string)
		problemId := submission["problem_id"].(string)
		result := vo.ParseResult(submission["result"].(string))
		epochSecond := (int64)(submission["epoch_second"].(float64))

		list = append(list, UserSubmission{
			UserId:    userId,
			ProblemId: problemId,
			Result:    result,
			EpochTime: epochSecond,
		})
	}

	return list, nil
}
