package command

import (
	"flag"
	"fmt"
)

type Command struct {
	Name    string
	Usage   string
	Options map[string]string
	Flag    flag.FlagSet
}

func (cmd *Command) Exec(args []string) {
	fmt.Println("Exec " + cmd.Name + " Command")
}

var List = make(map[string]*Command)
