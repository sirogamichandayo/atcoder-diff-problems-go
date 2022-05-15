package database

import (
	"diff-problems/domain/entity"
	"fmt"
)

type UserSolveProblemDifficultySumRepository struct {
	SqlHandler
}

func (u UserSolveProblemDifficultySumRepository) Upsert(list entity.UserSolveProblemDifficultySumList) error {
	placeholderValueList, err := list.MakeValuesForUpsertMySql()
	if err != nil {
		return err
	}

	for _, placeholderValue := range placeholderValueList {
		query := fmt.Sprintf(`
INSERT INTO user_solve_problem_difficulty_sum (user_id, clip_difficulty_sum, rnk)
VALUES %s
ON DUPLICATE KEY UPDATE
  user_id = VALUES(user_id)
`,
			placeholderValue.Placeholder())
		_, err := u.Execute(query, placeholderValue.Values()...)

		if err != nil {
			return err
		}
	}

	return nil
}
