package functions

import (
	"Terminal_Task_Managment_V2/messages"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func CreateDatabase() {
	_, err := os.Stat("./database/database.json")

	if os.IsNotExist(err) {
		fmt.Println(messages.DatabaseNotExist)
		time.Sleep(2 * time.Second)

		fmt.Println(messages.CreateDatabase)
		time.Sleep(2 * time.Second)
		barCreateDatabase := progressbar.Default(100, "Creation")
		for i := 0; i < 100; i++ {
			barCreateDatabase.Add(100)
		}

		fileCreate, err := os.Create("./database/database.json")
		if err != nil {
			log.Fatalln(err)
		}
		defer fileCreate.Close()

		_, err = io.Copy(fileCreate, strings.NewReader("[]"))
		if err != nil {
			log.Fatalln(err)
		}

		time.Sleep(1 * time.Second)
		fmt.Println(messages.DatabaseCreated)
		return

	}
}
