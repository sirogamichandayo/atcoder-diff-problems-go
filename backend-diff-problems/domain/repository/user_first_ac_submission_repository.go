//go:generate mockgen -source=$GOFILE -destination=./mock/mock_$GOFILE -package=$GOPACKAGE
package repository

import "diff-problems/domain/entity"

type UserFirstAcSubmissionRepository interface {
	BulkUpsert(entity.AcUserSubmissionList) error
}
