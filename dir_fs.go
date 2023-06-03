package writablefs

import (
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type DirFS struct {
	innerDirFS fs.FS
	baseDir    string
}

func NewDirFS(baseDir string) DirFS {
	return DirFS{
		innerDirFS: os.DirFS(baseDir),
		baseDir:    baseDir,
	}
}

func (dfs DirFS) Open(path string) (fs.File, error) {
	return dfs.innerDirFS.Open(path)
}

func (dfs DirFS) Mkdir(path string, permissions fs.FileMode) error {
	if err := checkPath(path, "mkdir"); err != nil {
		return err
	}

	err := os.Mkdir(filepath.Join(dfs.baseDir, path), permissions)
	if err != nil {
		// restore the original path instead of its joined version
		err.(*fs.PathError).Path = path

		return err
	}

	return nil
}

// This check is made for consistency with the implementation of `os.DirFS()`.
//
// # License
//
//	BSD 3-Clause "New" or "Revised" License
//	Copyright (C) 2009 The Go Authors
func checkPath(path string, operation string) error {
	if !fs.ValidPath(path) ||
		(runtime.GOOS == "windows" && strings.ContainsAny(path, `\:`)) {
		return &fs.PathError{Op: operation, Path: path, Err: fs.ErrInvalid}
	}

	return nil
}
