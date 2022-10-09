package shell

import (
	"bufio"

	"github.com/hum/gosh/shell/builtin"
)

const (
	Echo = "echo"
	Env  = "env"
	Exit = "exit"
	Ls   = "ls"
)

var builtinCmd = map[string]func(w *bufio.Writer, args []string) error{
	Echo: builtin.Echo,
	Env:  builtin.Env,
	Exit: builtin.Exit,
	Ls:   builtin.Ls,
}
