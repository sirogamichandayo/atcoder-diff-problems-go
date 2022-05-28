//go:generate mockgen -source=$GOFILE -destination=./mock/mock_$GOFILE -package=$GOPACKAGE
package client

import "diff-problems/domain/vo"

type ContestResultClient interface {
	All(string) (list vo.ContestResultList, err error)
}