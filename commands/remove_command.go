package commands

import (
	"Terminal_Task_Managment_V2/messages"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
)

func init() {
	rootCmd.AddCommand(removeCommand)
}

var removeCommand = &cobra.Command{
	Use:   "remove",
	Short: "Delete a task",
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
			fmt.Println(messages.NoDirExist)
			return
		}

		readFile, err := os.ReadFile("./database/database.json")
		if err != nil {
			fmt.Println(messages.NoFileOrError)
			return
		}

		tasks := []TaskStruct{}
		err = json.Unmarshal(readFile, &tasks)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(messages.DefineIDTask)
		var id string

		_, err = fmt.Scan(&id)
		if err != nil {
			log.Fatalln(err)
		}

		strConv, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println(messages.ErrorNotIntTask)
			return
		}

		for i, v := range tasks {
			if v.Id != strConv {
				i++
			} else if v.Id == strConv {
				fmt.Println("Found !") //Remove Task
				return
			}
		}
	},
}
