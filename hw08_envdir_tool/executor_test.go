package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("simple command with empty env", func(t *testing.T) {
		var cmd []string
		cmd = append(cmd, "pwd")
		code := RunCmd(cmd, make(Environment))

		require.Equal(t, 0, code)
	})
}
