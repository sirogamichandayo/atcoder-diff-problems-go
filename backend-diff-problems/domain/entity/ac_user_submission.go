package entity

import "fmt"

type AcUserSubmission struct {
	UserId    string
	ProblemId string
	EpochTime uint
}

func MakeAcUserSubmissionFromUserSubmission(rawSubmission UserSubmission) (
	acSubmission AcUserSubmission,
	err error,
) {
	if !rawSubmission.IsAc() {
		err = fmt.Errorf("this submission result is not ac")
		return
	}
	acSubmission = AcUserSubmission{
		UserId:    rawSubmission.UserId,
		ProblemId: rawSubmission.ProblemId,
		EpochTime: rawSubmission.EpochTime,
	}
	return
}

type AcUserSubmissionList []AcUserSubmission
