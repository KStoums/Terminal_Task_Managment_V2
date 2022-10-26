package commands

import (
	"Terminal_Task_Managment_V2/functions"
	"Terminal_Task_Managment_V2/messages"
	"encoding/json"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
	"log"
	"os"
	"time"
)

type TaskStruct struct {
	Id          int
	Name        string
	Status      string
	Priority    string
	DateCreated time.Time
	Due         int
}

func init() {
	rootCmd.AddCommand(createCommand)
}

var createCommand = &cobra.Command{
	Use:   "create",
	Short: "Create a new task",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println(messages.IncorrectSyntax)
			return
		}

		for i := 0; i < 100; i++ {
			fmt.Println("")
		}

		functions.CreateDir()
		time.Sleep(2 * time.Second)

		functions.CreateDatabase()
		time.Sleep(2 * time.Second)

		newTask := TaskStruct{}

		for i := 0; i < 100; i++ {
			fmt.Println("")
		}

		for {
			fmt.Print(messages.NewTaskName)
			var taskName string
			fmt.Scanln(&taskName)

			if taskName != "" {
				newTask.Name = taskName
				newTask.Status = "Todo"
				break
			}

			if taskName == "cancel" || taskName == "Cancel" {
				for i := 0; i < 100; i++ {
					fmt.Println("")
				}
				fmt.Println(messages.CancelCreateTask)
				return
			}

			fmt.Println(messages.ErrorTaskName)
		}

		for {
			fmt.Print(messages.NewTaskPriority)
			var taskPriority string
			fmt.Scanln(&taskPriority)

			if taskPriority == "Low" || taskPriority == "low" {
				newTask.Priority = "Low"
				break
			}

			if taskPriority == "High" || taskPriority == "high" {
				newTask.Priority = "High"
				break
			}

			if taskPriority == "cancel" || taskPriority == "Cancel" {
				for i := 0; i < 100; i++ {
					fmt.Println("")
				}
				fmt.Println(messages.CancelCreateTask)
				return
			}

			fmt.Println(messages.ErrorTaskPriority)
		}

		for {
			fmt.Print(messages.DefineDue)
			var dueTime string
			fmt.Scanln(&dueTime)

			if dueTime == "none" || dueTime == "None" {
				break
			}

			fmt.Println(messages.DueSoonFeature)
		}

		newTask.DateCreated = time.Now()

		openFile, err := os.ReadFile("./database/database.json")
		if err != nil {
			log.Fatalln(err)
		}

		var data []TaskStruct
		err = json.Unmarshal(openFile, &data)
		if err != nil {
			log.Fatalln(err)
		}

		var id int
		for i, task := range data {
			if i != task.Id {
				break
			}
			id++
		}
		newTask.Id = id

		for i := 0; i < 100; i++ {
			fmt.Println("")
		}

		fmt.Println(messages.SendTaskToDb)

		data = append(data, newTask)
		marshal, err := json.Marshal(data)
		if err != nil {
			log.Fatalln(err)
		}

		os.WriteFile("./database/database.json", marshal, 644)
		time.Sleep(1 * time.Second)

		barSendTask := progressbar.Default(100, "Saving in progress...")
		for i := 0; i < 100; i++ {
			barSendTask.Add(50)
		}

		fmt.Println(messages.TaskCreated)
	},
}
