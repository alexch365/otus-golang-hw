package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	// Place your code here
	t.Run("return error code for empty command", func(t *testing.T) {
		var emptyCommand []string

		exitCode := RunCmd(emptyCommand, Environment{})
		require.Equal(t, cmdErrorCode, exitCode)

		exitCode = RunCmd(nil, nil)
		require.Equal(t, cmdErrorCode, exitCode)
	})

	t.Run("return code from command", func(t *testing.T) {
		command := []string{"bash", "cmd"}
		commandCode := 127

		require.Equal(t, commandCode, RunCmd(command, Environment{}))
	})

	t.Run("set and unset env vars", func(t *testing.T) {
		command := []string{"bash"}
		env := Environment{"DEPLOY": "production", "UNSET": ""}

		exitCode := RunCmd(command, env)
		require.Equal(t, cmdSuccessCode, exitCode)

		deployEnv, ok := os.LookupEnv("DEPLOY")
		require.True(t, ok)
		require.Equal(t, "production", deployEnv)

		unsetEnv, ok := os.LookupEnv("UNSET")
		require.False(t, ok)
		require.Empty(t, unsetEnv)
	})
}
