package functions

import (
	"Terminal_Task_Managment_V2/messages"
	"fmt"
	"os"
	"os/exec"
)

func ClearTerminal() {
	cmd := exec.Command("cmd", "/c", "cls", "clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println(messages.ErrorWhenClearTerminal)
	}
}
