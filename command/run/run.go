package run

import (
	"fmt"

	"github.com/5046312/gom/command"
)

var cmd *command.Command

func init() {
	cmd = &command.Command{
		Name:  "run",
		Info:  `Run Go Project, Hot Reload When File Changes`,
		Usage: Usage,
		Exec:  Exec,
	}
	command.List[cmd.Name] = cmd
}

func Usage() {
	fmt.Println(cmd.Name, "Usage: ")
}

func Exec(args []string) {
	fmt.Println(cmd.Name, args)
}
