package controllers

import (
	"diff-problems/interfaces/controllers"
	"diff-problems/test_tool"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func Test_RankNearController_Show_正常系(t *testing.T) {
	handler, err := test_tool.TruncateTestTables()
	assert.Nil(t, err)

	_, err = handler.Execute(`
INSERT INTO user_solve_problem_difficulty_sum
VALUES 
("test1", 1200.0, 1),("test2", 1100.0, 2) ,("test3", 1000.0, 3),
("test4", 1000.0, 3),("test5", 1000.0, 3),("test6", 1000.0, 3),
("test7", 1000.0, 3),("test8", 1000.0, 3) 
`)
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "user_id", Value: "test1"}}
	assert.Nil(t, err)

	controller := controllers.NewRankNearController(handler)
	controller.Show(c)

	assert.JSONEq(t,
		`[
{"clipDifficultySum":1200,"rank":1,"userId":"test1"},
{"clipDifficultySum":1100,"rank":2,"userId":"test2"},
{"clipDifficultySum":1000,"rank":3,"userId":"test3"},
{"clipDifficultySum":1000,"rank":3,"userId":"test4"},
{"clipDifficultySum":1000,"rank":3,"userId":"test5"},
{"clipDifficultySum":1000,"rank":3,"userId":"test6"}
]`,
		w.Body.String(),
	)
}
