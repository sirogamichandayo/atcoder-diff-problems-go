package queryService

import (
	"diff-problems/domain/vo"
	"diff-problems/interfaces/api/atcoder_api"
	"diff-problems/interfaces/web"
	cqrsDto "diff-problems/usecase/cqrs_dto"
	"fmt"
)

type UserService struct {
	web.ScrapeHandler
	ContestResultClient atcoder_api.ContestResultClient
}

func (s UserService) FindByUserId(userId string) (cqrsDto.User, error) {

	url := fmt.Sprintf("https://atcoder.jp/users/%s", userId)
	userDocument, err := s.NewDocument(url)
	if err != nil {
		return cqrsDto.User{}, err
	}
	imageUrl, exist := userDocument.
		Find("#main-container > div.row > div.col-md-3.col-sm-12 > img:nth-child(2)").
		Attr("src")
	if !exist {
		imageUrl, exist = userDocument.
			Find("#main-container > div.row > div.col-md-3.col-sm-12 > img").
			Attr("src")

		if !exist {
			// TODO: logを残す
		}
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

	return cqrsDto.NewUser(userId, imageUrl, rating), nil
}
