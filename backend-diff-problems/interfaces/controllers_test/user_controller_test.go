package controllers_test

import (
	"diff-problems/config"
	"diff-problems/infrastructure"
	"diff-problems/interfaces/controllers"
	"diff-problems/test_tool"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func Test_Index正常系(t *testing.T) {
	sinDb := config.SinDb{
		Host:     "sin-mariadb",
		Port:     "3306",
		User:     "root",
		Password: "secret",
		Database: "sample",
	}
	handler := infrastructure.NewSqlHandler(sinDb)
	test_tool.TruncateTables(handler)

	_, err := handler.Execute("insert into users (id, first_name, last_name) values (?, ?, ?)", 1, "sirogami", "kurogami")
	assert.Nil(t, err)

	userController := controllers.NewUserController(
		infrastructure.NewSqlHandler(sinDb),
	)

	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)

	userController.Index(context)

	assert.Equal(t, `[{"Id":1,"FirstName":"sirogami","LastName":"kurogami"}]`, response.Body.String())
}
