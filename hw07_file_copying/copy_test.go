package main

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	const (
		destination = "/tmp/input.txt"
	)

	t.Run("when file does not exist", func(t *testing.T) {
		err := Copy("./test.txt", destination, 0, 0)

		require.Equal(t, ErrFileNotFound, err)
	})

	t.Run("when valid file provided", func(t *testing.T) {
		err := Copy("testdata/input.txt", destination, 0, 0)

		require.Nil(t, err)
	})

	t.Run("when same file provided", func(t *testing.T) {
		err := Copy("testdata/input.txt", "testdata/input.txt", 0, 0)

		require.Equal(t, ErrSamePath, err)
	})

	t.Run("when provided file without a size", func(t *testing.T) {
		err := Copy("/dev/urandom", destination, 0, 0)

		require.Equal(t, ErrUnsupportedFile, err)
	})

	t.Run("when provided file less then specified offset", func(t *testing.T) {
		err := Copy("testdata/out_offset0_limit10.txt", destination, 100, 0)

		require.Equal(t, ErrOffsetExceedsFileSize, err)
	})

	t.Run("when limit defined", func(t *testing.T) {
		matches := "testdata/out_offset0_limit10.txt"
		err := Copy("testdata/input.txt", destination, 0, 10)

		require.Nil(t, err)

		f1, err := os.Open(matches)
		require.Nil(t, err)
		b1, _ := io.ReadAll(f1)

		f2, err := os.Open(destination)
		require.Nil(t, err)
		b2, _ := io.ReadAll(f2)

		require.Equal(t, b1, b2)
	})

	t.Run("when offset and limit defined", func(t *testing.T) {
		matches := "testdata/out_offset100_limit1000.txt"
		err := Copy("testdata/input.txt", destination, 100, 1000)

		require.Nil(t, err)

		f1, err := os.Open(matches)
		require.Nil(t, err)
		b1, _ := io.ReadAll(f1)

		f2, err := os.Open(destination)
		require.Nil(t, err)
		b2, _ := io.ReadAll(f2)

		require.Equal(t, b1, b2)
	})
}
