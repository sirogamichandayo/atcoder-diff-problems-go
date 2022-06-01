//go:generate mockgen -source=$GOFILE -destination=./mock/mock_$GOFILE -package=$GOPACKAGE
package repository

type UserFirstAcSubmissionUpdatedAtRepository interface {
	Get() (int64, error)
	Update(int64) error
}
