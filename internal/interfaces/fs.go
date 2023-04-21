package interfaces

import (
	"io"
	"os"
)

type FS interface {
	OpenFile(name string, flag int, perm os.FileMode) (WriteCloser, error)
}

type WriteCloser interface {
	io.Writer
	io.Closer
}

type LocalFS struct{}

func NewLocalFS() *LocalFS {
	return &LocalFS{}
}

func (fs *LocalFS) OpenFile(name string, flag int, perm os.FileMode) (WriteCloser, error) {
	return os.OpenFile(name, flag, perm)
}
