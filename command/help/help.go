package help

import (
	"github.com/5046312/gom/util"

	"github.com/5046312/gom/command"
)

var (
	cmd     *command.Command
	cmdName string
)

func init() {
	cmd = &command.Command{
		Name:  "help",
		Usage: "help run",
		Exec:  exec,
	}
	// Flag
	cmd.Flag.StringVar(&cmdName, "cmd", "help", "Command Help Content")
	command.List[cmd.Name] = cmd
}

// exec
func exec(args []string) {
	if len(args) < 1 {
		util.Usage()
	}
}
