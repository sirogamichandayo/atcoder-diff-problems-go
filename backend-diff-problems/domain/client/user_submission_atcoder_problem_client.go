package client

import "diff-problems/domain/entity"

type UserSubmissionAtCoderProblemClient interface {
	Fetch(int64) (entity.UserSubmissionList, error)
}
