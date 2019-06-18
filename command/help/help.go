package help

import (
	"fmt"

	"github.com/5046312/gom/command"
)

var cmd *command.Command

func init() {
	cmd = &command.Command{
		Name:  "help",
		Info:  `Display Commands Help Information`,
		Usage: Usage,
		Exec:  Exec,
	}
	command.List[cmd.Name] = cmd
}

func Usage() {
	fmt.Println(`Gom Usage : gom [command] [...args]`)
	fmt.Println(`---------------------------------------`)
	for _, value := range command.List {
		info := fmt.Sprintf("%-5s : %s", value.Name, value.Info)
		fmt.Println(info)
	}
}

func Exec(args []string) {
	if len(args) == 0 {
		Usage()
		return
	}
	command.List[args[0]].Usage()
}
