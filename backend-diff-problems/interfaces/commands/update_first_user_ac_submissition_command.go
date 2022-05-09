package commands

import (
	"diff-problems/interfaces/api"
	"diff-problems/interfaces/database"
	"diff-problems/usecase"
)

type UpdateUserFirstAcSubmissionCommand struct {
	Interactor usecase.FirstUserAcSubmissionInteractor
}

func NewUpdateUserFirstAcSubmissionCommand(
	sqlHandler database.SqlHandler,
	requestHandler api.RequestHandler,
) *UpdateUserFirstAcSubmissionCommand {
	return &UpdateUserFirstAcSubmissionCommand{
		Interactor: usecase.UpdateUserFirstAcSubmissionInteractor{
			UserFirstAcSubmissionRepository: &database.UserFirstAcSubmissionRepository{
				SqlHandler: sqlHandler,
			},
			UserSubmissionRepository: &database.UserSubmissionRepository{
				RequestHandler: requestHandler,
			},
		},
	}
}

func (command *UpdateUserFirstAcSubmissionCommand) Exec() (err error) {

}
