package runner

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func run() bool {
	runnerLog("Running...")

	fmt.Println("dddddddddddddddddddddddd")
	fmt.Println(buildPath())
	if workDir() == "" {
		fmt.Println("workDir is empty")
	} else {
		fmt.Println(workDir())

	}
	appPath:=buildPath()
	wd:=workDir()
	os.Chdir(wd)
	cmd := exec.Command(appPath)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fatal(err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fatal(err)
	}

	err = cmd.Start()
	if err != nil {
		fatal(err)
	}

	go io.Copy(appLogWriter{}, stderr)
	go io.Copy(appLogWriter{}, stdout)

	go func() {
		<-stopChannel
		pid := cmd.Process.Pid
		runnerLog("Killing PID %d", pid)
		cmd.Process.Kill()
	}()

	return true
}
