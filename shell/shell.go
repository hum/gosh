package shell

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/hum/gosh/shell/env"
	"github.com/hum/gosh/shell/exec"
	"github.com/hum/gosh/shell/lexer"
	"github.com/hum/gosh/shell/process"
	"github.com/hum/gosh/shell/runtime"
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
		currPath := strings.Replace(runtime.CurrentPath, runtime.RootPath, "~", 1)
		prefix := fmt.Sprintf("%s :: %s > ", runtime.Username, currPath)
		w.WriteString(prefix)
		err := w.Flush()
		if err != nil {
			return err
		}

		s, err := r.ReadString(NEW_LINE_DELIMITER)
		if err != nil {
			return err
		}

		// Remove newline
		s = strings.TrimSuffix(s, "\n")

		// Split input line into the executable command and its arguments
		cmd, argv, err := lexer.Process(s)
		if err != nil {
			return fmt.Errorf("unable to parse the line, got error: %s", err)
		}

		// Happens when enter is pressed
		if cmd == "" {
			continue
		}

		if f, ok := builtinCmd[cmd]; ok {
			// If the command invocation is a builtin command
			if err = f(w, argv); err != nil {
				w.WriteString(fmt.Sprintf("gosh: %s\n", err))
				w.Flush()
				continue
			}
		} else {
			_, err := process.HandleExecutable(cmd, argv)
			if err != nil {
				w.WriteString(fmt.Sprintf("gosh: %s\n", err))
				w.Flush()
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
	)
	go handleSignal(sig)
}

func handleSignal(ch <-chan os.Signal) {
	for {
		select {
		case <-ch:
			if len(process.RunningProcesses) > 0 {
				process.KillChildren()
				continue
			}
			stop()
		default:
			// Ad-hoc sleep limit
			// maybe remove later
			time.Sleep(time.Millisecond * 200)
			continue
		}
	}
}

func stop() {
	w.WriteString("stopping the gosh shell")
	w.Flush()
	running = false
}
