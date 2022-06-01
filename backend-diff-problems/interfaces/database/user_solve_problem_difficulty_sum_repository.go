package database

import (
	"diff-problems/domain/entity"
	"fmt"
)

type UserSolveProblemDifficultySumRepository struct {
	SqlHandler
}

func (u UserSolveProblemDifficultySumRepository) FindById(s string) (entity.UserSolveProblemDifficultySum, error) {
	var sEntity entity.UserSolveProblemDifficultySum

	row, err := u.Query("SELECT user_id, clip_difficulty_sum, rnk FROM user_solve_problem_difficulty_sum WHERE user_id = 'aaa' LIMIT 1;")
	if err != nil {
		return sEntity, err
	}
	defer row.Close()

	if !row.Next() {
		return sEntity, fmt.Errorf("not found user_id = %s", s)
	}
	err = row.Scan(&sEntity.UserId, &sEntity.ClipDifficultySum, &sEntity.Rank)
	return sEntity, err
}

func (u UserSolveProblemDifficultySumRepository) SortedAll() (*entity.SortedUserSolveProblemDifficultySumList, error) {
	rows, err := u.Query("SELECT user_id, clip_difficulty_sum, rnk FROM user_solve_problem_difficulty_sum ORDER BY rnk asc;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sList entity.UserSolveProblemDifficultySumList
	for rows.Next() {
		var sum entity.UserSolveProblemDifficultySum
		if err := rows.Scan(&sum.UserId, &sum.ClipDifficultySum, &sum.Rank); err != nil {
			return nil, err
		}
		sList = append(sList, sum)
	}
	return entity.NewSortedUserSolveProblemDifficultySumList(sList)
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
   clip_difficulty_sum = VALUES(clip_difficulty_sum), rnk = VALUES(rnk)
`,
			placeholderValue.Placeholder())
		_, err := u.Execute(query, placeholderValue.Values()...)

		if err != nil {
			return err
		}
	}

	return nil
}
