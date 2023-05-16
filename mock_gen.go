package writablefs

import (
	"io/fs"
)

//go:generate mockery --name=FileInfo --case=underscore

// Interface `FileInfo` is only used for generating mocks.
type FileInfo interface {
	fs.FileInfo
}
