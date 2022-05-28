//go:generate mockgen -source=$GOFILE -destination=./mock/mock_$GOFILE -package=$GOPACKAGE
package cqrsService

import cqrsDto "diff-problems/usecase/cqrs_dto"

type UserProblemService interface {
	FindByUserId(string) (cqrsDto.UserProblem, error)
}
