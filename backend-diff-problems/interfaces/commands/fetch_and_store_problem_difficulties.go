package commands

import "diff-problems/interfaces/database"

type FetchAndStoreProblemDifficultyCommand struct {
	Interactor usecase.FetchAndStoreProblemDifficultyInteractor
}

func NewFetchAndStoreProblemDifficultyCommand(sqlHandler database.SqlHandler) *FetchAndStoreProblemDifficultyCommand {
	return FetchAndStoreProblemDifficultyCommand{
		Interactor: usecase.FetchAndStoreProblemDifficultyInteractor{
			ProblemDifficultyRepository: &database.ProblemDifficultyRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (command *FetchAndStoreProblemDifficultyCommand) Exec() {
	err := command.Interactor.Exec()
	if err != nil {

	}
}
