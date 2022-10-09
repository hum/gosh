package builtin

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hum/gosh/shell/format"
)

func Echo(w *bufio.Writer, in []string) error {
	fmt.Println(strings.Join(in, " "))
	return nil
}

func Ls(w *bufio.Writer, in []string) error {
	if len(in) == 0 {
		// If no argument passed, list home dir
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		in = append(in, homeDir)
	}

	for _, fp := range in {
		de, err := os.ReadDir(fp)
		if err != nil {
			return err
		}

		for _, e := range de {
			name := e.Name()
			if strings.HasPrefix(name, ".") {
				// Skip dotfiles by default
				continue
			}
			if e.IsDir() {
				name = format.SetColour(format.ColorBlue, name)
			}
			w.WriteString(name + "\n")
		}
		w.Flush()
	}
	return nil
}

func Exit(w *bufio.Writer, in []string) error {
	w.WriteString("Exiting gosh shell")
	w.Flush()
	os.Exit(0)
	return nil
}
