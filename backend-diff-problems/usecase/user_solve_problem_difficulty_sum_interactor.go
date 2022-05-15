package usecase

import (
	"diff-problems/domain/repository"
	"diff-problems/domain/repository/service"
)

type UserSolveProblemDifficultySumInteractor struct {
	UserFirstAcSubmissionUpdatedAtRepository         repository.UserFirstAcSubmissionUpdatedAtRepository
	CalcUserSolveProblemDifficultySumService         service.CalcUserSolveProblemDifficultySumService
	UserSolveProblemDifficultySumRepository          repository.UserSolveProblemDifficultySumRepository
	UserSolveProblemDifficultySumUpdatedAtRepository repository.UserSolveProblemDifficultySumUpdatedAtRepository
}

func (interactor UserSolveProblemDifficultySumInteractor) Update() error {
	submissionUpdatedAt, err := interactor.UserFirstAcSubmissionUpdatedAtRepository.Get()
	if err != nil {
		return err
	}

	diffSumList, err := interactor.CalcUserSolveProblemDifficultySumService.Calc()
	if err != nil {
		return err
	}

	err = interactor.UserSolveProblemDifficultySumRepository.Upsert(diffSumList)
	if err != nil {
		return err
	}

	return interactor.UserSolveProblemDifficultySumUpdatedAtRepository.Update(
		submissionUpdatedAt,
	)
}
