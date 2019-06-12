package main

import (
	"flag"

	"github.com/5046312/gom/command"
	_ "github.com/5046312/gom/command/build"
	_ "github.com/5046312/gom/command/help"
	_ "github.com/5046312/gom/command/run"
	"github.com/5046312/gom/util"
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
	util.Icon()
	flag.Usage = func() { util.Usage() }
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		util.Usage()
		return
	}

	cmd, exist := command.List[args[0]]
	if exist {
		cmd.Exec(args[1:])
	} else {
		// Command Not Exist
		util.Usage()
	}
}
