package writablefs

import (
	"io/fs"
)

type DummyFS struct{}

func (DummyFS) Open(name string) (fs.File, error) {
	return nil, MakeSpecificErrNotImplemented()
}

func (DummyFS) Mkdir(path string, permissions fs.FileMode) error {
	return MakeSpecificErrNotImplemented()
}

func (DummyFS) Create(path string) (WritableFile, error) {
	return nil, MakeSpecificErrNotImplemented()
}

func (DummyFS) CreateExcl(path string) (WritableFile, error) {
	return nil, MakeSpecificErrNotImplemented()
}

func (DummyFS) Rename(oldPath string, newPath string) error {
	return MakeSpecificErrNotImplemented()
}

func (DummyFS) Remove(path string) error {
	return MakeSpecificErrNotImplemented()
}
