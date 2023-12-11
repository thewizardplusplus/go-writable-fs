package fsutils

import (
	writablefs "github.com/thewizardplusplus/go-writable-fs"
)

func removePaths(wfs writablefs.WritableFS, paths []string) error {
	var firstErrorHolder firstErrorHolder
	for _, path := range paths {
		if err := wfs.Remove(path); err != nil {
			firstErrorHolder.updateErr(err)
		}
	}

	return firstErrorHolder.firstErr()
}
