package help

import (
	"fmt"

	"github.com/5046312/gom/command"
)

var cmd *command.Command

func init() {
	cmd = &command.Command{
		Name:  "help",
		Usage: Usage,
		Exec:  Exec,
	}
	command.List[cmd.Name] = cmd
}

func Usage() {
	fmt.Println(cmd.Name, "Usage: ")
}

func Exec(args []string) {
	if len(args) == 0 {
		Usage()
		return
	}
	command.List[args[0]].Usage()
}
