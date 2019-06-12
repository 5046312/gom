package build

import (
	"fmt"

	"github.com/5046312/gom/command"
)

func init() {
	cmd := &command.Command{}
	cmd.Name = "build"
	cmd.Exec = exec
	command.List[cmd.Name] = cmd
}

// exec
func exec(args []string) {
	fmt.Println("help usage")
}

func PlatWin() {

}

func PlatLinux() {

}

func PlatMac() {

}
