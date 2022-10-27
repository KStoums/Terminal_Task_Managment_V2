package commands

import (
	"Terminal_Task_Managment_V2/messages"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(optionsCommand)
}

var optionsCommand = &cobra.Command{
	Use:   "options",
	Short: "Change some Terminal Task Management options",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println(messages.IncorrectSyntax)
			return
		}

		for i := 0; i < 100; i++ {
			fmt.Println("")
		}

		fmt.Println(messages.CommandOptionsSoon)
	},
}
