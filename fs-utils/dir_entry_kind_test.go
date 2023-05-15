package fsutils

import (
	"io/fs"
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetDirEntryKind(test *testing.T) {
	type args struct {
		dirEntry fs.DirEntry
	}

	for _, data := range []struct {
		name string
		args args
		want DirEntryKind
	}{
		{
			name: "with a non-directory",
			args: args{
				dirEntry: func() fs.DirEntry {
					mapFS := fstest.MapFS{
						"non-directory": &fstest.MapFile{},
					}

					dirEntries, err := mapFS.ReadDir(".")
					require.NoError(test, err)
					require.Len(test, dirEntries, 1)

					return dirEntries[0]
				}(),
			},
			want: NonDirKind,
		},
		{
			name: "with a directory",
			args: args{
				dirEntry: func() fs.DirEntry {
					mapFS := fstest.MapFS{
						"directory": &fstest.MapFile{Mode: fs.ModeDir},
					}

					dirEntries, err := mapFS.ReadDir(".")
					require.NoError(test, err)
					require.Len(test, dirEntries, 1)

					return dirEntries[0]
				}(),
			},
			want: DirKind,
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			got := GetDirEntryKind(data.args.dirEntry)

			assert.Equal(test, data.want, got)
		})
	}
}
