# go-writable-fs

[![GoDoc](https://godoc.org/github.com/thewizardplusplus/go-writable-fs?status.svg)](https://godoc.org/github.com/thewizardplusplus/go-writable-fs)
[![Go Report Card](https://goreportcard.com/badge/github.com/thewizardplusplus/go-writable-fs)](https://goreportcard.com/report/github.com/thewizardplusplus/go-writable-fs)

The library that implements a writable file system.

## Features

- the interfaces:
  - of a writable file:
    - based on interfaces `fs.File` and `io.Writer`;
  - of a writable file system:
    - based on interface `fs.FS`;
    - additional methods:
      - `Mkdir(path string, permissions fs.FileMode) error`;
      - `Create(path string) (WritableFile, error)`;
      - `CreateExcl(path string) (WritableFile, error)`;
      - `Rename(oldPath string, newPath string) error`;
      - `Remove(path string) error`;
- the implementations:
  - `writablefs.DirFS`:
    - also implements interface `fs.StatFS`;
    - uses function `os.DirFS()` to implement interfaces `fs.FS` and `fs.StatFS`;
    - operates with a real file system rooted at a specified path;
  - `writablefs.DummyFS`:
    - all the methods return error `writablefs.ErrNotImplemented`;
    - is useful for adding missed (but unnecessary) methods to implementations of a writable file system;
  - `mocks.WritableFS`:
    - uses package `github.com/stretchr/testify/mock` (see https://github.com/stretchr/testify);
    - is useful for mocking a writable file system;
- the helper functions:
  - function `fsutils.ReadDirEntriesByKind()` that reads a specified directory via function `fs.ReadDir` and selects entries with a specified kind (directory / not directory).

## Installation

```
$ go get github.com/thewizardplusplus/go-writable-fs@latest
```

## Examples

`writablefs.DirFS`:

```go
package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	writablefs "github.com/thewizardplusplus/go-writable-fs"
)

type FSEntity struct {
	Name            string
	ChildFSEntities []FSEntity
}

func (e FSEntity) IsDir() bool {
	return len(e.ChildFSEntities) != 0
}

func MaterializeFSEntity(
	wfs writablefs.WritableFS,
	baseDir string,
	fsEntity FSEntity,
) error {
	enrichedName := filepath.Join(baseDir, fsEntity.Name)
	if fsEntity.IsDir() {
		if err := wfs.Mkdir(enrichedName, 0700); err != nil {
			return fmt.Errorf("unable to create directory %q: %w", fsEntity.Name, err)
		}

		return nil
	}

	file, err := wfs.Create(enrichedName)
	if err != nil {
		return fmt.Errorf("unable to create file %q: %w", fsEntity.Name, err)
	}
	defer file.Close()

	return nil
}

func MaterializeFSEntities(
	wfs writablefs.WritableFS,
	baseDir string,
	fsEntities []FSEntity,
) error {
	for _, fsEntity := range fsEntities {
		if err := MaterializeFSEntity(wfs, baseDir, fsEntity); err != nil {
			return err
		}

		if err := MaterializeFSEntities(
			wfs,
			filepath.Join(baseDir, fsEntity.Name),
			fsEntity.ChildFSEntities,
		); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	tempDir, err := os.MkdirTemp("", "example-*")
	if err != nil {
		log.Fatalf("unable to create a temporary directory: %s", err)
	}
	defer os.RemoveAll(tempDir)

	dfs := writablefs.NewDirFS(tempDir)
	if err := MaterializeFSEntities(dfs, ".", []FSEntity{
		{Name: "directory-1", ChildFSEntities: []FSEntity{
			{Name: "file-1"},
			{Name: "file-2"},
		}},
		{Name: "file-3"},
	}); err != nil {
		log.Fatalf("unable to create the FS entities: %s", err)
	}

	if err := fs.WalkDir(dfs, ".", func(
		path string,
		_ fs.DirEntry,
		err error,
	) error {
		fmt.Println(path)
		return err
	}); err != nil {
		log.Fatalf("unable to walk the temporary directory: %s", err)
	}

	// Output:
	// .
	// directory-1
	// directory-1/file-1
	// directory-1/file-2
	// file-3
}
```

`writablefs.DummyFS`:

```go
package main

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"

	writablefs "github.com/thewizardplusplus/go-writable-fs"
)

type FSEntity struct {
	Name            string
	ChildFSEntities []FSEntity
}

func (e FSEntity) IsDir() bool {
	return len(e.ChildFSEntities) != 0
}

func MaterializeFSEntity(
	wfs writablefs.WritableFS,
	baseDir string,
	fsEntity FSEntity,
) error {
	enrichedName := filepath.Join(baseDir, fsEntity.Name)
	if fsEntity.IsDir() {
		if err := wfs.Mkdir(enrichedName, 0700); err != nil {
			return fmt.Errorf("unable to create directory %q: %w", fsEntity.Name, err)
		}

		return nil
	}

	file, err := wfs.Create(enrichedName)
	if err != nil {
		return fmt.Errorf("unable to create file %q: %w", fsEntity.Name, err)
	}
	defer file.Close()

	return nil
}

func MaterializeFSEntities(
	wfs writablefs.WritableFS,
	baseDir string,
	fsEntities []FSEntity,
) error {
	for _, fsEntity := range fsEntities {
		if err := MaterializeFSEntity(wfs, baseDir, fsEntity); err != nil {
			return err
		}

		if err := MaterializeFSEntities(
			wfs,
			filepath.Join(baseDir, fsEntity.Name),
			fsEntity.ChildFSEntities,
		); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	expectedErr := writablefs.ErrNotImplemented

	var dfs writablefs.DummyFS
	if err := MaterializeFSEntities(dfs, ".", []FSEntity{
		{Name: "directory-1", ChildFSEntities: []FSEntity{
			{Name: "file-1"},
			{Name: "file-2"},
		}},
		{Name: "file-3"},
	}); !errors.Is(err, expectedErr) {
		log.Fatalf("expect error %q", expectedErr)
	}

	fmt.Println("done")

	// Output:
	// done
}
```

`fsutils.ReadDirEntriesByKind()`:

```go
package main

import (
	"fmt"
	"log"
	"testing/fstest"

	fsutils "github.com/thewizardplusplus/go-writable-fs/fs-utils"
)

func main() {
	mapFS := fstest.MapFS{
		"directory-1/file-1.1":                 &fstest.MapFile{},
		"directory-1/file-1.2":                 &fstest.MapFile{},
		"directory-2/file-2.1":                 &fstest.MapFile{},
		"directory-2/file-2.2":                 &fstest.MapFile{},
		"directory-2/directory-2.1/file-2.1.1": &fstest.MapFile{},
		"directory-2/directory-2.1/file-2.1.2": &fstest.MapFile{},
		"directory-2/directory-2.2/file-2.2.1": &fstest.MapFile{},
		"directory-2/directory-2.2/file-2.2.2": &fstest.MapFile{},
	}

	dirEntries, err :=
		fsutils.ReadDirEntriesByKind(mapFS, "directory-2", fsutils.NonDirKind)
	if err != nil {
		log.Fatal(err)
	}

	for _, dirEntry := range dirEntries {
		fmt.Println(dirEntry.Name())
	}

	// Output:
	// file-2.1
	// file-2.2
}
```

## License

The MIT License (MIT)

Copyright &copy; 2023 thewizardplusplus
