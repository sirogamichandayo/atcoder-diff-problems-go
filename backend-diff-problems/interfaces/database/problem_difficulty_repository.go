package database

import (
	"diff-problems/domain/entity"
	"fmt"
)

type ProblemDifficultyRepository struct {
	SqlHandler
}

func (repo *ProblemDifficultyRepository) BulkUpsert(u entity.ProblemDifficultyList) (err error) {
	placeholders, values := u.MakeValueForUpsertMySql()
	query := fmt.Sprintf(
		`
INSERT INTO problem_difficulties (problem_id, difficulty, clip_difficulty)
VALUES %s
ON DUPLICATE KEY UPDATE difficulty = VALUES(difficulty), clip_difficulty = VALUES(clip_difficulty);
`,
		placeholders)
	_, err = repo.Execute(query, values...)

	return
}
