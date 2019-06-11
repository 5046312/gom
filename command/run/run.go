package run

import (
	"github.com/5046312/gom/command"
)

func init() {
	command := &command.Command{
		Name: "run",
	}
	append(command.Commands, command)
}
