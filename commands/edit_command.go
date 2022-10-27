package commands

import (
	"Terminal_Task_Managment_V2/functions"
	"Terminal_Task_Managment_V2/messages"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
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
			var chooseEditign string
			fmt.Scan(&chooseEditign)

			if strings.EqualFold(chooseEditign, "1") {
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

				var found bool
				for _, v := range tasks {
					if v.Id == strConv {
						found = true
						fmt.Print(messages.EnterNewNameTask)
						var newNameTask string
						fmt.Scan(&newNameTask)

						v.Name = newNameTask
						fmt.Print("OK")
						break
					}
				}

				if !found {
					functions.ClearTerminal()
					fmt.Print(messages.TaskNotFoundSearhTask)
					return
				}
				break
			}

			if strings.EqualFold(chooseEditign, "2") {
				fmt.Print(messages.CommandEditSoon)
				break
			}

			if strings.EqualFold(chooseEditign, "3") {
				fmt.Print(messages.CommandEditSoon)
				break
			}

			if strings.EqualFold(chooseEditign, "4") {
				fmt.Print(messages.CommandEditSoon)
				break
			}

			if strings.EqualFold(chooseEditign, "5") {
				functions.ClearTerminal()

				fmt.Println(messages.EditMenuClosed)
				break
			}

			fmt.Println(messages.NotGoodResponse)
			time.Sleep(2 * time.Second)
		}
	},
}
