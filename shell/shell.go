package shell

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/hum/gosh/shell/env"
	"github.com/hum/gosh/shell/exec"
	"github.com/hum/gosh/shell/lexer"
	"github.com/hum/gosh/shell/process"
)

var (
	r *bufio.Reader = bufio.NewReader(os.Stdin)
	w *bufio.Writer = bufio.NewWriter(os.Stdout)
)

const (
	// Corresponds to the new line delimiter in the ASCII table
	// https://www.asciitable.com
	NEW_LINE_DELIMITER = 10

	// Displayed at the beginning of the line
	LINE_PREFIX = "> "
)

var running bool = true

func Execute() error {
	prepareSignalChan()

	// Load environment variables
	env.LoadEnv()
	err := exec.LoadExecutablesFromPath()
	if err != nil {
		return err
	}

	for running {
		fmt.Print(LINE_PREFIX)
		s, err := r.ReadString(NEW_LINE_DELIMITER)
		if err != nil {
			panic(err)
		}
		cmd, argv, err := lexer.Process(s)
		if err != nil {
			return fmt.Errorf("unable to parse the line, got error: %s", err)
		}
		if cmd == "" {
			continue
		}

		if f, ok := builtinCmd[cmd]; ok {
			// If the command invocation is a builtin command
			if err = f(w, argv); err != nil {
				fmt.Printf("gosh: %s\n", err)
				continue
			}
		} else {
			_, err := process.HandleExecutable(cmd, argv)
			if err != nil {
				fmt.Printf("gosh: %s\n", err)
				continue
			}
		}
	}
	return nil
}

func prepareSignalChan() {
	sig := make(chan os.Signal, 1)
	signal.Notify(
		sig,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGKILL,
	)

	go handleSignal(sig)
}

func handleSignal(ch <-chan os.Signal) {
	var s os.Signal = <-ch
	fmt.Printf("received %v signal\n", s)

	fmt.Println("stopping shell")
	stop()
}

func stop() {
	running = false

	fmt.Println("stopping child processes")
	process.KillChildren()
}
