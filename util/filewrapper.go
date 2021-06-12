package util

import (
	"bufio"
	"io"
)

type FileWrapper struct {
	r io.ReadCloser
	br *bufio.Reader
}

func NewFileWrapper(r io.ReadCloser) *FileWrapper {
	return &FileWrapper{
		r: r,
		br: bufio.NewReader(r),
	}
}

func (fw *FileWrapper) Read(p []byte) (int, error) {
	return fw.br.Read(p)
}

func (fw *FileWrapper) Close() error {
	fw.br = nil
	return fw.r.Close()
}
