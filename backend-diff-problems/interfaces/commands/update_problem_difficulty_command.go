package commands

import (
	"diff-problems/interfaces/api"
	"diff-problems/interfaces/api/atcoder_problems_api"
	"diff-problems/interfaces/database"
	"diff-problems/usecase"
)

type UpdateProblemDifficultyCommand struct {
	Interactor usecase.ProblemDifficultyInteractor
}

func NewUpdateProblemDifficultyCommand(
	sqlHandler database.SqlHandler,
	requestHandler api.RequestHandler,
) *UpdateProblemDifficultyCommand {
	return &UpdateProblemDifficultyCommand{
		Interactor: usecase.ProblemDifficultyInteractor{
			ProblemDifficultyRepository: &database.ProblemDifficultyRepository{
				SqlHandler: sqlHandler,
			},
			ProblemDifficultyAtCoderProblemClient: &atcoder_problems_api.ProblemDifficultyClient{
				RequestHandler: requestHandler,
			},
		},
	}
}

func (command *UpdateProblemDifficultyCommand) Exec() error {
	return command.Interactor.Update()
}
