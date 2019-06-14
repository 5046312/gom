package build

import (
	"fmt"

	"github.com/5046312/gom/command"
)

var cmd *command.Command

var (
	mainFile string = "main.go"
	platform string = "win"
)

func init() {
	cmd = &command.Command{
		Name:  "build",
		Usage: Usage,
		Exec:  Exec,
	}
	command.List[cmd.Name] = cmd
	cmd.Flag.StringVar(&platform, "p", "win", "")
}

func Usage() {
	fmt.Println(cmd.Name, `Usage:  gom build [main.go] [-p win|linux|mac]`)
}

func Exec(args []string) {
	fmt.Println(cmd.Name, ": ", args)
	if len(args) == 0 {
		// Default

	}
}

func PlatWin() {

}

func PlatLinux() {

}

func PlatMac() {

}
