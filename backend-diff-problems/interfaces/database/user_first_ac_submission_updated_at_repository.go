package database

type UserFirstAcSubmissionUpdatedAtRepository struct {
	SqlHandler
}

func (repo *UserFirstAcSubmissionUpdatedAtRepository) Update(epochTime int64) (err error) {
	err = repo.Transaction(func(handler TransactionHandler) error {
		_, err = handler.Execute("DELETE FROM user_first_ac_submission_updated_at")
		if err != nil {
			return err
		}
		_, err = handler.Execute("INSERT INTO user_first_ac_submission_updated_at VALUES (?)", epochTime)
		return err
	})
	return
}

func (repo *UserFirstAcSubmissionUpdatedAtRepository) Get() (updatedEpochTime int64, err error) {
	row, err := repo.Query("SELECT updated_epoch_time FROM user_first_ac_submission_updated_at")
	if err != nil {
		return
	}
	defer func(row Row) {
		err := row.Close()
		if err != nil {
			return
		}
	}(row)

	row.Next()
	if err = row.Scan(&updatedEpochTime); err != nil {
		return
	}

	return
}
