package build

import (
	"fmt"

	"github.com/5046312/gom/command"
)

var cmd *command.Command

func init() {
	cmd = &command.Command{
		Name:  "build",
		Usage: Usage,
		Exec:  Exec,
	}
	command.List[cmd.Name] = cmd
}

func Usage() {
	fmt.Println(cmd.Name, "Usage: ")
}

func Exec(args []string) {
	fmt.Println(cmd.Name, ": ", args)
}

func PlatWin() {

}

func PlatLinux() {

}

func PlatMac() {

}
