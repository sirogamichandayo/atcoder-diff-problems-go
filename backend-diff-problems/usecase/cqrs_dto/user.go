package cqrsDto

import "diff-problems/domain/vo"

type User struct {
	userId   string
	imageUrl string
	ranking  int
	rating   vo.Rating
}

func NewUser(userId string, imageUrl string, ranking int, rating vo.Rating) User {
	return User{userId: userId, imageUrl: imageUrl, ranking: ranking, rating: rating}
}

func (u User) UserId() string {
	return u.userId
}

func (u User) ImageUrl() string {
	return u.imageUrl
}

func (u User) Ranking() int {
	return u.ranking
}

func (u User) Rating() vo.Rating {
	return u.rating
}
