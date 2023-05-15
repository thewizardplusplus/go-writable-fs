package fsutils

import (
	"io/fs"
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadDirEntriesByKind(test *testing.T) {
	type args struct {
		fsInstance fs.FS
		path       string
		kind       DirEntryKind
	}

	for _, data := range []struct {
		name    string
		args    args
		want    []fs.DirEntry
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success with kind `NonDirKind`",
			args: args{
				fsInstance: fstest.MapFS{
					"directory-1/file-1.1":                 &fstest.MapFile{},
					"directory-1/file-1.2":                 &fstest.MapFile{},
					"directory-2/file-2.1":                 &fstest.MapFile{},
					"directory-2/file-2.2":                 &fstest.MapFile{},
					"directory-2/directory-2.1/file-2.1.1": &fstest.MapFile{},
					"directory-2/directory-2.1/file-2.1.2": &fstest.MapFile{},
					"directory-2/directory-2.2/file-2.2.1": &fstest.MapFile{},
					"directory-2/directory-2.2/file-2.2.2": &fstest.MapFile{},
				},
				path: "directory-2",
				kind: NonDirKind,
			},
			want: func() []fs.DirEntry {
				mapFS := fstest.MapFS{
					"file-2.1": &fstest.MapFile{},
					"file-2.2": &fstest.MapFile{},
				}

				dirEntries, err := mapFS.ReadDir(".")
				require.NoError(test, err)

				return dirEntries
			}(),
			wantErr: assert.NoError,
		},
		{
			name: "success with kind `DirKind`",
			args: args{
				fsInstance: fstest.MapFS{
					"directory-1/file-1.1":                 &fstest.MapFile{},
					"directory-1/file-1.2":                 &fstest.MapFile{},
					"directory-2/file-2.1":                 &fstest.MapFile{},
					"directory-2/file-2.2":                 &fstest.MapFile{},
					"directory-2/directory-2.1/file-2.1.1": &fstest.MapFile{},
					"directory-2/directory-2.1/file-2.1.2": &fstest.MapFile{},
					"directory-2/directory-2.2/file-2.2.1": &fstest.MapFile{},
					"directory-2/directory-2.2/file-2.2.2": &fstest.MapFile{},
				},
				path: "directory-2",
				kind: DirKind,
			},
			want: func() []fs.DirEntry {
				mapFS := fstest.MapFS{
					"directory-2.1": &fstest.MapFile{Mode: fs.ModeDir},
					"directory-2.2": &fstest.MapFile{Mode: fs.ModeDir},
				}

				dirEntries, err := mapFS.ReadDir(".")
				require.NoError(test, err)

				return dirEntries
			}(),
			wantErr: assert.NoError,
		},
		{
			name: "error",
			args: args{
				fsInstance: fstest.MapFS{
					"directory-1/file-1.1":                 &fstest.MapFile{},
					"directory-1/file-1.2":                 &fstest.MapFile{},
					"directory-2/file-2.1":                 &fstest.MapFile{},
					"directory-2/file-2.2":                 &fstest.MapFile{},
					"directory-2/directory-2.1/file-2.1.1": &fstest.MapFile{},
					"directory-2/directory-2.1/file-2.1.2": &fstest.MapFile{},
					"directory-2/directory-2.2/file-2.2.1": &fstest.MapFile{},
					"directory-2/directory-2.2/file-2.2.2": &fstest.MapFile{},
				},
				path: "non-existing-entry",
				kind: NonDirKind,
			},
			want:    nil,
			wantErr: assert.Error,
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			got, err := ReadDirEntriesByKind(
				data.args.fsInstance,
				data.args.path,
				data.args.kind,
			)

			assert.Equal(test, data.want, got)
			data.wantErr(test, err)
		})
	}
}
