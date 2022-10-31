package edit_tasks

import (
	"Terminal_Task_Managment_V2/functions"
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
	dirDatabase := functions.CheckDirDatabase()
	if dirDatabase == false {
		return
	}

	checkDatabase := functions.CheckDatabase()
	if checkDatabase == false {
		return
	}

	fmt.Print(messages.DefineIDTask)
	var id string
	fmt.Scan(&id)

	if strings.EqualFold(id, "cancel") {
		functions.ClearTerminal()
		fmt.Print(messages.EditMenuClosed)
		return
	}

	strConv, err := strconv.Atoi(id)
	if err != nil {
		functions.ClearTerminal()
		fmt.Println(messages.ErrorNotIntTask)
		return
	}

	tasks := []*TaskStruct{}

	readFile, err := os.ReadFile("./database/database.json")
	if err != nil {
		functions.ClearTerminal()
		log.Fatalln(err)
	}
	err = json.Unmarshal(readFile, &tasks)
	if err != nil {
		functions.ClearTerminal()
		log.Fatalln(err)
	}

	var found bool
	for _, v := range tasks {
		if v.Id == strConv {
			found = true

			functions.ClearTerminal()
			fmt.Print(messages.ConfirmDoneTask)
			var confirm string
			fmt.Scan(&confirm)

			if strings.EqualFold(confirm, "y") {
				v.Status = "Done"

				marshal, err := json.Marshal(&tasks)
				if err != nil {
					functions.ClearTerminal()
					fmt.Print(err)
					return
				}

				err = os.WriteFile("./database/database.json", marshal, 644)
				if err != nil {
					functions.ClearTerminal()
					fmt.Print(err)
					return
				}

				functions.ClearTerminal()
				fmt.Println(messages.TaskEdited)
				break
			}

			if strings.EqualFold(confirm, "n") {
				functions.ClearTerminal()
				fmt.Println(messages.EditMenuClosed)
				break
			}

			functions.ClearTerminal()
			fmt.Println(messages.NotGoodResponse)
		}
	}

	if !found {
		functions.ClearTerminal()
		fmt.Print(messages.TaskNotFoundSearhTask)
		return
	}
}
