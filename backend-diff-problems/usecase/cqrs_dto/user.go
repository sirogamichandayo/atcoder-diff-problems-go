package cqrsDto

import "diff-problems/domain/vo"

type User struct {
	userId   string
	imageUrl string
	rating   vo.Rating
}

func NewUser(userId string, imageUrl string, rating vo.Rating) User {
	return User{userId: userId, imageUrl: imageUrl, rating: rating}
}

func (u User) UserId() string {
	return u.userId
}

func (u User) ImageUrl() string {
	return u.imageUrl
}

func (u User) Rating() vo.Rating {
	return u.rating
}
