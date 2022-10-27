package functions

import (
	"os"
	"os/exec"
)

func ClearTerminal() {
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}
