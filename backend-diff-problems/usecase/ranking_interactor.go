package usecase

type RankingInteractor struct {
	UserFirstAcSubmissionRepository          UserFirstAcSubmissionRepository
	UserFirstAcSubmissionUpdatedAtRepository UserFirstAcSubmissionUpdatedAtRepository
	ProblemDifficultyRepository              ProblemDifficultyRepository
	RankingRepository                        RankingRepository
	RankingUpdatedAtRepository               RankingUpdatedAtRepository
}

func (interactor *RankingInteractor) Update() error {
	submissionUpdatedAt, err := interactor.UserFirstAcSubmissionUpdatedAtRepository.Get()
	if err != nil {
		return err
	}

}

type RankingRepository interface {
	BulkUpsert()
}

type RankingUpdatedAtRepository interface {
	Update(uint) error
}
