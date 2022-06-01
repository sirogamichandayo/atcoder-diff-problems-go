package usecase

import (
	"diff-problems/domain/client"
	"diff-problems/domain/entity"
	"diff-problems/domain/repository"
	"fmt"
	"time"
)

type UserFirstAcSubmissionInteractor struct {
	UserFirstAcSubmissionRepository          repository.UserFirstAcSubmissionRepository
	UserFirstAcSubmissionUpdatedAtRepository repository.UserFirstAcSubmissionUpdatedAtRepository
	UserSubmissionAtCoderProblemClient       client.UserSubmissionClient
}

// UpdateAll は最初から最後まで更新します
func (interactor UserFirstAcSubmissionInteractor) UpdateAll() (err error) {
	err = interactor.updateToTheEnd(0)
	return
}

// UpdateFromUpdatedAt は更新済みの時間から、提出がなくなるまで更新します
func (interactor UserFirstAcSubmissionInteractor) UpdateFromUpdatedAt() (err error) {
	updatedEpochTime, err := interactor.UserFirstAcSubmissionUpdatedAtRepository.Get()
	if err != nil {
		return
	}
	err = interactor.updateToTheEnd(updatedEpochTime)
	return
}

// update 与えられたepoch_timeから提出がなくなるまで更新します
func (interactor UserFirstAcSubmissionInteractor) updateToTheEnd(updatedEpochTime int64) (err error) {
	var isLast bool
	for {
		updatedEpochTime, isLast, err = interactor.fetchSubmissionAndUpdate(updatedEpochTime + 1)
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
func (interactor UserFirstAcSubmissionInteractor) fetchSubmissionAndUpdate(sinceEpochTime int64) (
	lastEpochTime int64,
	isLast bool,
	err error,
) {
	retryCount := 10
	var userSubmissionList entity.UserSubmissionList
	for true {
		userSubmissionList, err = interactor.UserSubmissionAtCoderProblemClient.FetchSinceByEpochTime(sinceEpochTime)
		if retryCount == 0 {
			return
		}
		if err != nil {
			fmt.Errorf(err.Error())
			retryCount--
		} else {
			break
		}
	}

	if userSubmissionList.IsEmpty() {
		return sinceEpochTime, true, nil
	}
	userAcSubmissionList, err := userSubmissionList.ExactByAc()
	if err != nil {
		return
	}

	if userAcSubmissionList.IsEmpty() {
		return userSubmissionList.LastEpochTime(), false, nil
	}
	err = interactor.UserFirstAcSubmissionRepository.BulkUpsert(userAcSubmissionList)
	if err != nil {
		return
	}

	lastEpochTime = userSubmissionList.LastEpochTime()

	err = interactor.UserFirstAcSubmissionUpdatedAtRepository.Update(lastEpochTime)

	return lastEpochTime, false, err
}
