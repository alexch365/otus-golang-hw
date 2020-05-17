package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	outFile, _ := ioutil.TempFile("", "*")
	defer os.Remove(outFile.Name())

	t.Run("limit exceeds file size", func(t *testing.T) {
		resultFile, _ := os.Open("testdata/out_offset100_limit10000.txt")
		resultStat, _ := resultFile.Stat()
		defer resultFile.Close()

		result := Copy("testdata/input.txt", outFile.Name(), 100, 10000)
		outStat, _ := outFile.Stat()

		require.Nil(t, result)
		require.Equal(t, resultStat.Size(), outStat.Size())
	})

	t.Run("offset exceeds file size", func(t *testing.T) {
		result := Copy("testdata/input.txt", outFile.Name(), 10000, 0)

		require.Equal(t, ErrOffsetExceedsFileSize, result)
	})

	t.Run("unsupported file", func(t *testing.T) {
		result := Copy("/dev/null", outFile.Name(), 0, 1000)
		require.Equal(t, ErrUnsupportedFile, result)

		result = Copy("/", outFile.Name(), 0, 1000)
		require.Equal(t, ErrUnsupportedFile, result)
	})
}
