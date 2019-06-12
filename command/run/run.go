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
	cmd.Usage = "run main.go"
	cmd.Exec = exec
	// Flag
	cmd.Flag.StringVar(&mainFile, "main", "main.go", "Need Project Main File")
	command.List[cmd.Name] = cmd
}

// exec
func exec(args []string) {
	fmt.Println("help usage")
}
