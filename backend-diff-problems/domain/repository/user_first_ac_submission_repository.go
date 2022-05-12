package repository

import "diff-problems/domain/entity"

type UserFirstAcSubmissionRepository interface {
	BulkUpsert(entity.AcUserSubmissionList) error
}
