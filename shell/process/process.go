package process

import (
	"fmt"
	"os"
	"syscall"

	"github.com/hum/gosh/shell/exec"
)

type RunningProcess struct {
	Pid  int
	Path string
}

var RunningProcesses []RunningProcess

func HandleExecutable(cmd string, args []string) (pid int, err error) {
	var bin string = cmd

	v, ok := exec.Executables[bin]
	if ok {
		bin = v
	} else {
		fi, err := os.Stat(cmd)
		if err != nil {
			return 0, fmt.Errorf("error stat syscall: %s", err)
		}

		if fi.IsDir() {
			return 0, fmt.Errorf("cannot execute a directory: %s", cmd)
		}

		if !fileIsExecutable(fi.Mode()) {
			return 0, fmt.Errorf("not an executable: %s", cmd)
		}
	}

	pid, err = forkProcess(bin, args)
	if err != nil {
		return 0, fmt.Errorf("could not fork child process, got error: %s", err)
	}

	rp := RunningProcess{
		Pid:  pid,
		Path: cmd,
	}
	RunningProcesses = append(RunningProcesses, rp)
	return
}

func KillChildren() error {
	for _, c := range RunningProcesses {
		err := syscall.Kill(c.Pid, syscall.SIGINT)
		if err != nil {
			// TODO: handle
			fmt.Printf("could not kill process pid(%d), got error: %s\n", c.Pid, err)
			continue
		}
	}
	// Clear the child process stack
	RunningProcesses = []RunningProcess{}
	return nil
}

func forkProcess(proc string, args []string) (pid int, err error) {
	argv := make([]string, 0, len(args)+1)
	argv = append(argv, proc)
	argv = append(argv, args...)

	return syscall.ForkExec(
		argv[0],
		argv,
		&syscall.ProcAttr{
			Env: syscall.Environ(),
			Sys: &syscall.SysProcAttr{
				Setsid: true,
			},
			Files: []uintptr{0, 1, 2},
		})
}

func fileIsExecutable(fm os.FileMode) bool {
	return fm&0100 != 0
}
