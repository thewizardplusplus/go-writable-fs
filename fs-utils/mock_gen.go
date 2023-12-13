package fsutils

import (
	"io/fs"

	writablefs "github.com/thewizardplusplus/go-writable-fs"
)

//go:generate mockery --name=ReadDirWritableFile --case=underscore --with-expecter

// Interface `ReadDirWritableFile` is only used for generating mocks.
type ReadDirWritableFile interface {
	writablefs.WritableFile
	fs.ReadDirFile
}
