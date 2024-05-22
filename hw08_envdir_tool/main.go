package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	if size := len(args); size < 2 {
		log.Fatalf("not enough arguments to execute, must be at least 2, received %d", size)
	}

	dir := args[0]
	cmd := args[1:]

	env, err := ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(RunCmd(cmd, env))
}
