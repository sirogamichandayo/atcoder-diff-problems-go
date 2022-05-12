package database

import (
	"diff-problems/domain/entity"
	"fmt"
)

type UserFirstAcSubmissionRepository struct {
	SqlHandler
}

func (repo *UserFirstAcSubmissionRepository) BulkUpsert(u entity.AcUserSubmissionList) (err error) {
	placeholders, values := u.MakeValueForUpsertMySql()

	// on duplicate key updateのときはfirst_solved_atが小さくなるような更新だけをする
	query := fmt.Sprintf(
		`
INSERT INTO user_first_ac_submissions (user_id, problem_id, first_solved_epoch_time)
VALUES %s
ON DUPLICATE KEY UPDATE
  user_id = VALUES(user_id),
  problem_id = VALUES(problem_id),
  first_solved_epoch_time = LEAST(first_solved_epoch_time, VALUES(first_solved_epoch_time))  
`,
		placeholders)
	_, err = repo.Execute(query, values...)

	return
}
