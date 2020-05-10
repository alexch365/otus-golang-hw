package main

import (
	"errors"
	"io"
	"os"
	"time"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath string, toPath string, offset, limit int64) error {
	source, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer source.Close()

	sourceStat, err := source.Stat()
	if err != nil {
		return err
	}

	if !sourceStat.Mode().IsRegular() {
		return ErrUnsupportedFile
	}

	sourceSize := sourceStat.Size()
	if offset > sourceSize {
		return ErrOffsetExceedsFileSize
	}

	if _, err := source.Seek(offset, 0); err != nil {
		return err
	}

	dest, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer dest.Close()

	if limit == 0 || limit > sourceSize {
		limit = sourceSize
	}
	if (sourceSize - offset) < limit {
		limit = sourceSize - offset
	}

	var refreshRate time.Duration = 10
	bar := pb.Start64(limit).SetRefreshRate(time.Millisecond * refreshRate)
	defer bar.Finish()

	sourceReader := bar.NewProxyReader(source)
	if _, err := io.CopyN(dest, sourceReader, limit); err != nil {
		return err
	}

	return nil
}
