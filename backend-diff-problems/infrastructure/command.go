package infrastructure

import (
	conf "diff-problems/config"
	"diff-problems/interfaces/commands"
	"os"

	"github.com/spf13/cobra"
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

var updateAllUserFirstAcSubmissionCmd = &cobra.Command{
	Use:   "update-all-ac-submission",
	Short: "atcoderの最初のAC提出を更新",
	Long:  "atcoder problemsのapiを叩いて、得られた情報を元にDBに保存します",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		config, err := conf.LoadConfig()
		if err != nil {
			return
		}

		command := commands.NewUpdateUserFirstAcSubmissionCommand(
			NewSqlHandler(config.SinDb),
			NewRequestHandler(),
		)
		err = command.UpdateAll()
		return
	},
}

var updateDifficultySumCmd = &cobra.Command{
	Use:   "update-difficulty-sum",
	Short: "ランキング情報をupdate",
	Long:  "すでにクロール済みの情報からランキングを更新します",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		config, err := conf.LoadConfig()
		if err != nil {
			return
		}

		command := commands.NewUserSolveProblemDifficultySumCommand(
			NewSqlHandler(config.SinDb),
		)
		err = command.Update()
		return
	},
}

var startApiCmd = &cobra.Command{
	Use:   "api",
	Short: "apiを起動",
	Long:  "atcoder diff problemsのapiを起動します",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		return RouterInitialize().Run()
	},
}

var updateUserFirstAcSubmissionCmd = &cobra.Command{
	Use:   "update-ac-submission",
	Short: "atcoderの最初のAC提出を更新",
	Long:  "atcoder problemsのapiを叩いて、得られた情報を元にDBに保存します",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		config, err := conf.LoadConfig()
		if err != nil {
			return
		}

		command := commands.NewUpdateUserFirstAcSubmissionCommand(
			NewSqlHandler(config.SinDb),
			NewRequestHandler(),
		)
		err = command.UpdateFromUpdatedAt()
		return
	},
}

func Execute() {
	rootCmd.AddCommand(updateProblemDifficultiesCmd)
	rootCmd.AddCommand(updateAllUserFirstAcSubmissionCmd)
	rootCmd.AddCommand(updateUserFirstAcSubmissionCmd)
	rootCmd.AddCommand(updateDifficultySumCmd)
	rootCmd.AddCommand(startApiCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
