package run

import (
	"fmt"

	"github.com/5046312/gom/command"
)

var (
	mainFile string
)

func init() {
	cmd := &command.Command{}
	cmd.Name = "run"
	// Flag
	cmd.Flag.Usage = usage
	cmd.Flag.StringVar(&mainFile, "main", "main.go", "Need Project Main File")
	command.List[cmd.Name] = cmd
}

func usage() {
	fmt.Println("run usage")
}
