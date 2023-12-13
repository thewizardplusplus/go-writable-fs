package fsutils

import (
	"io/fs"
	"sort"

	writablefs "github.com/thewizardplusplus/go-writable-fs"
)

//go:generate mockery --name=RemoveAllFS --case=underscore --with-expecter

type RemoveAllFS interface {
	writablefs.WritableFS

	RemoveAll(path string) error
}

func RemoveAll(wfs writablefs.WritableFS, path string) error {
	if wfs, ok := wfs.(RemoveAllFS); ok {
		return wfs.RemoveAll(path)
	}

	if !fs.ValidPath(path) {
		// use the "RemoveAll" operation, since the `os.RemoveAll()` uses it
		return &fs.PathError{Op: "RemoveAll", Path: path, Err: fs.ErrInvalid}
	}

	var firstErrorHolder firstErrorHolder
	var filePaths, dirPaths []string
	fs.WalkDir(wfs, path, func( //nolint: errcheck
		path string,
		dirEntry fs.DirEntry,
		err error,
	) error {
		if err != nil {
			firstErrorHolder.updateErr(err)
			return nil
		}

		if dirEntry.IsDir() {
			dirPaths = append(dirPaths, path)
		} else {
			filePaths = append(filePaths, path)
		}

		return nil
	})

	if err := removePaths(wfs, filePaths); err != nil {
		firstErrorHolder.updateErr(err)
	}

	// sort directories from more deeper to less deeper
	sort.Slice(dirPaths, func(i int, j int) bool {
		return countPathElements(dirPaths[i]) >
			countPathElements(dirPaths[j]) // reverse order
	})

	if err := removePaths(wfs, dirPaths); err != nil {
		firstErrorHolder.updateErr(err)
	}

	return firstErrorHolder.firstErr()
}
