package fsutils

import (
	"errors"
	"io/fs"
	pathpkg "path"

	writablefs "github.com/thewizardplusplus/go-writable-fs"
)

//go:generate mockery --name=CreateTempFS --case=underscore --with-expecter

type CreateTempFS interface {
	writablefs.WritableFS

	CreateTemp(baseDir string, pathPattern string) (writablefs.WritableFile, error)
}

func CreateTemp(
	wfs writablefs.WritableFS,
	baseDir string,
	pathPattern string,
) (writablefs.WritableFile, error) {
	if wfs, ok := wfs.(CreateTempFS); ok {
		return wfs.CreateTemp(baseDir, pathPattern)
	}

	fullPath := pathpkg.Join(baseDir, pathPattern)
	if !fs.ValidPath(fullPath) {
		innerErr := fs.ErrInvalid
		return nil, &fs.PathError{Op: "createtemp", Path: fullPath, Err: innerErr}
	}

	for try := 0; ; try++ {
		randomSuffix, err := generateRandomSuffix()
		if err != nil {
			return nil, &fs.PathError{Op: "createtemp", Path: fullPath, Err: err}
		}

		path := injectRandomSuffix(pathPattern, randomSuffix)
		fullPath := pathpkg.Join(baseDir, path)
		file, err := wfs.CreateExcl(fullPath)
		if err != nil {
			if errors.Is(err, fs.ErrExist) && try < maxCountOfTempTries {
				continue
			}

			return nil, err
		}

		return file, nil
	}
}
