package controllers

import (
	"diff-problems/interfaces/controllers"
	"diff-problems/test_tool"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_RankController_Show_正常系(t *testing.T) {
	handler, err := test_tool.TruncateTestTables()
	assert.Nil(t, err)

	_, err = handler.Execute(`
INSERT INTO user_solve_problem_difficulty_sum
VALUES 
("test1", 1000.0, 1),("test2", 900.0, 2),("test3", 700.0, 3),
("test4", 600.0, 4) ,("test5", 500.0, 5),("test6", 400.0, 6),
("test7", 300.0, 7) ,("test8", 200.0, 2) 
`)
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, err = http.NewRequest("GET", "http://127.0.0.1:8080/api/v1/ranks/paging?offset=1&limit=3", nil)
	assert.Nil(t, err)

	controller := controllers.NewRankController(handler)
	controller.Show(c)

	assert.JSONEq(t,
		`
[{"clipDifficultySum":900,"rank":2,"userId":"test2"},
{"clipDifficultySum":200,"rank":2,"userId":"test8"},
{"clipDifficultySum":700,"rank":3,"userId":"test3"}]`,
		w.Body.String(),
	)
}
