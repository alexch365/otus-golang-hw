package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Environment map[string]string

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	resultEnv := Environment{}

	for _, fileInfo := range files {
		fileName := fileInfo.Name()
		skipCondition := strings.ContainsRune(fileName, '=') || fileInfo.IsDir()
		if skipCondition {
			continue
		}

		filePath := filepath.Join(dir, fileInfo.Name())
		envValue, err := envValueFromFile(filePath)
		if err != nil {
			return nil, err
		}

		resultEnv[fileName] = envValue
	}

	return resultEnv, nil
}

func envValueFromFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		return "", err
	}

	if fileStat.Size() == 0 {
		return "", nil
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return "", err
	}

	envValue := string(scanner.Bytes())
	envValue = strings.ReplaceAll(envValue, "\x00", "\n")
	envValue = strings.TrimRight(envValue, " \t")

	return envValue, nil
}
