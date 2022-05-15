package commands

import (
	"diff-problems/interfaces/commands"
	"diff-problems/test_tool"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Update_正常系(t *testing.T) {
	handler, err := test_tool.TruncateTestTables()
	assert.Nil(t, err)

	updatedEpochTime := int64(12345)
	userId1 := "user_id1"
	userId2 := "user_id2"
	problemId1 := "problem_id1"
	problem1Difficulty := 1000.000
	problem1ClipDifficulty := 1000.000
	problemId2 := "problem_id2"
	problem2Difficulty := 2000.000
	problem2ClipDifficulty := 2000.000
	solvedEpochTime := int64(12345555)

	_, err = handler.Execute(`INSERT INTO user_first_ac_submission_updated_at (updated_epoch_time) values (?)`, updatedEpochTime)
	assert.Nil(t, err)
	_, err = handler.Execute(`
INSERT INTO user_first_ac_submissions (user_id, problem_id, first_solved_epoch_time)
VALUES (?,?,?),(?,?,?),(?,?,?)`,
		userId1, problemId1, solvedEpochTime, userId1, problemId2, solvedEpochTime, userId2, problemId1, solvedEpochTime,
	)
	assert.Nil(t, err)
	_, err = handler.Execute(`
INSERT INTO problem_difficulties (problem_id, difficulty, clip_difficulty)
VALUES (?,?,?),(?,?,?)
`, problemId1, problem1Difficulty, problem1ClipDifficulty, problemId2, problem2Difficulty, problem2ClipDifficulty)
	assert.Nil(t, err)

	command := commands.NewUserSolveProblemDifficultySumCommand(handler)
	err = command.Update()
	assert.Nil(t, err)

	var userId string
	var difficultySum float64
	var rank uint64
	rows, err := handler.Query("SELECT * from user_solve_problem_difficulty_sum ORDER BY user_id ASC")
	assert.Nil(t, err)
	defer rows.Close()

	assert.True(t, rows.Next())
	err = rows.Scan(&userId, &difficultySum, &rank)
	assert.Nil(t, err)
	assert.Equal(t, userId1, userId)
	assert.Equal(t, difficultySum, problem1ClipDifficulty+problem2ClipDifficulty)
	assert.Same(t, uint64(1), rank)

	assert.True(t, rows.Next())
	err = rows.Scan(&userId, &difficultySum, &rank)
	assert.Nil(t, err)
	assert.Equal(t, userId2, userId)
	assert.Equal(t, difficultySum, problem1ClipDifficulty)
	assert.Same(t, uint64(2), rank)

	assert.False(t, rows.Next())

	var epochTime int64
	row, err := handler.Query("SELECT * from user_solve_problem_difficulty_sum_updated_at")
	assert.Nil(t, err)

	assert.True(t, row.Next())
	err = row.Scan(&epochTime)
	assert.Nil(t, err)
	assert.Equal(t, updatedEpochTime, epochTime)

	assert.False(t, row.Next())
}
