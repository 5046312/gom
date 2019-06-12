package main

import (
	"flag"

	"github.com/5046312/gom/command"
	_ "github.com/5046312/gom/command/build"
	_ "github.com/5046312/gom/command/help"
	_ "github.com/5046312/gom/command/run"
	"github.com/5046312/gom/helper"
)

type Gom struct {
	cmd  map[string]*command.Command
	flag flag.FlagSet
}

var G *Gom
var cmd string

func init() {
	G = &Gom{}
	G.cmd = command.List
}

func main() {
	helper.Icon()
	flag.Usage = func() { helper.Usage() }
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		helper.Usage()
		return
	}

	cmdName := args[0]
	cmd, exist := command.List[cmdName]
	if exist {
		cmd.Exec(args[1:])
	} else {
		// Command Not Exist
		helper.Usage()
	}
}
