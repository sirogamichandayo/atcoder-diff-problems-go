package usecase

import "diff-problems/domain"

type UserRepository interface {
	Store(domain.User) (uint64, error)
	FindById(uint64) (domain.User, error)
	FindAll() (domain.Users, error)
}
