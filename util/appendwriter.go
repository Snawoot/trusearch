package util

import (
	"bufio"
	"os"
)

const (
	APPEND_BUF_SIZE = 32 * 1024
)

type AppendWriter struct {
	path string
}

func NewAppendWriter(path string) *AppendWriter {
	return &AppendWriter{path}
}

func (wr *AppendWriter) Write(p []byte) (int, error) {
	f, err := os.OpenFile(wr.path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.Write(p)
}

func NewBufferedAppendWriter(path string) *bufio.Writer {
	return bufio.NewWriterSize(NewAppendWriter(path), APPEND_BUF_SIZE)
}
