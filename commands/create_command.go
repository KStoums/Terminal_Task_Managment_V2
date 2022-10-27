package commands

import (
	"Terminal_Task_Managment_V2/functions"
	"Terminal_Task_Managment_V2/messages"
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
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

		functions.ClearTerminal()

		functions.CreateDir()

		functions.CreateDatabase()
		time.Sleep(2 * time.Second)

		newTask := TaskStruct{}

		functions.ClearTerminal()

		fmt.Print(messages.NewTaskName)

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			taskName := scanner.Text()
			if strings.EqualFold(taskName, "cancel") {
				functions.ClearTerminal()
				fmt.Print(messages.CancelCreateTask)
				return
			}

			if taskName != "" {
				newTask.Name = taskName
				newTask.Status = "Todo"
				break
			}

			functions.ClearTerminal()
			fmt.Println(messages.ErrorTaskName)
		}

		for {
			fmt.Print(messages.NewTaskPriority)
			var taskPriority string
			fmt.Scan(&taskPriority)

			if strings.EqualFold(taskPriority, "low") {
				newTask.Priority = "Low"
				break
			}

			if strings.EqualFold(taskPriority, "high") {
				newTask.Priority = "High"
				break
			}

			if strings.EqualFold(taskPriority, "cancel") {
				functions.ClearTerminal()
				fmt.Print(messages.CancelCreateTask)
				return
			}

			functions.ClearTerminal()
			fmt.Println(messages.ErrorTaskPriority)
		}

		for {
			fmt.Print(messages.DefineDue)
			var dueTime string
			fmt.Scan(&dueTime)

			if strings.EqualFold(dueTime, "none") {
				break
			}

			functions.ClearTerminal()
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
			functions.ClearTerminal()
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

		functions.ClearTerminal()

		fmt.Println(messages.SendTaskToDb)

		data = append(data, newTask)
		marshal, err := json.Marshal(data)
		if err != nil {
			functions.ClearTerminal()
			log.Fatalln(err)
		}

		os.WriteFile("./database/database.json", marshal, 644)
		time.Sleep(1 * time.Second)

		barSendTask := progressbar.Default(100, "Saving in progress...")
		for i := 0; i < 100; i++ {
			barSendTask.Add(50)
		}

		time.Sleep(1 * time.Second)
		functions.ClearTerminal()
		fmt.Println(messages.TaskCreated)
	},
}
