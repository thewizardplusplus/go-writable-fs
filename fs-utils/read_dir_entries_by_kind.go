package fsutils

import (
	"io/fs"
)

func ReadDirEntriesByKind(
	fsInstance fs.FS,
	path string,
	kind DirEntryKind,
) ([]fs.DirEntry, error) {
	dirEntries, err := fs.ReadDir(fsInstance, path)
	if err != nil {
		return nil, err
	}

	files := make([]fs.DirEntry, 0, len(dirEntries))
	for _, dirEntry := range dirEntries {
		if GetDirEntryKind(dirEntry) == kind {
			files = append(files, dirEntry)
		}
	}

	return files, nil
}
