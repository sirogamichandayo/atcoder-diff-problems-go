package cqrsDto

import "diff-problems/domain/vo"

type UserRate struct {
	userId   string
	imageUrl string
	ranking  *int
	rating   vo.Rating
}

func NewUserRate(userId string, imageUrl string, ranking *int, rating vo.Rating) UserRate {
	return UserRate{userId: userId, imageUrl: imageUrl, ranking: ranking, rating: rating}
}

func (u UserRate) UserId() string {
	return u.userId
}

func (u UserRate) ImageUrl() string {
	return u.imageUrl
}

func (u UserRate) Ranking() *int {
	return u.ranking
}

func (u UserRate) Rating() vo.Rating {
	return u.rating
}
