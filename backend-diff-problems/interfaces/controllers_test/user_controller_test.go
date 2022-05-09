package controllers_test

import (
	"diff-problems/interfaces/controllers"
	"diff-problems/test_tool"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func Test_Index正常系(t *testing.T) {
	handler, err := test_tool.TruncateTestTables()
	assert.Nil(t, err)

	_, err = handler.Execute("insert into users (id, first_name, last_name) values (?, ?, ?)", 1, "sirogami", "kurogami")
	assert.Nil(t, err)

	userController := controllers.NewUserController(handler)

	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)

	userController.Index(context)

	assert.Equal(t, `[{"Id":1,"FirstName":"sirogami","LastName":"kurogami"}]`, response.Body.String())
}
