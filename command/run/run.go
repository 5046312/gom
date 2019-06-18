package run

import (
	"fmt"
	"os"

	"github.com/5046312/gom/command"
)

var runCmd *command.Command
var mainfile string = "main.go"

func init() {
	runCmd = &command.Command{
		Name:  "run",
		Info:  `Run Go Project, Hot Reload When File Changes`,
		Usage: Usage,
		Exec:  Exec,
	}
	command.List[runCmd.Name] = runCmd
}

func Usage() {
	fmt.Println(runCmd.Name, "Usage: ")
}

func Exec(args []string) {
	// fmt.Println(runCmd.Name, args)
	dir, _ := os.Getwd()
	watcher(dir, mainfile)
}
