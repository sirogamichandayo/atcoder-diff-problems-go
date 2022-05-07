package infrastructure

import (
	conf "diff-problems/config"
	"diff-problems/interfaces/commands"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "hoge",
}

var fetchAndStoreProblemDifficultiesCmd = &cobra.Command{
	Use:   "fetch-and-store-problem-difficulties",
	Short: "atcoder problems apiを叩いて問題のdiffを永続化する",
	Run: func(cmd *cobra.Command, args []string) {
		cDir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		config, err := conf.LoadConfig(cDir)
		if err != nil {
			panic(err)
		}

		command := commands.NewFetchAndStoreProblemDifficultyCommand(
			NewSqlHandler(config.SinDb),
			NewRequestHandler(),
		)
		command.Exec()
	},
}

func Execute() {
	rootCmd.AddCommand(fetchAndStoreProblemDifficultiesCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
