package lexer

import (
	"fmt"
	"strings"
)

// TODO: actual parsing logic

func Process(in string) (cmd string, argv []string, err error) {
	// TODO: handle more edge-cases
	in = strings.ReplaceAll(in, "\n", "")

	// TODO: better whitespace and argument handling
	argv = strings.Split(in, " ")
	if len(argv) == 0 {
		err = fmt.Errorf("could not parse input data")
		return
	}
	cmd = argv[0]
	argv = argv[1:]
	return
}
