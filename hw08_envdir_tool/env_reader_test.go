package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("when provided non-existing directory", func(t *testing.T) {
		env, err := ReadDir("random")

		require.Nil(t, env)
		require.Equal(t, ErrDirectoryDoesNotExist, err)
	})

	t.Run("when provided directory with ENV files", func(t *testing.T) {
		env, err := ReadDir("testdata/env")

		expected := Environment{
			"BAR":   EnvValue{Value: "bar", NeedRemove: false},
			"EMPTY": EnvValue{Value: "", NeedRemove: false},
			"FOO":   EnvValue{Value: "   foo\nwith new line", NeedRemove: false},
			"HELLO": EnvValue{Value: "\"hello\"", NeedRemove: false},
			"UNSET": EnvValue{Value: "", NeedRemove: true},
		}

		require.NoError(t, err)
		require.Equal(t, expected, env)
	})
}
