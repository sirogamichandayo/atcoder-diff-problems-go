package usecase

import "diff-problems/domain/entity"

type UserRepository interface {
	Store(entity.User) (uint64, error)
	FindById(uint64) (entity.User, error)
	FindAll() (entity.Users, error)
}
