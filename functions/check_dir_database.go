package functions

import (
	"Terminal_Task_Managment_V2/messages"
	"fmt"
	"os"
)

func CheckDirDatabase() bool {
	_, err := os.Stat("./database")
	if err != nil {
		if os.IsNotExist(err) {
			ClearTerminal()
			fmt.Println(messages.NoDirOrNoPermission)
			return false
		}

		ClearTerminal()
		fmt.Println(err)
		return false
	}

	return true
}
