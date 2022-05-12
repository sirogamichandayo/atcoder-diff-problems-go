package repository

type UserFirstAcSubmissionUpdatedAtRepository interface {
	Get() (int64, error)
	Update(int64) error
}
