package client

import "diff-problems/domain/entity"

type UserSubmissionClient interface {
	FetchSinceByEpochTime(int64) (entity.UserSubmissionList, error)
}
