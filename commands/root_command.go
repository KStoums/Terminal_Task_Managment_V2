package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "ttm",
	Short: "Terminal Task Management is a program allowing you to list edit_tasks.",
	Long: `Terminal Task Management is a program allowing you to list edit_tasks with a definable status and priority. 
			An email is sent to you if the task arrives at the time defined by you.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
