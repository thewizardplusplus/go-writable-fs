package writablefs

import (
	"io"
	"io/fs"
)

//go:generate mockery --name=WritableFile --case=underscore --with-expecter

type WritableFile interface {
	fs.File
	io.Writer
}

//go:generate mockery --name=WritableFS --case=underscore --with-expecter

type WritableFS interface {
	fs.FS

	Mkdir(path string, permissions fs.FileMode) error
	Create(path string) (WritableFile, error)
	CreateExcl(path string) (WritableFile, error)
	Rename(oldPath string, newPath string) error
	Remove(path string) error
}
