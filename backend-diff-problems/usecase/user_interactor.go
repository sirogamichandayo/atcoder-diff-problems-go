package usecase

import "diff-problems/domain/entity"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u entity.User) (err error) {
	_, err = interactor.UserRepository.Store(u)
	return
}

func (interactor *UserInteractor) Users() (user entity.Users, err error) {
	user, err = interactor.UserRepository.FindAll()
	return
}

func (interactor *UserInteractor) UserById(identifier uint64) (user entity.User, err error) {
	user, err = interactor.UserRepository.FindById(identifier)
	return
}
