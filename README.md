# go-writable-fs

[![GoDoc](https://godoc.org/github.com/thewizardplusplus/go-writable-fs?status.svg)](https://godoc.org/github.com/thewizardplusplus/go-writable-fs)
[![Go Report Card](https://goreportcard.com/badge/github.com/thewizardplusplus/go-writable-fs)](https://goreportcard.com/report/github.com/thewizardplusplus/go-writable-fs)

## Installation

```
$ go get github.com/thewizardplusplus/go-writable-fs@latest
```

## Examples

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
