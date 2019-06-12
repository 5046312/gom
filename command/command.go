package command

import "flag"

type Command struct {
	Name    string
	Usage   string
	Options map[string]string
	Flag    flag.FlagSet
}

var List = make(map[string]*Command)
