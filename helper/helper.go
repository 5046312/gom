package helper

import (
	"fmt"

	"github.com/5046312/gom/command"
)

func Icon() {
	gom := `
 	 ██████╗   ██████╗  ███╗   ███╗
 	██╔════╝  ██╔═══██╗ ████╗ ████║
 	██║  ███╗ ██║   ██║ ██╔████╔██║
 	██║   ██║ ██║   ██║ ██║╚██╔╝██║
 	╚██████╔╝ ╚██████╔╝ ██║ ╚═╝ ██║
 	 ╚═════╝   ╚═════╝  ╚═╝     ╚═╝
	`
	fmt.Println(gom)
}

func Usage(args ...string) {
	if len(args) == 0 {
		fmt.Println("0 All Usage")
	} else if len(args) != 1 {
		fmt.Println("!=1 All Usage")
	} else {
		cmd, exist := command.List[args[0]]
		fmt.Println(args, cmd, exist)
		if exist {
			fmt.Println(cmd.Usage)
		} else {
			fmt.Println("Don't Know")
		}
	}
}
