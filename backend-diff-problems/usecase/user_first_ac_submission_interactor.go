package usecase

import (
	"diff-problems/domain/entity"
	"fmt"
	"time"
)

type UserFirstAcSubmissionInteractor struct {
	UserFirstAcSubmissionRepository          UserFirstAcSubmissionRepository
	UserFirstAcSubmissionUpdatedAtRepository UserFirstAcSubmissionUpdatedAtRepository
	UserSubmissionAtcoderProblemClient       UserSubmissionAtCoderProblemClient
}

// UpdateAll は最初から最後まで更新します
func (interactor *UserFirstAcSubmissionInteractor) UpdateAll() (err error) {
	err = interactor.updateToTheEnd(0)
	return
}

// UpdateFromUpdatedAt は更新済みの時間から、提出がなくなるまで更新します
func (interactor *UserFirstAcSubmissionInteractor) UpdateFromUpdatedAt() (err error) {
	updatedEpochTime, err := interactor.UserFirstAcSubmissionUpdatedAtRepository.Get()
	if err != nil {
		return
	}
	err = interactor.updateToTheEnd(updatedEpochTime)
	return
}

// update 与えられたepoch_timeから提出がなくなるまで更新します
func (interactor *UserFirstAcSubmissionInteractor) updateToTheEnd(updatedEpochTime int64) (err error) {
	var isLast bool
	for {
		updatedEpochTime, isLast, err = interactor.fetchSubmissionAndUpdate(updatedEpochTime)
		fmt.Println("updated epoch time : ", updatedEpochTime)
		if err != nil {
			return
		}
		if isLast {
			break
		}
		time.Sleep(time.Second)
	}
	return
}

// fetchSubmissionAndUpdate はsinceEpochTimeからの提出をapiで取得して時間の最も早い提出を保存し、更新した最後の時間を保存します
func (interactor *UserFirstAcSubmissionInteractor) fetchSubmissionAndUpdate(sinceEpochTime int64) (
	lastEpochTime int64,
	isLast bool,
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

	// isEmptyがtrueのときにlastEpochTimeに0が入るのを防ぐif文
	isEmpty := userSubmissionList.IsEmpty()
	if isEmpty {
		lastEpochTime = sinceEpochTime
	} else {
		lastEpochTime = userSubmissionList.LastEpochTime()
	}
	err = interactor.UserFirstAcSubmissionUpdatedAtRepository.Update(lastEpochTime)

	return
}

type UserFirstAcSubmissionUpdatedAtRepository interface {
	Get() (int64, error)
	Update(int64) error
}

type UserFirstAcSubmissionRepository interface {
	BulkUpsert(entity.AcUserSubmissionList) error
}

type UserSubmissionAtCoderProblemClient interface {
	Fetch(int64) (entity.UserSubmissionList, error)
}
