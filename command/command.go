package command

import (
	"flag"
)

type Command struct {
	Name  string
	Flag  flag.FlagSet
	Usage func()
	Exec  func(args []string)
}

var List = make(map[string]*Command)
