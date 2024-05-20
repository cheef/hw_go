package main

import (
	"errors"
	"io"
	"math"
	"os"

	pb "github.com/cheggaaa/pb/v3"
)

var (
	ErrFileNotFound          = errors.New("file not found")
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrSamePath              = errors.New("trying to copy to the same path")
)

const (
	bufferSize = 1
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	from, err := validateFileToCopy(fromPath, toPath, offset)
	if err != nil {
		return err
	}
	defer from.Close()

	to, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer to.Close()

	return copyByBuffer(from, to, bufferSize, offset, limit)
}

func validateFileToCopy(fromPath, toPath string, offset int64) (f *os.File, e error) {
	from, err := os.Open(fromPath)
	if err != nil {
		return nil, ErrFileNotFound
	}

	info, err := from.Stat()

	if err != nil || info.Size() == 0 {
		return nil, ErrUnsupportedFile
	}

	if info.Size() <= offset {
		return nil, ErrOffsetExceedsFileSize
	}

	info2, err := os.Lstat(toPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return from, nil
		}

		return nil, err
	}

	if os.SameFile(info, info2) {
		return nil, ErrSamePath
	}

	return from, nil
}

func copyByBuffer(from *os.File, to *os.File, bufferSize, startOffset, limit int64) error {
	buffer := make([]byte, bufferSize)
	position := startOffset
	lastPosition := limit + startOffset

	bar, err := getProgressBar(from, startOffset, limit)
	if err != nil {
		return err
	}

	defer bar.Finish()

	for {
		if limit != 0 && position >= lastPosition {
			return nil
		}

		read, err := from.ReadAt(buffer, position)

		if err != nil && err != io.EOF {
			return err
		}

		if read == 0 {
			return nil
		}

		if _, err := to.Write(buffer); err != nil {
			return err
		}

		bar.Increment()
		position += bufferSize
	}
}

func getProgressBar(f *os.File, offset, limit int64) (*pb.ProgressBar, error) {
	info, err := f.Stat()
	if err != nil {
		return nil, err
	}

	barSize := info.Size()
	barSize -= offset

	if limit != 0 {
		barSize = int64(math.Min(float64(barSize), float64(limit)))
	}

	bar := pb.StartNew(int(barSize))

	return bar, nil
}
