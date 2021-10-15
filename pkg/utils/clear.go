package utils

import (
	"os"
	"os/exec"
)

func ClearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
