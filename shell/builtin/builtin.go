package builtin

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/hum/gosh/shell/format"
	"github.com/hum/gosh/shell/runtime"
)

func Cd(w *bufio.Writer, in []string) error {
	if len(in) == 0 {
		// Go to root
		w.WriteString(fmt.Sprintf("changing path to: %s\n", runtime.RootPath))
		runtime.CurrentPath = runtime.RootPath
		return nil
	}

	if len(in) > 1 {
		return errors.New("too many arguments")
	}

	// Build the path
	var arg string = runtime.RootPath + "/" + in[0]

	// Assert the location exists
	fi, err := os.Stat(arg)
	if err != nil {
		return errors.New(fmt.Sprintf("could not open location: %s", arg))
	}

	// Assert the location is a directory
	if !fi.IsDir() {
		return errors.New(fmt.Sprintf("location: %s is not a directory", arg))
	}

	w.WriteString(fmt.Sprintf("changing path to: %s\n", arg))
	w.Flush()

	// TODO: Handle case for parent directory
	// currently this goes to $HOME parent directory
	runtime.CurrentPath = arg
	return nil
}

func Echo(w *bufio.Writer, in []string) error {
	w.WriteString(strings.Join(in, " ") + "\n")
	return w.Flush()
}

func Ls(w *bufio.Writer, in []string) error {
	if len(in) == 0 {
		// If no argument passed, list home dir
		in = append(in, runtime.CurrentPath)
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
