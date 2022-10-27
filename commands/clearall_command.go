package commands

import (
	"Terminal_Task_Managment_V2/functions"
	"Terminal_Task_Managment_V2/messages"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
	"time"
)

func init() {
	rootCmd.AddCommand(clearAllCommand)
}

var clearAllCommand = &cobra.Command{
	Use:   "clearall",
	Short: "Delete the \"./database/*\" file",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println(messages.IncorrectSyntax)
			return
		}

		functions.ClearTerminal()

		_, err := os.Stat("./database")
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println(messages.NoDirOrNoPermission)
				return
			}

			fmt.Println(err)
			return
		}

		for {
			fmt.Println(messages.ConfirmClearAll)
			var confirm string
			fmt.Scan(&confirm)

			if strings.EqualFold(confirm, "n") {
				functions.ClearTerminal()
				fmt.Print(messages.ClearAllCanceled)
				return
			}

			if strings.EqualFold(confirm, "y") {
				break
			}

			fmt.Println(messages.NotGoodResponse)
		}

		functions.ClearTerminal()

		fmt.Println(messages.DeletingDir)
		err = os.RemoveAll("./database")
		if err != nil {
			log.Fatalln(err)
		}

		time.Sleep(2 * time.Second)
		barDeleteDir := progressbar.Default(100, "Deleting")

		for i := 0; i < 100; i++ {
			barDeleteDir.Add(50)
		}
		time.Sleep(1 * time.Second)
		fmt.Println(messages.DirDeleted)
	},
}
