package commands

import (
	"diff-problems/interfaces/api"
	"diff-problems/interfaces/api/atcoder_problems_api"
	"diff-problems/interfaces/database"
	"diff-problems/usecase"
)

type UpdateUserFirstAcSubmissionCommand struct {
	Interactor usecase.UserFirstAcSubmissionInteractor
}

func NewUpdateUserFirstAcSubmissionCommand(
	sqlHandler database.SqlHandler,
	requestHandler api.RequestHandler,
) *UpdateUserFirstAcSubmissionCommand {
	return &UpdateUserFirstAcSubmissionCommand{
		Interactor: usecase.UserFirstAcSubmissionInteractor{
			UserFirstAcSubmissionRepository: &database.UserFirstAcSubmissionRepository{
				SqlHandler: sqlHandler,
			},
			UserFirstAcSubmissionUpdatedAtRepository: &database.UserFirstAcSubmissionUpdatedAtRepository{
				SqlHandler: sqlHandler,
			},
			UserSubmissionAtCoderProblemClient: &atcoder_problems_api.UserSubmissionClient{
				RequestHandler: requestHandler,
			},
		},
	}
}

func (command *UpdateUserFirstAcSubmissionCommand) UpdateAll() (err error) {
	return command.Interactor.UpdateAll()
}

func (command *UpdateUserFirstAcSubmissionCommand) UpdateFromUpdatedAt() (err error) {
	return command.Interactor.UpdateFromUpdatedAt()
}
