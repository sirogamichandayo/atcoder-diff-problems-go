package controllers_test

import (
	"diff-problems/interfaces/controllers"
	"diff-problems/test_tool"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
)

func Test_ShowDiff_正常系(t *testing.T) {
	handler, err := test_tool.TruncateTestTables()
	assert.Nil(t, err)

	_, err = handler.Execute(`INSERT INTO user_solve_problem_difficulty_sum (user_id, clip_difficulty_sum, rnk) VALUES ('aaa', 1000.0, 123);`)
	assert.Nil(t, err)
	_, err = handler.Execute(`INSERT INTO user_solve_problem_difficulty_sum_updated_at (updated_epoch_time) VALUES (12345);`)
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	controller := controllers.NewUserController(handler)
	controller.ShowDiff(c)

	assert.Equal(t, `{"UserSum":{"UserId":"aaa","ClipDifficultySum":1000,"Rank":123},"RankUpdatedEpochTime":12345}`, w.Body.String())
}
