package command

type Command struct {
	Name    string
	Usage   string
	Options map[string]string
}

var Commands []*Command
