package commands_test

import (
	"diff-problems/interfaces/commands"
	mock_api "diff-problems/interfaces/commands_test/mock"
	"diff-problems/test_tool"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_正常系(t *testing.T) {
	sqlHandler, err := test_tool.TruncateTestTables()
	assert.Nil(t, err)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	requestHandlerMock := mock_api.NewMockRequestHandler(ctrl)
	responseMock := mock_api.NewMockResponse(ctrl)

	requestHandlerMock.
		EXPECT().
		Get(
			gomock.Eq("https://kenkoooo.com/atcoder/resources/problem-models.json"),
			gomock.Eq(map[string]string(nil)),
		).
		Return(responseMock, nil)

	responseMock.
		EXPECT().BodyBytes().
		Return([]byte(`{"abc138_a": {"difficulty": -848}, "abc138_b": {}}`), nil)

	command := commands.NewUpdateProblemDifficultyCommand(sqlHandler, requestHandlerMock)
	err = command.Exec()
	assert.Nil(t, err)

	rows, err := sqlHandler.Query("SELECT * FROM product_difficulties")
	assert.Nil(t, err)

	expected := map[string]func(*float64) bool{
		"abc138_a": func(val *float64) bool {
			return *val == float64(-848)
		},
		"abc138_b": func(val *float64) bool {
			return val == nil
		},
	}

	var problemId string
	var difficulty *float64

	for rows.Next() {
		err := rows.Scan(&problemId, &difficulty)
		assert.Nil(t, err)

		fn, ok := expected[problemId]
		assert.True(t, ok)
		assert.True(t, fn(difficulty))
	}
}
