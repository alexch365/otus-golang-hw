package main

import (
	"os"
	"os/exec"
)

const (
	cmdErrorCode   = 1
	cmdSuccessCode = 0
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	if len(cmd) == 0 {
		return cmdErrorCode
	}

	command := exec.Command(cmd[0], cmd[1:]...) //nolint:gosec

	for envName, envValue := range env {
		var err error
		if envValue == "" {
			err = os.Unsetenv(envName)
		} else {
			err = os.Setenv(envName, envValue)
		}

		if err != nil {
			return cmdErrorCode
		}
	}

	command.Env = os.Environ()
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin

	if err := command.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError.ExitCode()
		}
		return cmdErrorCode
	}

	return cmdSuccessCode
}
