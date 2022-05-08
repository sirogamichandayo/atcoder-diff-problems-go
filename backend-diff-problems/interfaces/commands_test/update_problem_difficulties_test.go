package commands_test

import (
	"diff-problems/interfaces/commands"
	mock_api "diff-problems/interfaces/commands_test/mock"
	"diff-problems/interfaces/database"
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

	var problemId string
	var difficulty *float64
	var rows database.Row

	rows, err = sqlHandler.Query(`SELECT * FROM product_difficulties where "problem_id" = "abc138_a";`)
	defer rows.Close()
	assert.Nil(t, err)
	rows.Next()
	assert.Nil(t, rows.Scan(&problemId, &difficulty))
	assert.Equal(t, "abc138_a", problemId)
	assert.Equal(t, -848, difficulty)

	// TODO: テスト修正

}
