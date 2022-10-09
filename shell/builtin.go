package shell

import (
	"bufio"

	"github.com/hum/gosh/shell/builtin"
)

const (
	Echo = "echo"
	Exit = "exit"
	Ls   = "ls"
)

var builtinCmd = map[string]func(w *bufio.Writer, args []string) error{
	Echo: builtin.Echo,
	Ls:   builtin.Ls,
	Exit: builtin.Exit,
}
