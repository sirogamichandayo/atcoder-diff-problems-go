package controllers_test

import (
	"diff-problems/domain/vo"
	api "diff-problems/interfaces/api/mock"
	"diff-problems/interfaces/controllers"
	web "diff-problems/interfaces/web/mock"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func Test_Show_正常系(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	scrapeHandlerMock := web.NewMockScrapeHandler(ctrl)
	documentMock := web.NewMockDocument(ctrl)
	selectionMock1 := web.NewMockSelection(ctrl)
	selectionMock2 := web.NewMockSelection(ctrl)
	requestHandlerMock := api.NewMockRequestHandler(ctrl)
	responseMock := api.NewMockResponse(ctrl)

	userId := "test-user-id"
	imageUrl := "https://image.com/test_user.png"
	ranking := 123
	lastRating := 1234

	gomock.InOrder(
		scrapeHandlerMock.
			EXPECT().
			NewDocument("https://atcoder.jp/users/"+userId).
			Return(documentMock, nil).
			Times(1),
		documentMock.
			EXPECT().
			Find("#main-container > div.row > div.col-md-3.col-sm-12 > img:nth-child(2)").
			Return(selectionMock1).
			Times(1),
		selectionMock1.
			EXPECT().
			Attr("src").
			Return(imageUrl, true).
			Times(1),
		documentMock.
			EXPECT().
			Find("#main-container > div.row > div.col-md-9.col-sm-12 > table > tbody > tr:nth-child(1) > td").
			Return(selectionMock2).
			Times(1),
		selectionMock2.
			EXPECT().
			Text().
			Return(fmt.Sprintf("%dth", ranking)).
			Times(1),
		requestHandlerMock.
			EXPECT().
			Get(fmt.Sprintf("https://atcoder.jp/users/%s/history/json", userId), nil).
			Return(responseMock, nil).
			Times(1),
		responseMock.
			EXPECT().
			BodyBytes().
			Return([]byte(fmt.Sprintf(`[{"IsRated":true,"Place":59,"OldRating":0,"NewRating":1255,"Performance":2455,"InnerPerformance":2455,"ContestScreenName":"arc061.contest.atcoder.jp","ContestName":"AtCoder Regular Contest 061","ContestNameEn":"","EndTime":"2016-09-11T22:40:00+09:00"},{"IsRated":true,"Place":9,"OldRating":1255,"NewRating":2155,"Performance":3192,"InnerPerformance":3192,"ContestScreenName":"arc064.contest.atcoder.jp","ContestName":"AtCoder Regular Contest 064","ContestNameEn":"","EndTime":"2016-12-04T22:40:00+09:00"},{"IsRated":true,"Place":8,"OldRating":2155,"NewRating":2676,"Performance":3623,"InnerPerformance":3623,"ContestScreenName":"agc010.contest.atcoder.jp","ContestName":"AtCoder Grand Contest 010","ContestNameEn":"","EndTime":"2017-02-04T22:50:00+09:00"},{"IsRated":true,"Place":41,"OldRating":2676,"NewRating":2728,"Performance":2974,"InnerPerformance":2974,"ContestScreenName":"agc014.contest.atcoder.jp","ContestName":"AtCoder Grand Contest 014","ContestNameEn":"","EndTime":"2017-05-06T23:10:00+09:00"},{"IsRated":true,"Place":45,"OldRating":2728,"NewRating":%d,"Performance":2813,"InnerPerformance":2813,"ContestScreenName":"agc015.contest.atcoder.jp","ContestName":"AtCoder Grand Contest 015","ContestNameEn":"","EndTime":"2017-05-27T22:50:00+09:00"}]`, lastRating))).
			Times(1),
	)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "user_id", Value: userId}}

	controller := controllers.NewUserController(scrapeHandlerMock, requestHandlerMock)
	controller.Show(c)

	expected := fmt.Sprintf(`
{
"userId":"%s",
"imageUrl":"%s",
"ranking":%d,
"rating":%d,
"color":"%s"
}
`, userId, imageUrl, ranking, lastRating, vo.CyanRating)

	assert.JSONEq(t, expected, w.Body.String())
}
