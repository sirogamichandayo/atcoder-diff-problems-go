package infrastructure

import (
	"diff-problems/interfaces/commands"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "hoge",
}

var fetchAndStoreProblemDifficultiesCmd = &cobra.Command{
	Use: "fetch-and-store-problem-difficulties",
	Run: func(cmd *cobra.Command, args []string) {
		commands.FetchAndStoreProblemDifficulties()
	},
}

func Execute() {
	rootCmd.AddCommand(fetchAndStoreProblemDifficultiesCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
