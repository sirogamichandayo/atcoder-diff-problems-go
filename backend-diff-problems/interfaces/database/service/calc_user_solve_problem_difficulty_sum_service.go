package service

import (
	"diff-problems/domain/entity"
	"diff-problems/interfaces/database"
)

// CalcUserSolveProblemDifficultySumService はSQLでユーザーが解いた問題のdiffを計算するstructです
type CalcUserSolveProblemDifficultySumService struct {
	database.SqlHandler
}

func (repo CalcUserSolveProblemDifficultySumService) Calc() (entity.UserSolveProblemDifficultySumList, error) {
	// NOTE: アプリケーションコードでは計算時間がかかりすぎるのでSQLで書いた
	query := `
SELECT user_id, clip_diff_sum, RANK() OVER(order by clip_diff_sum desc) 'rank'
from   
(
 SELECT user_id, SUM(IFNULL(clip_difficulty, 0)) as clip_diff_sum 
 FROM user_first_ac_submissions
 JOIN problem_difficulties
 USING (problem_id)
 GROUP BY user_id
)t;
`
	rows, err := repo.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sumEntityList := make(entity.UserSolveProblemDifficultySumList, 0)

	var userId string
	var clipDifficultySum float64
	var rank uint64

	for rows.Next() {
		if err := rows.Scan(&userId, &clipDifficultySum, &rank); err != nil {
			return nil, err
		}

		sumEntity := entity.UserSolveProblemDifficultySum{
			UserId: userId, ClipDifficultySum: clipDifficultySum, Rank: rank,
		}
		sumEntityList = append(sumEntityList, sumEntity)
	}

	return sumEntityList, nil
}
