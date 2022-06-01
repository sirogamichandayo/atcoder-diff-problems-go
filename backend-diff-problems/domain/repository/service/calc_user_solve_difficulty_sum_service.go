//go:generate mockgen -source=$GOFILE -destination=./mock/mock_$GOFILE -package=$GOPACKAGE
package service

import "diff-problems/domain/entity"

type CalcUserSolveProblemDifficultySumService interface {
	Calc() (entity.UserSolveProblemDifficultySumList, error)
}
