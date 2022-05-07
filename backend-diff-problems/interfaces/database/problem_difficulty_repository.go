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
INSERT INTO product_difficulties (problem_id, difficulty)
VALUES %s
ON DUPLICATE KEY UPDATE problem_id = VALUES(problem_id);
`,
		placeholders)
	_, err = repo.Execute(query, values...)

	return
}
