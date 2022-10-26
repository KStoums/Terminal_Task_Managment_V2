package commands

import (
	"Terminal_Task_Managment_V2/messages"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
	"log"
	"os"
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

		for i := 0; i < 100; i++ {
			fmt.Println("")
		}

		_, err := os.Stat("./database")
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println(messages.NoDirExist)
				return
			}

			fmt.Println(err)
			return
		}

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
