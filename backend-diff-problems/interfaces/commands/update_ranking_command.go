package commands

import (
	"diff-problems/interfaces/api"
	"diff-problems/interfaces/database"
	"diff-problems/usecase"
)

type UpdateRankingCommand struct {
	Interactor usecase.RankingInteractor
}

func NewUpdateRankingCommand(
	sqlHandler database.SqlHandler,
	requestHandler api.RequestHandler,
) *UpdateRankingCommand {
	return &UpdateRankingCommand{
		Interactor: usecase.RankingInteractor{
			UserFirstAcSubmissionRepository: &database.UserFirstAcSubmissionRepository{
				SqlHandler: sqlHandler,
			},
			UserFirstAcSubmissionUpdatedAtRepository: &database.UserFirstAcSubmissionUpdatedAtRepository{
				SqlHandler: sqlHandler,
			},
			ProblemDifficultyRepository: &database.ProblemDifficultyRepository{
				SqlHandler: sqlHandler,
			},
			/*
				RankingRepository: &database.RankingRepository{
					SqlHandler: sqlHandler,
				},
				RankingUpdatedAtRepository: &database.RankingUpdatedAtRepository{
					SqlHandler: sqlHandler,
				},*

			*/
		},
	}
}

func (command *UpdateRankingCommand) Update() (err error) {
	return command.Interactor.Update()
}
