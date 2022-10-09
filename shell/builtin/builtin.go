package builtin

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/hum/gosh/shell/format"
)

func Echo(w *bufio.Writer, in []string) error {
	w.WriteString(strings.Join(in, " "))
	return w.Flush()
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

func Env(w *bufio.Writer, in []string) error {
	for _, env := range syscall.Environ() {
		w.WriteString(fmt.Sprintf("%s\n", env))
	}
	return w.Flush()
}

func Exit(w *bufio.Writer, in []string) error {
	w.WriteString("Exiting gosh shell")
	err := w.Flush()
	if err != nil {
		return err
	}
	os.Exit(0)
	return nil
}
