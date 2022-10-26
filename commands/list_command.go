package commands

import (
	"Terminal_Task_Managment_V2/messages"
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func init() {
	rootCmd.AddCommand(listCommand)
}

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "See all your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println(messages.IncorrectSyntax)
			return
		}

		_, err := os.Stat("./database")
		if err != nil {
			fmt.Println(messages.NoDirOrError)
			return
		}

		readFile, err := os.ReadFile("./database/database.json")
		if err != nil {
			log.Fatalln(err)
		}

		tasks := []TaskStruct{}

		err = json.Unmarshal(readFile, &tasks)
		if err != nil {
			log.Fatalln(err)
		}

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
	},
}
