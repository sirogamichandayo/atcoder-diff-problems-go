package commands

import (
	"diff-problems/interfaces/database"
	"diff-problems/interfaces/database/service"
	"diff-problems/usecase"
)

type UserSolveProblemDifficultySumCommand struct {
	Interactor usecase.UserSolveProblemDifficultySumInteractor
}

func NewUserSolveProblemDifficultySumCommand(
	sqlHandler database.SqlHandler,
) *UserSolveProblemDifficultySumCommand {
	return &UserSolveProblemDifficultySumCommand{
		Interactor: usecase.UserSolveProblemDifficultySumInteractor{
			UserFirstAcSubmissionUpdatedAtRepository: &database.UserFirstAcSubmissionUpdatedAtRepository{
				SqlHandler: sqlHandler,
			},
			CalcUserSolveProblemDifficultySumService: &service.CalcUserSolveProblemDifficultySumService{
				SqlHandler: sqlHandler,
			},
			UserSolveProblemDifficultySumRepository: &database.UserSolveProblemDifficultySumRepository{
				SqlHandler: sqlHandler,
			},
			UserSolveProblemDifficultySumUpdatedAtRepository: &database.UserSolveProblemDifficultySumUpdatedAtRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (command *UserSolveProblemDifficultySumCommand) Update() (err error) {
	return command.Interactor.Update()
}
