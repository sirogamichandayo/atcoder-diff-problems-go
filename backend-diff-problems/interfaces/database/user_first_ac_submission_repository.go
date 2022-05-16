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

	fmt.Println(values)
	// on duplicate key updateのときはfirst_solved_atが小さくなるような更新だけをする
	// TODO: first_solved_epoch_timeあたりをアプリケーションで制御するようにしたい
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

	rows, err := repo.Query(
		`select count(*) from user_first_ac_submissions order by user_id, problem_id;`,
	)
	defer rows.Close()
	rows.Next()
	var a int
	rows.Scan(&a)
	fmt.Println("count: ", a)

	return
}
