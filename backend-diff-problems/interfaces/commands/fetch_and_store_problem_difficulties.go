package commands

import (
	"diff-problems/interfaces/api"
	"diff-problems/interfaces/api/atcoder_problems_api"
	"diff-problems/interfaces/database"
	"diff-problems/usecase"
)

type FetchAndStoreProblemDifficultyCommand struct {
	Interactor usecase.ProblemDifficultyInteractor
}

func NewFetchAndStoreProblemDifficultyCommand(sqlHandler database.SqlHandler, requestHandler api.RequestHandler) *FetchAndStoreProblemDifficultyCommand {
	return &FetchAndStoreProblemDifficultyCommand{
		Interactor: usecase.ProblemDifficultyInteractor{
			ProblemDifficultyRepository: &database.ProblemDifficultyRepository{
				SqlHandler: sqlHandler,
			},
			ProblemDifficultyAtcoderProblemClient: &atcoder_problems_api.ProblemDifficultyClient{
				RequestHandler: requestHandler,
			},
		},
	}
}

func (command *FetchAndStoreProblemDifficultyCommand) Exec() {
	err := command.Interactor.FetchAndStore()
	if err != nil {

	}
}
