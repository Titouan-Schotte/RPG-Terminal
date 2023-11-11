package utils

import (
	"os"
	"os/exec"
)

func ClearConsole() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	cmd = exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
