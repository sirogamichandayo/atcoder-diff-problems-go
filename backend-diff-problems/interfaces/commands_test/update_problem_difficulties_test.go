package commands_test

import (
	api "diff-problems/interfaces/api/mock"
	"diff-problems/interfaces/commands"
	"diff-problems/interfaces/database"
	"diff-problems/test_tool"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_正常系(t *testing.T) {
	sqlHandler, err := test_tool.TruncateTestTables()
	assert.Nil(t, err)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	requestHandlerMock := api.NewMockRequestHandler(ctrl)
	responseMock := api.NewMockResponse(ctrl)

	requestHandlerMock.
		EXPECT().
		Get(
			gomock.Eq("https://kenkoooo.com/atcoder/resources/problem-models.json"),
			gomock.Eq(map[string]string(nil)),
		).
		Return(responseMock, nil).
		Times(1)

	responseMock.
		EXPECT().BodyBytes().
		Return([]byte(`{"abc138_a": {"difficulty": -849}, "abc138_b": {}}`)).
		Times(1)

	command := commands.NewUpdateProblemDifficultyCommand(sqlHandler, requestHandlerMock)
	err = command.Exec()
	assert.Nil(t, err)

	var rawDifficulty *float64
	var clipDifficulty *float64

	row1, err := sqlHandler.Query(
		`
SELECT 
    difficulty, clip_difficulty
FROM 
    problem_difficulties
WHERE
    problem_id = "abc138_a"`)

	defer func(row1 database.Row) {
		err := row1.Close()
		assert.Nil(t, err)
	}(row1)

	assert.True(t, row1.Next())
	assert.Nil(t, row1.Scan(&rawDifficulty, &clipDifficulty))
	assert.Equal(t, float64(-849), *rawDifficulty)
	assert.Equal(t, 17.61876534994966, *clipDifficulty)

	row2, err := sqlHandler.Query(
		`
SELECT
    difficulty, clip_difficulty
FROM
    problem_difficulties
WHERE
    problem_id = "abc138_b"`)

	defer func(row2 database.Row) {
		err := row2.Close()
		assert.Nil(t, err)
	}(row2)

	assert.True(t, row2.Next())
	assert.Nil(t, row2.Scan(&rawDifficulty, &clipDifficulty))
	assert.Nil(t, rawDifficulty)
	assert.Nil(t, clipDifficulty)
}
