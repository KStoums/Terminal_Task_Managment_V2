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
			functions.ClearTerminal()
			fmt.Println(messages.IncorrectSyntax)
			return
		}

		functions.ClearTerminal()
		checkDirDatabase := functions.CheckDirDatabase()
		if checkDirDatabase == false {
			return
		}

		checkDatabase := functions.CheckDatabase()
		if checkDatabase == false {
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
		err := os.RemoveAll("./database")
		if err != nil {
			functions.ClearTerminal()
			log.Fatalln(err)
		}

		time.Sleep(2 * time.Second)
		barDeleteDir := progressbar.Default(100, "Deleting")

		for i := 0; i < 100; i++ {
			barDeleteDir.Add(50)
		}
		time.Sleep(1 * time.Second)
		functions.ClearTerminal()
		fmt.Println(messages.DirDeleted)
	},
}
