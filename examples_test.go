package writablefs_test

import (
	"errors"
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

func ExampleDirFS() {
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

func ExampleDummyFS() {
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
