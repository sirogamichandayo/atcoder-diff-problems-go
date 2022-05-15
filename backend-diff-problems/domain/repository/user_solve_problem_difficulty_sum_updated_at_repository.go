//go:generate mockgen -source=$GOFILE -destination=./mock/mock_$GOFILE -package=$GOPACKAGE
package repository

type UserSolveProblemDifficultySumUpdatedAtRepository interface {
	Update(int64) error
	Get() (int64, error)
}
