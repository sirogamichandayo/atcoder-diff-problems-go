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

<<<<<<< Updated upstream
=======
var updateAllUserFirstAcSubmissionCmd = &cobra.Command{
	Use:   "update-all-ac-submission",
	Short: "全てのatcoderの最初のAC提出を更新",
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

var startApiCmd = &cobra.Command{
	Use:   "api",
	Short: "apiを起動",
	Long:  "atcoder diff problemsのapiを起動します",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		RouterInitialize()
		return Router.Run()
	},
}

var updateUserFirstAcSubmissionCmd = &cobra.Command{
	Use:   "update-ac-submission",
	Short: "すでに更新済みの部分からatcoderの最初のAC提出を更新",
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

>>>>>>> Stashed changes
func Execute() {
	rootCmd.AddCommand(updateProblemDifficultiesCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
