package controllers

import (
	"diff-problems/interfaces/controllers"
	"diff-problems/test_tool"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func Test_Show_正常系(t *testing.T) {
	handler, err := test_tool.TruncateTestTables()
	assert.Nil(t, err)

	problemId1 := "problem1"
	problemDiff1 := 800
	problemId2 := "problem2"
	problemDiff2 := 1600

	userId1 := "test-user1"
	userId2 := "test-user2"

	updatedEpochTime := 123456789

	_, err = handler.Execute(
		"INSERT problem_difficulties VALUES (?,?,?),(?,?,?)",
		problemId1, problemDiff1, problemDiff1,
		problemId2, problemDiff2, problemDiff2,
	)
	assert.Nil(t, err)
	_, err = handler.Execute(
		"INSERT user_first_ac_submissions VALUES (?,?,?),(?,?,?),(?,?,?)",
		userId1, problemId1, 111111,
		userId1, problemId2, 222222,
		userId2, problemId1, 333333,
	)
	assert.Nil(t, err)
	_, err = handler.Execute(
		"INSERT user_first_ac_submission_updated_at VALUES (?)",
		updatedEpochTime,
	)
	assert.Nil(t, err)

	w1 := httptest.NewRecorder()
	c1, _ := gin.CreateTestContext(w1)
	c1.Params = []gin.Param{{Key: "user_id", Value: userId1}}

	controller := controllers.NewUserProblemController(handler)
	controller.Show(c1)

	expectedUser1Body := `{"problemCount":{"Black":{"all":0,"solved":0},"Blue":{"all":1,"solved":1},"Bronze":{"all":0,"solved":0},"Brown":{"all":0,"solved":0},"Cyan":{"all":0,"solved":0},"Gold":{"all":0,"solved":0},"Gray":{"all":0,"solved":0},"Green":{"all":1,"solved":1},"Orange":{"all":0,"solved":0},"Red":{"all":0,"solved":0},"Silver":{"all":0,"solved":0},"Yellow":{"all":0,"solved":0}},"updatedEpochTime":"123456789","userId":"test-user1","userSolveClipDifficultyTotal":2400}`
	assert.JSONEq(t, expectedUser1Body, w1.Body.String())

	w2 := httptest.NewRecorder()
	c2, _ := gin.CreateTestContext(w2)
	c2.Params = []gin.Param{{Key: "user_id", Value: userId2}}

	controller2 := controllers.NewUserProblemController(handler)
	controller2.Show(c2)

	expectedUser2Body := `{"problemCount":{"Black":{"all":0,"solved":0},"Blue":{"all":1,"solved":0},"Bronze":{"all":0,"solved":0},"Brown":{"all":0,"solved":0},"Cyan":{"all":0,"solved":0},"Gold":{"all":0,"solved":0},"Gray":{"all":0,"solved":0},"Green":{"all":1,"solved":1},"Orange":{"all":0,"solved":0},"Red":{"all":0,"solved":0},"Silver":{"all":0,"solved":0},"Yellow":{"all":0,"solved":0}},"updatedEpochTime":"123456789","userId":"test-user2","userSolveClipDifficultyTotal":800}`
	assert.JSONEq(t, expectedUser2Body, w2.Body.String())
}
