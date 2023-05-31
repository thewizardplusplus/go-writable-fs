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
