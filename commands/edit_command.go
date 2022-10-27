package commands

import (
	"Terminal_Task_Managment_V2/functions"
	"Terminal_Task_Managment_V2/messages"
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
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
			var chooseEditing string
			fmt.Scan(&chooseEditing)

			if strings.EqualFold(chooseEditing, "1") {
				checkDirDatabase := functions.CheckDirDatabase()
				if checkDirDatabase == false {
					return
				}

				readFile, err := os.ReadFile("./database/database.json")
				if err != nil {
					functions.ClearTerminal()
					fmt.Println(messages.NoFileOrError)
					break
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
				json.Unmarshal(readFile, &tasks)

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
				break
			}

			if strings.EqualFold(chooseEditing, "2") {
				functions.ClearTerminal()

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
					break
				}

				strConv, err := strconv.Atoi(id)
				if err != nil {
					functions.ClearTerminal()
					log.Fatalln(err)
				}

				tasks := []TaskStruct{}

				var found bool
				for _, v := range tasks { //## !! A DEBUG !! ##//
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
							fmt.Println(messages.TaskEditedToDone)
							return
						}

						if strings.EqualFold(id, "n") {
							functions.ClearTerminal()
							fmt.Println(messages.EditMenuClosed)
							return
						}

						functions.ClearTerminal()
						fmt.Println(messages.IncorrectSyntax)
					}
				}

				if !found {
					functions.ClearTerminal()
					fmt.Print(messages.TaskNotFoundSearhTask)
					return
				}
			}

			if strings.EqualFold(chooseEditing, "3") {
				fmt.Print(messages.CommandEditSoon)
				break
			}

			if strings.EqualFold(chooseEditing, "4") {
				fmt.Print(messages.CommandEditSoon)
				break
			}

			if strings.EqualFold(chooseEditing, "5") {
				functions.ClearTerminal()

				fmt.Println(messages.EditMenuClosed)
				break
			}

			fmt.Println(messages.NotGoodResponse)
			time.Sleep(2 * time.Second)
		}
	},
}
