package usecase

import "diff-problems/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) (err error) {
	_, err = interactor.UserRepository.Store(u)
	return
}

func (interactor *UserInteractor) Users() (user domain.Users, err error) {
	user, err = interactor.UserRepository.FindAll()
	return
}

func (interactor *UserInteractor) UserById(identifier uint64) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindById(identifier)
	return
}
