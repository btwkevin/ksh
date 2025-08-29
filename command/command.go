package command

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func HandleCmd(input string) error {
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			fmt.Fprint(os.Stderr, "cd: missing argunment")
			return nil
		}
		return HandleChdir(args[1])
	case "exit":
		HandleExit()
	}
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func HandleChdir(dir string) error {
	return os.Chdir(dir)
}

func HandleExit() {
	os.Exit(0)
}
