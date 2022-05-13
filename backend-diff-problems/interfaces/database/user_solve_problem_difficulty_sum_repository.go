package database

import (
	"diff-problems/domain/entity"
	"fmt"
)

type UserSolveProblemDifficultySumRepository struct {
	SqlHandler
}

func (u UserSolveProblemDifficultySumRepository) Upsert(list entity.UserSolveProblemDifficultySumList) error {
	placeholders, values := list.MakeValueForUpsertMySql()
	query := fmt.Sprintf(`
INSERT INTO user_solve_problem_difficulty_sum (user_id, clip_difficulty_sum)
VALUES %s
ON DUPLICATE KEY UPDATE
  user_id = VALUES(user_id)
`,
		placeholders)
	_, err := u.Execute(query, values)

	return err
}
