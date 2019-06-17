package build

import (
	"fmt"
	"net/http"
	"os"

	"github.com/5046312/gom/command"
)

var cmd *command.Command

func init() {
	cmd = &command.Command{
		Name:  "file",
		Usage: Usage,
		Exec:  Exec,
	}
	command.List[cmd.Name] = cmd
}

func Usage() {
	fmt.Println(cmd.Name, `Usage:  gom file [port]`)
}

func Exec(args []string) {
	var address string = "127.0.0.1"
	var port string = "8080"
	if len(args) == 1 {
		port = args[0]
	} else if len(args) > 1 {
		port = args[0]
		address = args[1]
	}

	dir, _ := os.Getwd()
	fmt.Println("File Server:", address+":"+port, dir)
	http.ListenAndServe(address+":"+port, http.FileServer(http.Dir(dir)))
}
