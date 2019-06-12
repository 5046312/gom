package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/5046312/gom/command"
	_ "github.com/5046312/gom/command/build"
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
	flag.Usage = helper.Usage
	flag.Parse()
	args := flag.Args()
	fmt.Println(args)

	if len(args) < 1 {
		flag.Usage()
		os.Exit(2)
		return
	}
}
