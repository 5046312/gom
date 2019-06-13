package run

import (
	"fmt"

	"github.com/5046312/gom/command"
)

var cmd *command.Command

func init() {
	cmd = &command.Command{
		Name:  "run",
		Usage: Usage,
		Exec:  Exec,
	}
	command.List[cmd.Name] = cmd
}

func Usage() {
	fmt.Println("help usage")
}

func Exec(args []string) {
	fmt.Println("help usage")
}

func PlatWin() {

}

func PlatLinux() {

}

func PlatMac() {

}
