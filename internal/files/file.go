package files

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type FileScanInfo struct {
	filePath   string
	lineNumber int
	text       string
}

func (f FileScanInfo) FilePath() string {
	return f.filePath
}

func (f FileScanInfo) LineNumber() int {
	return f.lineNumber
}

func (f FileScanInfo) Text() string {
	return f.text
}

type ScanLine func(info FileScanInfo) error

type File struct {
	path string
}

func NewFile(path string) *File {
	return &File{path: path}
}

func (f *File) Path() string {
	return f.path
}

func (f *File) Scan(execLine ScanLine) error {
	file, err := os.Open(f.path)
	if err != nil {
		return fmt.Errorf("could not open file: %w", err)
	}
	defer func() {
		if errClose := file.Close(); errClose != nil {
			panic(errClose)
		}
	}()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	lineNumber := 0

	var errFinal error

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()

		info := FileScanInfo{
			filePath:   f.path,
			lineNumber: lineNumber,
			text:       line,
		}

		errFinal = errors.Join(errFinal, execLine(info))
	}

	if errScan := scanner.Err(); errScan != nil {
		return fmt.Errorf("could not scan file: %w", errScan)
	}

	return errFinal
}
