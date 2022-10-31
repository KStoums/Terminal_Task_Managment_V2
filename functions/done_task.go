package functions

import (
	"Terminal_Task_Managment_V2/messages"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
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

func DoneTask() {
	fmt.Print(messages.DefineIDTask)
	var id string
	fmt.Scan(&id)

	if strings.EqualFold(id, "cancel") {
		ClearTerminal()
		fmt.Print(messages.EditMenuClosed)
		return
	}

	strConv, err := strconv.Atoi(id)
	if err != nil {
		ClearTerminal()
		fmt.Println(messages.ErrorNotIntTask)
		return
	}

	tasks := []TaskStruct{}

	readFile, err := os.ReadFile("./database/database.json")
	if err != nil {
		log.Fatalln(err)
	}
	json.Unmarshal(readFile, &tasks)

	var found bool
	for _, v := range tasks {
		if v.Id == strConv {
			found = true

			ClearTerminal()
			fmt.Print(v)
			fmt.Print(messages.ConfirmDoneTask)
			var confirm string
			fmt.Scan(&confirm)

			if strings.EqualFold(confirm, "y") {
				v.Status = "Done"

				marshal, err := json.Marshal(&tasks)
				if err != nil {
					ClearTerminal()
					fmt.Print(err)
					return
				}

				err = os.WriteFile("./database/database.json", marshal, 644)
				if err != nil {
					ClearTerminal()
					fmt.Print(err)
					return
				}

				ClearTerminal()
				fmt.Println(messages.TaskEditedToDone)
				break
			}

			if strings.EqualFold(confirm, "n") {
				ClearTerminal()
				fmt.Println(messages.EditMenuClosed)
				break
			}

			ClearTerminal()
			fmt.Println(messages.NotGoodResponse)
		}
	}

	if !found {
		ClearTerminal()
		fmt.Print(messages.TaskNotFoundSearhTask)
		return
	}
}
