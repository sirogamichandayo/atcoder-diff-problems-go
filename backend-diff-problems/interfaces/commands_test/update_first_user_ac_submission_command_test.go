package commands

import (
	"diff-problems/interfaces/commands"
	mockApi "diff-problems/interfaces/commands_test/mock"
	"diff-problems/test_tool"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_UpdateAll_正常系(t *testing.T) {
	sqlHandler, err := test_tool.TruncateTestTables()
	assert.Nil(t, err)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	requestHandlerMock := mockApi.NewMockRequestHandler(ctrl)
	responseMock := mockApi.NewMockResponse(ctrl)

	gomock.InOrder(
		requestHandlerMock.
			EXPECT().
			Get(
				gomock.Eq("https://kenkoooo.com/atcoder/atcoder-api/v3/from/1"),
				gomock.Eq(map[string]string(nil)),
			).
			Return(responseMock, nil).
			Times(1),
		responseMock.
			EXPECT().
			IsSuccess().
			Return(true).
			Times(1),
		responseMock.
			EXPECT().
			BodyBytes().
			Return([]byte(
				`[
{"epoch_second":111111111,"problem_id":"problem1","user_id":"test_user1","result":"AC"},
{"epoch_second":222222222,"problem_id":"problem2","user_id":"test_user2","result":"WA"},
{"epoch_second":333333333,"problem_id":"problem3","user_id":"test_user3","result":"AC"},
{"epoch_second":222222222,"problem_id":"problem3","user_id":"test_user3","result":"AC"},
{"epoch_second":333333333,"problem_id":"problem3","user_id":"test_user3","result":"AC"}
]`,
			), nil).
			Times(1),
		requestHandlerMock.
			EXPECT().
			Get(
				gomock.Eq("https://kenkoooo.com/atcoder/atcoder-api/v3/from/333333334"),
				gomock.Eq(map[string]string(nil)),
			).
			Return(responseMock, nil).
			Times(1),
		responseMock.
			EXPECT().
			IsSuccess().
			Return(true).
			Times(1),
		responseMock.
			EXPECT().
			BodyBytes().
			Return([]byte(`[]`), nil).
			Times(1),
	)

	command := commands.NewUpdateUserFirstAcSubmissionCommand(sqlHandler, requestHandlerMock)
	err = command.UpdateAll()
	assert.Nil(t, err)

	rows, err := sqlHandler.Query(
		`select * from user_first_ac_submissions order by user_id, problem_id`,
	)
	assert.Nil(t, err)

	expected := [][]interface{}{
		{"test_user1", "problem1", 111111111},
		{"test_user3", "problem3", 222222222},
	}

	var actual [][]interface{}
	var a string
	var b string
	var c int

	for rows.Next() {
		err := rows.Scan(&a, &b, &c)
		assert.Nil(t, err)
		actual = append(actual, []interface{}{a, b, c})
	}

	assert.Equal(t, expected, actual)
}
