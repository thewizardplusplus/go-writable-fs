package writablefs

import (
	"io/fs"
	"os"
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
