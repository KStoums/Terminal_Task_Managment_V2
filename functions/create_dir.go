package functions

import (
	"Terminal_Task_Managment_V2/messages"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"os"
	"time"
)

func CreateDir() {
	_, err2 := os.Stat("./database")
	if os.IsNotExist(err2) {
		os.Mkdir("./database", 0755)
		fmt.Println(messages.NoDirExist)
		time.Sleep(2 * time.Second)

		fmt.Println(messages.CreateDir)
		time.Sleep(1 * time.Second)
		barCreateDir := progressbar.Default(100, "Creating the \"./database/\" file")
		for i := 0; i < 100; i++ {
			barCreateDir.Add(50)
		}
		time.Sleep(1 * time.Second)
		fmt.Println(messages.DirCreated)
		time.Sleep(2 * time.Second)

		for i := 0; i < 100; i++ {
			fmt.Println("")
		}
	}
}
