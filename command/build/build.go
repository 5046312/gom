package build

import "github.com/5046312/gom/command"

func init() {
	cmd := &command.Command{
		Name: "build",
	}
	command.List[cmd.Name] = cmd
}

func PlatWin() {

}

func PlatLinux() {

}

func PlatMac() {

}
