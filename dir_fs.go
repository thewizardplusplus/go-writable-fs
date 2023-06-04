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

// Method `Stat()` is added for consistency with the implementation
// of `os.DirFS()`, which returns an instance that implements `fs.StatFS`.
func (dfs DirFS) Stat(path string) (fs.FileInfo, error) {
	return dfs.innerDirFS.(fs.StatFS).Stat(path)
}

func (dfs DirFS) Mkdir(path string, permissions fs.FileMode) error {
	if err := checkPath(path, "mkdir"); err != nil {
		return err
	}

	err := os.Mkdir(dfs.joinWithBaseDir(path), permissions)
	if err != nil {
		// restore the original path instead of its joined version
		updatePathInPathError(err, path)

		return err
	}

	return nil
}

func (dfs DirFS) Create(path string) (WritableFile, error) {
	// use the "open" operation, since the `os.Create()` uses it
	if err := checkPath(path, "open"); err != nil {
		return nil, err
	}

	file, err := os.Create(dfs.joinWithBaseDir(path))
	if err != nil {
		// restore the original path instead of its joined version
		updatePathInPathError(err, path)

		return nil, err
	}

	return file, nil
}

// Method `CreateExcl()` acts by analogy with function `os.Create()`,
// but replaces flag `os.O_TRUNC` with `os.O_EXCL`.
func (dfs DirFS) CreateExcl(path string) (WritableFile, error) {
	// use the "open" operation, since the `os.Create()` uses it
	if err := checkPath(path, "open"); err != nil {
		return nil, err
	}

	fullPath := dfs.joinWithBaseDir(path)
	file, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		// restore the original path instead of its joined version
		updatePathInPathError(err, path)

		return nil, err
	}

	return file, nil
}

func (dfs DirFS) Rename(oldPath string, newPath string) error {
	if err := checkPathsForRenaming(oldPath, newPath); err != nil {
		return err
	}

	err := os.Rename(dfs.joinWithBaseDir(oldPath), dfs.joinWithBaseDir(newPath))
	if err != nil {
		// restore the original paths instead of their joined versions
		updatePathsInLinkError(err, oldPath, newPath)

		return err
	}

	return nil
}

func (dfs DirFS) Remove(path string) error {
	if err := checkPath(path, "remove"); err != nil {
		return err
	}

	if err := os.Remove(dfs.joinWithBaseDir(path)); err != nil {
		// restore the original path instead of its joined version
		updatePathInPathError(err, path)

		return err
	}

	return nil
}

func (dfs DirFS) joinWithBaseDir(path string) string {
	return filepath.Join(dfs.baseDir, path)
}

// This check is made for consistency with the implementation of `os.DirFS()`.
//
// # License
//
//	BSD 3-Clause "New" or "Revised" License
//	Copyright (C) 2009 The Go Authors
func checkPath(path string, operation string) error {
	if !isValidPath(path) {
		return &fs.PathError{Op: operation, Path: path, Err: fs.ErrInvalid}
	}

	return nil
}

// This check is made for consistency with the implementation of `os.DirFS()`.
//
// # License
//
//	BSD 3-Clause "New" or "Revised" License
//	Copyright (C) 2009 The Go Authors
func checkPathsForRenaming(oldPath string, newPath string) error {
	if !isValidPath(oldPath) || !isValidPath(newPath) {
		innerErr := fs.ErrInvalid
		return &os.LinkError{Op: "rename", Old: oldPath, New: newPath, Err: innerErr}
	}

	return nil
}

// This check is made for consistency with the implementation of `os.DirFS()`.
//
// # License
//
//	BSD 3-Clause "New" or "Revised" License
//	Copyright (C) 2009 The Go Authors
func isValidPath(path string) bool {
	return fs.ValidPath(path) &&
		(runtime.GOOS != "windows" || !strings.ContainsAny(path, `\:`))
}

func updatePathInPathError(err error, path string) {
	err.(*fs.PathError).Path = path
}

func updatePathsInLinkError(err error, oldPath string, newPath string) {
	typedErr := err.(*os.LinkError)
	typedErr.Old = oldPath
	typedErr.New = newPath
}
