package commands

import (
	"Terminal_Task_Managment_V2/messages"
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
	"strings"
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

		//##### TABLEUR #####//
		data := [][]string{}

		for _, task := range tasks {
			data = append(data, []string{fmt.Sprintf("%d", task.Id), task.Name, task.Status, task.Priority,
				fmt.Sprintf("%d", task.Due), task.DateCreated.Format("2006-01-02 15:04:05")})
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "NAME", "STATUS", "PRIORITY", "DUE", "CREATION"})

		for _, v := range data {
			table.Append(v)
		}
		table.Render()
		//##########//

		fmt.Print("\n" + messages.DefineIDTask)
		var id string

		_, err = fmt.Scan(&id)
		if err != nil {
			log.Fatalln(err)
		}

		if strings.EqualFold(id, "cancel") {
			for i := 0; i < 100; i++ {
				fmt.Println("")
			}

			fmt.Print(messages.RemoveTaskCanceled)
			return
		}

		strConv, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println(messages.ErrorNotIntTask)
			return
		}

		var found bool
		for i, v := range tasks {
			if v.Id == strConv {
				tasks = append(tasks[:i], tasks[i+1:]...)
				bytes, err := json.Marshal(tasks)
				if err != nil {
					log.Fatalln(err)
				}

				found = true
				err = os.WriteFile("./database/database.json", bytes, 644)
				if err != nil {
					log.Fatalln(err)
				}

				for i := 0; i < 100; i++ {
					fmt.Println("")
				}

				fmt.Println(fmt.Sprintf(">> Task [%d] has been deleted successfully!", strConv))
				break
			}
		}

		if !found {
			fmt.Println(messages.TaskNotFoundSearhTask)
		}
	},
}
