package commands

import (
	"Terminal_Task_Managment_V2/messages"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(helpCommand)
}

var helpCommand = &cobra.Command{
	Use:   "help",
	Short: "Allows you to see the list of commands",
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < 100; i++ {
			fmt.Println("")
		}

		fmt.Println(messages.HeaderHelpCommand + "\n" + messages.CommandsList)
	},
}
