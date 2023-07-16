package fsutils

import (
	"errors"
	"io/fs"
	pathpkg "path"

	writablefs "github.com/thewizardplusplus/go-writable-fs"
)

const (
	tempDirPermission   = fs.FileMode(0700)
	maxCountOfTempTries = 10000
)

//go:generate mockery --name=MkdirTempFS --case=underscore --with-expecter

type MkdirTempFS interface {
	writablefs.WritableFS

	MkdirTemp(baseDir string, pathPattern string) (string, error)
}

func MkdirTemp(
	wfs writablefs.WritableFS,
	baseDir string,
	pathPattern string,
) (string, error) {
	if wfs, ok := wfs.(MkdirTempFS); ok {
		return wfs.MkdirTemp(baseDir, pathPattern)
	}

	fullPath := pathpkg.Join(baseDir, pathPattern)
	if !fs.ValidPath(fullPath) {
		return "", &fs.PathError{Op: "mkdirtemp", Path: fullPath, Err: fs.ErrInvalid}
	}

	for try := 0; ; try++ {
		randomSuffix, err := generateRandomSuffix()
		if err != nil {
			return "", &fs.PathError{Op: "mkdirtemp", Path: fullPath, Err: err}
		}

		path := injectRandomSuffix(pathPattern, randomSuffix)
		fullPath := pathpkg.Join(baseDir, path)
		if err := wfs.Mkdir(fullPath, tempDirPermission); err != nil {
			if errors.Is(err, fs.ErrExist) && try < maxCountOfTempTries {
				continue
			}

			return "", err
		}

		return fullPath, nil
	}
}
