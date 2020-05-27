package main

import (
	"log"
	"os"
)

func main() {
	minArgsLen := 3
	if len(os.Args) < minArgsLen {
		log.Fatalln("Two or more arguments needed. Usage: ./go-envdir /path/to/env/dir command arg1 arg2")
	}

	envDir := os.Args[1]
	cmdAndArgs := os.Args[2:]
	env, err := ReadDir(envDir)
	if err != nil {
		log.Fatalln(err)
	}

	os.Exit(RunCmd(cmdAndArgs, env))
}
