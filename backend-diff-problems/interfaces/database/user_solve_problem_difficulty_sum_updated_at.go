package database

type UserSolveProblemDifficultySumUpdatedAtRepository struct {
	SqlHandler
}

func (repo UserSolveProblemDifficultySumUpdatedAtRepository) Update(epochTime int64) error {
	return repo.Transaction(func(handler TransactionHandler) error {
		if _, err := handler.Execute(
			"DELETE FROM user_solve_problem_difficulty_sum_updated_at",
		); err != nil {
			return err
		}

		if _, err := handler.Execute(
			"INSERT INTO user_solve_problem_difficulty_sum_updated_at VALUES (?)",
			epochTime,
		); err != nil {
			return err
		}

		return nil
	})
}