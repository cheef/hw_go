package main

import (
	"errors"
	"os"
	"os/exec"
)

var (
	errorExitCode   = 1
	successExitCode = 0
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	for name, v := range env {
		if v.NeedRemove {
			err := os.Unsetenv(name)
			if err != nil {
				return errorExitCode
			}
			continue
		}

		err := os.Setenv(name, v.Value)
		if err != nil {
			return errorExitCode
		}
	}

	c := exec.Command(cmd[0], cmd[1:]...) //nolint:gosec
	c.Env = os.Environ()
	c.Stderr = os.Stderr
	c.Stdout = os.Stdout
	c.Stdin = os.Stdin

	if err := c.Run(); err != nil {
		var exitErr *exec.ExitError

		if errors.As(err, &exitErr) {
			return exitErr.ExitCode()
		}

		return errorExitCode
	}

	return successExitCode
}
