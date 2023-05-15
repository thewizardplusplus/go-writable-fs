package fsutils

import (
	"io/fs"
)

type DirEntryKind int

const (
	NonDirKind DirEntryKind = iota
	DirKind
)

func GetDirEntryKind(dirEntry fs.DirEntry) DirEntryKind {
	if dirEntry.IsDir() {
		return DirKind
	}

	return NonDirKind
}
