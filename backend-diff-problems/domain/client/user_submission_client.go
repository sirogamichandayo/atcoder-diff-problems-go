//go:generate mockgen -source=$GOFILE -destination=./mock/mock_$GOFILE -package=$GOPACKAGE
package client

import "diff-problems/domain/entity"

type UserSubmissionClient interface {
	FetchSinceByEpochTime(int64) (entity.UserSubmissionList, error)
}
