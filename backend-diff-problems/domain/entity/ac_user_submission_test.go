package entity

import (
	"diff-problems/domain/vo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MakeAcUserSubmissionFromUserSubmission_正常系(t *testing.T) {
	userId := "user_id"
	problemId := "problem_id"
	result := vo.Ac
	epochTime := int64(1234567)
	userSubmission := UserSubmission{
		UserId:    userId,
		ProblemId: problemId,
		Result:    result,
		EpochTime: epochTime,
	}

	actual, err := MakeAcUserSubmissionFromUserSubmission(userSubmission)
	assert.Nil(t, err)
	assert.Equal(t, userId, actual.UserId)
	assert.Equal(t, problemId, actual.ProblemId)
	assert.Equal(t, epochTime, actual.EpochTime)
}

func Test_MakeAcUserSubmissionFromUserSubmission_異常系_resultがAcじゃない(t *testing.T) {
	userId := "user_id"
	problemId := "problem_id"
	result := vo.Wa
	epochTime := int64(1234567)
	userSubmission := UserSubmission{
		UserId:    userId,
		ProblemId: problemId,
		Result:    result,
		EpochTime: epochTime,
	}

	actual, err := MakeAcUserSubmissionFromUserSubmission(userSubmission)
	assert.NotNil(t, err)
	assert.Equal(t, "this submission result is not ac", err.Error())
	assert.Equal(t, AcUserSubmission{}, actual)
}

func Test_MakeValueForUpsertMySql_正常系(t *testing.T) {
	userId1 := "user_id1"
	problemId1 := "problem_id1"
	epochTime1 := int64(11111)
	userId2 := "user_id2"
	problemId2 := "problem_id2"
	epochTime2 := int64(22222)

	list := AcUserSubmissionList{
		AcUserSubmission{UserId: userId1, ProblemId: problemId1, EpochTime: epochTime1},
		AcUserSubmission{UserId: userId2, ProblemId: problemId2, EpochTime: epochTime2},
	}

	placeholder, values := list.MakeValueForUpsertMySql()
	assert.Equal(
		t, "(?,?,?),(?,?,?)",
		placeholder,
		[]interface{}{userId1, problemId1, epochTime1, userId1, problemId1, epochTime1},
	)
	assert.Equal(
		t,
		[]interface{}{userId1, problemId1, epochTime1, userId2, problemId2, epochTime2},
		values,
	)
}
