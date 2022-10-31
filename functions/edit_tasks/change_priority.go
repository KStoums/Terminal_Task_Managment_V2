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
)

func ChangePriority() {
	checkDirDatabase := functions.CheckDirDatabase()
	if checkDirDatabase == false {
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
		fmt.Println(messages.CancelEditTask)
		return
	}

	readFile, err := os.ReadFile("./database/database.json")
	if err != nil {
		functions.ClearTerminal()
		log.Fatalln(err)
	}

	tasks := []*TaskStruct{}

	err = json.Unmarshal(readFile, &tasks)
	if err != nil {
		functions.ClearTerminal()
		log.Fatalln(err)
	}

	strConv, err := strconv.Atoi(id)
	if err != nil {
		functions.ClearTerminal()
		fmt.Println(messages.ErrorNotIntTask)
		return
	}

	var found bool
	for _, v := range tasks {
		if v.Id == strConv {
			found = true

			fmt.Println(messages.DefineNewPriority)
			var newPriority string
			fmt.Scan(&newPriority)

			if strings.EqualFold(newPriority, "cancel") {
				functions.ClearTerminal()
				fmt.Println(messages.CancelEditTask)
				return
			}

			if strings.EqualFold(newPriority, "low") {
				v.Priority = "Low"

				bytes, err := json.Marshal(tasks)
				if err != nil {
					functions.ClearTerminal()
					log.Fatalln(err)
				}

				err = os.WriteFile("./database/database.json", bytes, 644)
				if err != nil {
					functions.ClearTerminal()
					log.Fatalln(err)
				}

				functions.ClearTerminal()
				fmt.Println(messages.TaskEdited)
				return
			}

			if strings.EqualFold(newPriority, "high") {
				v.Priority = "High"

				bytes, err := json.Marshal(tasks)
				if err != nil {
					functions.ClearTerminal()
					log.Fatalln(err)
				}

				err = os.WriteFile("./database/database.json", bytes, 644)
				if err != nil {
					functions.ClearTerminal()
					log.Fatalln(err)
				}

				functions.ClearTerminal()
				fmt.Println(messages.TaskEdited)
				return
			}

			functions.ClearTerminal()
			fmt.Println(messages.NotGoodResponse)
		}
	}

	if !found {
		functions.ClearTerminal()
		fmt.Println(messages.TaskNotFoundSearhTask)
		return
	}
}
