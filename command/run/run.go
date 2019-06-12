package run

import (
	"github.com/5046312/gom/command"
)

func init() {
	cmd := &command.Command{
		Name: "run",
	}
	command.List[cmd.Name] = cmd
}
