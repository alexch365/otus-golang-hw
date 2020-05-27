package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	testDir := "testdata/env"
	testFiles, _ := ioutil.ReadDir(testDir)

	t.Run("skip file with '=' in name", func(t *testing.T) {
		tempfile, _ := ioutil.TempFile(testDir, "S=KIP")
		defer os.Remove(tempfile.Name())

		envMap, err := ReadDir(testDir)
		require.Nil(t, err)
		require.Equal(t, len(testFiles), len(envMap))
	})

	t.Run("remove whitespaces and tabs from env var", func(t *testing.T) {
		tempfile, _ := ioutil.TempFile(testDir, "tabbed lines")
		defer os.Remove(tempfile.Name())

		ioutil.WriteFile(tempfile.Name(), []byte("env value\t is\x00 \t \nskipped"), os.ModePerm)
		envMap, err := ReadDir(testDir)
		require.Nil(t, err)
		require.Equal(t, len(testFiles)+1, len(envMap))
		require.Equal(t, "env value\t is\n", envMap[filepath.Base(tempfile.Name())])
	})

	t.Run("convert empty file to empty string env value", func(t *testing.T) {
		tempfile, _ := ioutil.TempFile(testDir, "empty")
		defer os.Remove(tempfile.Name())

		envMap, err := ReadDir(testDir)
		require.Nil(t, err)
		require.Equal(t, len(envMap), len(testFiles)+1)
		require.Equal(t, "", envMap[filepath.Base(tempfile.Name())])
	})
}
