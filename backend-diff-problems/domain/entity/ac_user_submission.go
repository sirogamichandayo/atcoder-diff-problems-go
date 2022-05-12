package entity

import (
	"fmt"
	"strings"
)

type AcUserSubmission struct {
	UserId    string
	ProblemId string
	EpochTime int64
}

func MakeAcUserSubmissionFromUserSubmission(rawSubmission UserSubmission) (
	acSubmission AcUserSubmission,
	err error,
) {
	if !rawSubmission.IsAc() {
		err = fmt.Errorf("this submission result is not ac")
		return
	}
	acSubmission = AcUserSubmission{
		UserId:    rawSubmission.UserId,
		ProblemId: rawSubmission.ProblemId,
		EpochTime: rawSubmission.EpochTime,
	}
	return
}

type AcUserSubmissionList []AcUserSubmission

func (list *AcUserSubmissionList) IsEmpty() bool {
	return len(*list) == 0
}

func (list *AcUserSubmissionList) MakeValueForUpsertMySql() (string, []interface{}) {
	listSize := len(*list)
	placeholders := make([]string, 0, listSize)
	for i := 0; i < listSize; i++ {
		placeholders = append(placeholders, "(?,?,?)")
	}

	values := make([]interface{}, 0, listSize*3)
	for _, problem := range *list {
		values = append(values, problem.UserId, problem.ProblemId, problem.EpochTime)
	}
	return strings.Join(placeholders, ","), values
}
