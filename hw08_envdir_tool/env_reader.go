package main

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

var (
	ErrDirectoryDoesNotExist       = errors.New("specified directory doesn't exist")
	ErrEnvContainsDirectory        = errors.New("contains directory, but only files are supported")
	ErrEnvNameContainsInvalidChars = errors.New("ENV name contains invalid chars")
)

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, ErrDirectoryDoesNotExist
	}

	env := make(Environment)

	for _, entry := range entries {
		if entry.IsDir() {
			return nil, ErrEnvContainsDirectory
		}

		name := entry.Name()

		if strings.Contains(name, "=") {
			return nil, ErrEnvNameContainsInvalidChars
		}

		path := filepath.Join(dir, name)
		info, err := entry.Info()
		if err != nil {
			return nil, err
		}

		toRemove := info.Size() == 0
		value, err := readFileContent(path)
		if err != nil {
			return nil, err
		}

		env[name] = EnvValue{Value: value, NeedRemove: toRemove}
	}

	return env, nil
}

func readFileContent(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return "", err
	}

	value := scanner.Text()
	value = strings.ReplaceAll(value, "\x00", "\n")
	value = strings.TrimRight(value, " \t")

	return value, nil
}
