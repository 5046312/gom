package build

import (
	"fmt"
	"os"
	"runtime"

	"github.com/5046312/gom/command"
	"github.com/5046312/gom/util"
)

var cmd *command.Command

var (
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
	mainfile := "main.go"
	if len(args) != 0 {
		mainfile = args[0]
	}
	switch platform {
	case "win", "windows":
		os.Setenv("GOOS", "windows")
	case "mac":
		os.Setenv("GOOS", "darwin")
	case "linux":
		os.Setenv("GOOS", "linux")
	}
	os.Setenv("GOARCH", runtime.GOARCH)
	util.Command("build", mainfile)
}
