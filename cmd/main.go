package main

import (
	"bufio"
	// "fmt"
	"os"

	"github.com/btwkevin/ksh/command"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	inputArr := []string{}
	inputArr = append(inputArr, input)
	command.HandleCmd(inputArr)
}
