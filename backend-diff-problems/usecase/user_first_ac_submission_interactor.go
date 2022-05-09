package usecase

import "diff-problems/domain/entity"

type UserFirstAcSubmissionInteractor struct {
	UserFirstAcSubmissionRepository    UserFirstAcSubmissionRepository
	UserSubmissionAtcoderProblemClient UserSubmissionAtCoderProblemClient
}

func (interactor *UserFirstAcSubmissionInteractor) update(sinceEpochTime uint) (
	lastEpochTime uint,
	err error,
) {
	userSubmissionList, err := interactor.UserSubmissionAtcoderProblemClient.Fetch(sinceEpochTime)
	if err != nil {
		return
	}
	userAcSubmissionList, err := userSubmissionList.ExactByAc()
	if err != nil {
		return
	}
	err = interactor.UserFirstAcSubmissionRepository.BulkUpsert(userAcSubmissionList)
	if err != nil {
		return
	}
	return userSubmissionList.LastEpochTime(), nil
}

type UserFirstAcSubmissionRepository interface {
	BulkUpsert(u entity.AcUserSubmissionList) (err error)
}

type UserSubmissionAtCoderProblemClient interface {
	Fetch(sinceEpochTime uint) (u entity.UserSubmissionList, err error)
}
