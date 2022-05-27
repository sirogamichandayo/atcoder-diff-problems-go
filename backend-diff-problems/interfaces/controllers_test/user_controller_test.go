package controllers_test

import (
	api "diff-problems/interfaces/api/mock"
	"diff-problems/interfaces/controllers"
	web "diff-problems/interfaces/web/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func Test_Show_正常系(t *testing.T) {
	ctrl := gomock.NewController(t)
	scrapeHandlerMock := web.NewMockScrapeHandler(ctrl)
	requestHandlerMock := api.NewMockRequestHandler(ctrl)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "user_id", Value: "sirogamichandayo"}}

	controller := controllers.NewUserController(scrapeHandlerMock, requestHandlerMock)
	controller.Show(c)

	assert.Equal(t, `{}`, w.Body.String())
}
