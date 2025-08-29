package command

import (
	"fmt"
)

func HandleCmd(cmd []string) {
	for _, val := range cmd {
		fmt.Println(val)
	}
}
