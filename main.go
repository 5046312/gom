package main

import (
	"fmt"

	"github.com/5046312/gom/command"
	_ "github.com/5046312/gom/command/build"
	_ "github.com/5046312/gom/command/run"
	"github.com/5046312/gom/helper"
)

type Gom struct {
	cmd map[string]*command.Command
}

var G *Gom

func init() {
	G = &Gom{}
	G.cmd = command.List
}

func main() {
	fmt.Println(helper.Icon)
	fmt.Println(len(G.cmd))
}
