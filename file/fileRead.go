package file

import (
	"bufio"
	"os"
)

func NewFileRead(file string) *FileRead {
	return &FileRead{file: file}
}

type FileRead struct {
	file string

	f *os.File
}

func (f *FileRead) open() error {
	file, err := os.Open(f.file)
	if err != nil {
		return err
	}
	f.f = file
	return nil
}

func (f *FileRead) Read(callback func(line string)) error {
	if f.f == nil {
		if err := f.open(); err != nil {
			return err
		}
	}
	defer f.close()

	buf := bufio.NewScanner(f.f)
	for buf.Scan() {
		line := buf.Text()
		callback(line)
	}

	return nil
}

func (f *FileRead) close() {
	f.f.Close()
}
