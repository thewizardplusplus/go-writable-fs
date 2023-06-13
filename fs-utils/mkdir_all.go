package fsutils

import (
	"errors"
	"io/fs"
	pathpkg "path"
	"strings"

	writablefs "github.com/thewizardplusplus/go-writable-fs"
)

//go:generate mockery --name=MkdirAllFS --case=underscore --with-expecter

type MkdirAllFS interface {
	writablefs.WritableFS

	MkdirAll(path string, permissions fs.FileMode) error
}

func MkdirAll(
	wfs writablefs.WritableFS,
	path string,
	permissions fs.FileMode,
) error {
	if wfs, ok := wfs.(MkdirAllFS); ok {
		return wfs.MkdirAll(path, permissions)
	}

	if !fs.ValidPath(path) {
		// use the "mkdir" operation, since the `os.MkdirAll()` uses it
		return &fs.PathError{Op: "mkdir", Path: path, Err: fs.ErrInvalid}
	}

	var currentPath string
	for _, pathElement := range strings.Split(path, "/") {
		currentPath = pathpkg.Join(currentPath, pathElement)

		err := wfs.Mkdir(currentPath, permissions)
		if err != nil && !errors.Is(err, fs.ErrExist) {
			return err
		}
	}

	return nil
}
