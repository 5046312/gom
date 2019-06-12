package help

import (
	"fmt"

	"github.com/5046312/gom/command"
)

func init() {
	cmd := &command.Command{}
	cmd.Name = "help"
	// Flag
	cmd.Flag.Usage = usage
	command.List[cmd.Name] = cmd
}

func usage() {
	fmt.Println("help usage")
}
