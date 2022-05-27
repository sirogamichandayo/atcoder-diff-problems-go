package queryService

import (
	"diff-problems/domain/client"
	"diff-problems/domain/vo"
	"diff-problems/interfaces/web"
	cqrsDto "diff-problems/usecase/cqrs_dto"
	"fmt"
	"strconv"
)

type UserService struct {
	web.ScrapeHandler
	ContestResultClient client.ContestResultClient
}

func (s UserService) FindByUserId(userId string) (cqrsDto.User, error) {
	url := fmt.Sprintf("https://atcoder.jp/users/%s", userId)
	userDocument, err := s.NewDocument(url)
	if err != nil {
		return cqrsDto.User{}, err
	}
	// 優勝経験あるユーザーのselector
	imageUrl, exist := userDocument.
		Find("#main-container > div.row > div.col-md-3.col-sm-12 > img:nth-child(2)").
		Attr("src")
	if !exist {
		// 優勝経験ないユーザーのselector
		imageUrl, _ = userDocument.
			Find("#main-container > div.row > div.col-md-3.col-sm-12 > img").
			Attr("src")
	}

	rating := vo.NewNoRating()
	resultList, err := s.ContestResultClient.All(userId)
	if err != nil {
		return cqrsDto.User{}, err
	}
	if !resultList.Empty() {
		lastResult, err := resultList.Last()
		if err != nil {
			return cqrsDto.User{}, err
		}
		rating = lastResult.Rating
	}

	var ranking *int
	rankingStr := userDocument.Find("#main-container > div.row > div.col-md-9.col-sm-12 > table > tbody > tr:nth-child(1) > td").Text()
	if len(rankingStr) > 2 {
		t, _ := strconv.Atoi(rankingStr[0 : len(rankingStr)-2])
		ranking = &t
	}
	return cqrsDto.NewUser(userId, imageUrl, ranking, rating), nil
}
