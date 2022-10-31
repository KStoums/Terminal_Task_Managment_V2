package edit_tasks

import (
	"Terminal_Task_Managment_V2/functions"
	"Terminal_Task_Managment_V2/messages"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func EditNameTask() {
	checkDirDatabase := functions.CheckDirDatabase()
	if checkDirDatabase == false {
		return
	}

	readFile, err := os.ReadFile("./database/database.json")
	if err != nil {
		functions.ClearTerminal()
		fmt.Println(messages.NoFileOrError)
		return
	}

	fmt.Print(messages.DefineIDTask)
	var id string
	fmt.Scan(&id)

	if strings.EqualFold(id, "cancel") {
		functions.ClearTerminal()
		fmt.Print(messages.CancelEditTask)
		return
	}

	strConv, err := strconv.Atoi(id)
	if err != nil {
		functions.ClearTerminal()
		fmt.Print(messages.ErrorNotIntTask)
		return
	}

	tasks := []TaskStruct{}

	err = json.Unmarshal(readFile, &tasks)
	if err != nil {
		log.Fatalln(err)
	}

	var found bool
	for _, v := range tasks {
		if v.Id == strConv {
			found = true
			fmt.Print(messages.EnterNewNameTask)

			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				v.Name = scanner.Text() //## !! A DEBUG !! ##//
				return
			}

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
			fmt.Println(messages.TaskNameEdited)
			break
		}
	}

	if !found {
		functions.ClearTerminal()
		fmt.Print(messages.TaskNotFoundSearhTask)
		return
	}
	return
}
