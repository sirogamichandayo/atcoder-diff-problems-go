package entity

import "diff-problems/domain/vo"

type UserSubmission struct {
	UserId    string
	ProblemId string
	Result    vo.Result
	EpochTime uint
}

func (s UserSubmission) IsAc() bool {
	return s.Result.IsAc()
}

type UserSubmissionList []UserSubmission

func (list *UserSubmissionList) ExactByAc() (acList AcUserSubmissionList, err error) {
	for _, rawSubmission := range *list {
		if rawSubmission.IsAc() {
			newSubmission, err := MakeAcUserSubmissionFromUserSubmission(rawSubmission)
			if err != nil {
				return
			}
			acList = append(acList, newSubmission)
		}
	}
	return
}

func (list *UserSubmissionList) LastEpochTime() (res uint) {
	res = 0
	for _, rawSubmission := range *list {
		if rawSubmission.EpochTime > res {
			res = rawSubmission.EpochTime
		}
	}
	return
}

func MakeUserSubmissionListFromJsonBytes(bytes []byte) (list UserSubmissionList, err error) {

}
