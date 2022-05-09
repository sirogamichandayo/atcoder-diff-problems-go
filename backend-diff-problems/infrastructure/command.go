package infrastructure

import (
	conf "diff-problems/config"
	"diff-problems/interfaces/commands"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "AtCoderDiffProblemsCli",
}

var updateProblemDifficultiesCmd = &cobra.Command{
	Use:   "update-problem-difficulties",
	Short: "atcoderの問題のdiffを更新",
	Long:  "atcoder problemsのapiを叩いて、得られた情報ををDBに保存します",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		config, err := conf.LoadConfig()
		if err != nil {
			return
		}

		command := commands.NewUpdateProblemDifficultyCommand(
			NewSqlHandler(config.SinDb),
			NewRequestHandler(),
		)
		err = command.Exec()
		return
	},
}

func Execute() {
	rootCmd.AddCommand(updateProblemDifficultiesCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
