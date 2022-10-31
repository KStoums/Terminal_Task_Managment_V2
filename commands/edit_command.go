package commands

import (
	"Terminal_Task_Managment_V2/functions"
	"Terminal_Task_Managment_V2/functions/edit_tasks"
	"Terminal_Task_Managment_V2/messages"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
	"time"
)

func init() {
	rootCmd.AddCommand(editCommand)
}

var editCommand = &cobra.Command{
	Use:   "edit",
	Short: "Change the status or priority of a task",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println(messages.IncorrectSyntax)
			return
		}

		for {
			functions.ClearTerminal()
			fmt.Print(fmt.Sprintf("%s \n%s \n%s \n%s \n%s \n%s \n\n%s", messages.HeaderEditCommand, messages.ChooseOneEdit, messages.ChooseTwoEdit,
				messages.ChooseThreeEdit, messages.ChooseForEdit, messages.ChooseFiveEdit, messages.DefineEditChoose))
			var chooseEditing string
			fmt.Scan(&chooseEditing)

			if strings.EqualFold(chooseEditing, "1") {
				functions.ClearTerminal()
				edit_tasks.EditNameTask()
				break
			}

			if strings.EqualFold(chooseEditing, "2") {
				functions.ClearTerminal()
				edit_tasks.DoneTask()
				break
			}

			if strings.EqualFold(chooseEditing, "3") {
				functions.ClearTerminal()
				edit_tasks.ChangePriority()
				break
			}

			if strings.EqualFold(chooseEditing, "4") {
				fmt.Print(messages.CommandEditSoon)
				break
			}

			if strings.EqualFold(chooseEditing, "5") {
				functions.ClearTerminal()
				fmt.Print(messages.EditMenuClosed)
				break
			}

			fmt.Print(messages.NotGoodResponse)
			time.Sleep(2 * time.Second)
		}
	}}
