package database

import "diff-problems/domain/entity"

// CalcUserSolveProblemDifficultySumService はSQLでユーザーが解いた問題のdiffを計算するstructです
type CalcUserSolveProblemDifficultySumService struct {
	SqlHandler
}

func (repo CalcUserSolveProblemDifficultySumService) Calc() (entity.UserSolveProblemDifficultySumList, error) {
	// NOTE: アプリケーションコードでは計算時間がかかりすぎるのでSQLで書いた
	query := `
SELECT user_id, SUM(IFNULL(clip_difficulty, 0)) as clip_diff_sum 
FROM 'user_first_ac_submissions''
JOIN 'problem_difficulties'
USING ('problem_id')
GROUP BY 'user_id'
`
	rows, err := repo.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sumEntityList := make(entity.UserSolveProblemDifficultySumList, 0)

	var userId string
	var clipDifficultySum float64

	for rows.Next() {
		if err := rows.Scan(&userId, &clipDifficultySum); err != nil {
			return nil, err
		}

		sumEntity := entity.UserSolveProblemDifficultySum{
			UserId: userId, ClipDifficultySum: clipDifficultySum,
		}
		sumEntityList = append(sumEntityList, sumEntity)
	}

	return sumEntityList, nil
}
