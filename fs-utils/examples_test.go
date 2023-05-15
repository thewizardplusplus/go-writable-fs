package fsutils_test

import (
	"fmt"
	"log"
	"testing/fstest"

	fsutils "github.com/thewizardplusplus/go-writable-fs/fs-utils"
)

func ExampleReadDirEntriesByKind() {
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
