package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/btwkevin/ksh/command"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if err = command.HandleCmd(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
