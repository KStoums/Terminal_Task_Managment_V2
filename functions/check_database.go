package functions

import (
	"Terminal_Task_Managment_V2/messages"
	"fmt"
	"os"
)

func CheckDatabase() bool {
	_, err := os.Stat("./database/database.json")
	if err != nil {
		if os.IsNotExist(err) {
			ClearTerminal()
			fmt.Println(messages.NoFileOrError)
			return false
		}

		ClearTerminal()
		fmt.Println(err)
		return false
	}

	return true
}
