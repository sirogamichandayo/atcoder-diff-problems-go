//go:generate mockgen -source=$GOFILE -destination=./mock/mock_$GOFILE -package=$GOPACKAGE
package client

import "diff-problems/domain/entity"

type ProblemDifficultyClient interface {
	Fetch() (list entity.ProblemDifficultyList, err error)
}
