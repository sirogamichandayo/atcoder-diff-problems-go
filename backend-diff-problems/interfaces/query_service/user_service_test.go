package queryService

import (
	client "diff-problems/domain/client/mock"
	"diff-problems/domain/vo"
	web "diff-problems/interfaces/web/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_FindByUserId_正常系_優勝経験あるユーザー(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	scrapeHandlerMock := web.NewMockScrapeHandler(ctrl)
	documentMock := web.NewMockDocument(ctrl)
	selectorMock1 := web.NewMockSelection(ctrl)
	selectorMock2 := web.NewMockSelection(ctrl)
	contestResultClientMock := client.NewMockContestResultClient(ctrl)

	rating1, _ := vo.NewRating(1200)
	rating2, _ := vo.NewRating(1199)
	result1 := vo.ContestResult{
		IsRated: false,
		Rating:  rating1,
		EndTime: 100,
	}
	result2 := vo.ContestResult{
		IsRated: false,
		Rating:  rating2,
		EndTime: 200,
	}
	resultList := vo.ContestResultList{result1, result2}
	imgUrl := "https://example.com/imageurl.png"
	rankingStr := "123th"

	userId := "sirogamichandayo"
	gomock.InOrder(
		scrapeHandlerMock.
			EXPECT().
			NewDocument("https://atcoder.jp/users/"+userId).
			Return(documentMock, nil).
			Times(1),
		documentMock.
			EXPECT().
			Find("#main-container > div.row > div.col-md-3.col-sm-12 > img:nth-child(2)").
			Return(selectorMock1).
			Times(1),
		selectorMock1.
			EXPECT().
			Attr("src").
			Return(imgUrl, true).
			Times(1),
		contestResultClientMock.
			EXPECT().
			All(userId).
			Return(resultList, nil).
			Times(1),
		documentMock.
			EXPECT().
			Find("#main-container > div.row > div.col-md-9.col-sm-12 > table > tbody > tr:nth-child(1) > td").
			Return(selectorMock2),
		selectorMock2.
			EXPECT().
			Text().
			Return(rankingStr).
			Times(1),
	)

	userService := UserService{scrapeHandlerMock, contestResultClientMock}
	actual, err := userService.FindByUserId(userId)
	assert.Nil(t, err)
	assert.Equal(t, result2.Rating, actual.Rating())
	assert.Equal(t, userId, actual.UserId())
	assert.Equal(t, imgUrl, actual.ImageUrl())
	assert.Equal(t, 123, *(actual.Ranking()))
}

func Test_FindByUserId_正常系_優勝経験ないユーザー(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	scrapeHandlerMock := web.NewMockScrapeHandler(ctrl)
	documentMock := web.NewMockDocument(ctrl)
	selectorMock1 := web.NewMockSelection(ctrl)
	selectorMock2 := web.NewMockSelection(ctrl)
	selectorMock3 := web.NewMockSelection(ctrl)
	contestResultClientMock := client.NewMockContestResultClient(ctrl)

	rating1, _ := vo.NewRating(1200)
	rating2, _ := vo.NewRating(1199)
	result1 := vo.ContestResult{
		IsRated: false,
		Rating:  rating1,
		EndTime: 100,
	}
	result2 := vo.ContestResult{
		IsRated: false,
		Rating:  rating2,
		EndTime: 200,
	}
	resultList := vo.ContestResultList{result1, result2}
	imgUrl := "https://example.com/imageurl.png"
	rankingStr := "123th"

	userId := "sirogamichandayo"
	gomock.InOrder(
		scrapeHandlerMock.
			EXPECT().
			NewDocument("https://atcoder.jp/users/"+userId).
			Return(documentMock, nil).
			Times(1),
		documentMock.
			EXPECT().
			Find("#main-container > div.row > div.col-md-3.col-sm-12 > img:nth-child(2)").
			Return(selectorMock1).
			Times(1),
		selectorMock1.
			EXPECT().
			Attr("src").
			Return("", false).
			Times(1),
		documentMock.
			EXPECT().
			Find("#main-container > div.row > div.col-md-3.col-sm-12 > img").
			Return(selectorMock3).
			Times(1),
		selectorMock3.
			EXPECT().
			Attr("src").
			Return(imgUrl, true).
			Times(1),
		contestResultClientMock.
			EXPECT().
			All(userId).
			Return(resultList, nil).
			Times(1),
		documentMock.
			EXPECT().
			Find("#main-container > div.row > div.col-md-9.col-sm-12 > table > tbody > tr:nth-child(1) > td").
			Return(selectorMock2),
		selectorMock2.
			EXPECT().
			Text().
			Return(rankingStr).
			Times(1),
	)

	userService := UserService{scrapeHandlerMock, contestResultClientMock}
	actual, err := userService.FindByUserId(userId)
	assert.Nil(t, err)
	assert.Equal(t, result2.Rating, actual.Rating())
	assert.Equal(t, userId, actual.UserId())
	assert.Equal(t, imgUrl, actual.ImageUrl())
	assert.Equal(t, 123, *(actual.Ranking()))
}
