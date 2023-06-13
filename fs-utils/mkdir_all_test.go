package fsutils

import (
	"io/fs"
	"os"
	"path/filepath"
	"syscall"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	writablefs "github.com/thewizardplusplus/go-writable-fs"
	fsutilsmocks "github.com/thewizardplusplus/go-writable-fs/fs-utils/mocks"
	writablefsmocks "github.com/thewizardplusplus/go-writable-fs/mocks"
)

func TestMkdirAll(test *testing.T) {
	type args struct {
		wfs         func(test *testing.T) writablefs.WritableFS
		path        string
		permissions fs.FileMode
	}

	for _, data := range []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success/interface `MkdirAllFS`",
			args: args{
				wfs: func(test *testing.T) writablefs.WritableFS {
					wfs := fsutilsmocks.NewMkdirAllFS(test)
					wfs.EXPECT().
						MkdirAll("path-one/path-two/path-three", fs.FileMode(0700)).
						Return(nil)

					return wfs
				},
				path:        "path-one/path-two/path-three",
				permissions: 0700,
			},
			wantErr: assert.NoError,
		},
		{
			name: "success/no errors",
			args: args{
				wfs: func(test *testing.T) writablefs.WritableFS {
					wfs := writablefsmocks.NewWritableFS(test)
					wfs.EXPECT().Mkdir("path-one", fs.FileMode(0700)).Return(nil)
					wfs.EXPECT().Mkdir("path-one/path-two", fs.FileMode(0700)).Return(nil)
					wfs.EXPECT().
						Mkdir("path-one/path-two/path-three", fs.FileMode(0700)).
						Return(nil)

					return wfs
				},
				path:        "path-one/path-two/path-three",
				permissions: 0700,
			},
			wantErr: assert.NoError,
		},
		{
			name: "success/one path element",
			args: args{
				wfs: func(test *testing.T) writablefs.WritableFS {
					wfs := writablefsmocks.NewWritableFS(test)
					wfs.EXPECT().Mkdir("path", fs.FileMode(0700)).Return(nil)

					return wfs
				},
				path:        "path",
				permissions: 0700,
			},
			wantErr: assert.NoError,
		},
		{
			name: "success/error `fs.ErrExist`",
			args: args{
				wfs: func(test *testing.T) writablefs.WritableFS {
					wfs := writablefsmocks.NewWritableFS(test)
					wfs.EXPECT().Mkdir("path-one", fs.FileMode(0700)).Return(fs.ErrExist)
					wfs.EXPECT().
						Mkdir("path-one/path-two", fs.FileMode(0700)).
						Return(fs.ErrExist)
					wfs.EXPECT().
						Mkdir("path-one/path-two/path-three", fs.FileMode(0700)).
						Return(nil)

					return wfs
				},
				path:        "path-one/path-two/path-three",
				permissions: 0700,
			},
			wantErr: assert.NoError,
		},
		{
			name: "error/interface `MkdirAllFS`",
			args: args{
				wfs: func(test *testing.T) writablefs.WritableFS {
					wfs := fsutilsmocks.NewMkdirAllFS(test)
					wfs.EXPECT().
						MkdirAll("path-one/path-two/path-three", fs.FileMode(0700)).
						Return(iotest.ErrTimeout)

					return wfs
				},
				path:        "path-one/path-two/path-three",
				permissions: 0700,
			},
			wantErr: func(test assert.TestingT, err error, _ ...interface{}) bool {
				return assert.ErrorIs(test, err, iotest.ErrTimeout)
			},
		},
		{
			name: "error/invalid path",
			args: args{
				wfs: func(test *testing.T) writablefs.WritableFS {
					return writablefsmocks.NewWritableFS(test)
				},
				path:        "/invalid-path",
				permissions: 0700,
			},
			wantErr: func(test assert.TestingT, err error, _ ...interface{}) bool {
				wantPath := "/invalid-path"
				wantErr := &fs.PathError{Op: "mkdir", Path: wantPath, Err: fs.ErrInvalid}
				return assert.Equal(test, wantErr, err)
			},
		},
		{
			name: "error/not error `fs.ErrExist`",
			args: args{
				wfs: func(test *testing.T) writablefs.WritableFS {
					wfs := writablefsmocks.NewWritableFS(test)
					wfs.EXPECT().Mkdir("path-one", fs.FileMode(0700)).Return(nil)
					wfs.EXPECT().
						Mkdir("path-one/path-two", fs.FileMode(0700)).
						Return(iotest.ErrTimeout)

					return wfs
				},
				path:        "path-one/path-two/path-three",
				permissions: 0700,
			},
			wantErr: func(test assert.TestingT, err error, _ ...interface{}) bool {
				return assert.ErrorIs(test, err, iotest.ErrTimeout)
			},
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			err := MkdirAll(data.args.wfs(test), data.args.path, data.args.permissions)

			data.wantErr(test, err)
		})
	}
}

func TestMkdirAll_withDirFS(test *testing.T) {
	type args struct {
		path        string
		permissions fs.FileMode
	}

	for _, data := range []struct {
		name        string
		preparation func(test *testing.T, tempDir string)
		args        args
		want        func(test *testing.T, tempDir string)
		wantErr     func(test *testing.T, err error)
	}{
		{
			name:        "success/non-existent path elements",
			preparation: func(_ *testing.T, _ string) {},
			args: args{
				path:        "path-one/path-two/path-three",
				permissions: 0700,
			},
			want: func(test *testing.T, tempDir string) {
				fullPath := tempDir
				for _, pathElement := range []string{"path-one", "path-two", "path-three"} {
					fullPath = filepath.Join(fullPath, pathElement)

					if !assert.DirExists(test, fullPath) {
						continue
					}

					stat, err := os.Stat(fullPath)
					if !assert.NoError(test, err) {
						continue
					}

					assert.Equal(test, 0700|fs.ModeDir, stat.Mode())
				}
			},
			wantErr: func(test *testing.T, err error) {
				assert.NoError(test, err)
			},
		},
		{
			name: "success/existent path elements",
			preparation: func(test *testing.T, tempDir string) {
				path := filepath.Join(tempDir, "path-one/path-two")
				err := os.MkdirAll(path, 0700)
				require.NoError(test, err)
			},
			args: args{
				path:        "path-one/path-two/path-three",
				permissions: 0700,
			},
			want: func(test *testing.T, tempDir string) {
				path := filepath.Join(tempDir, "path-one/path-two/path-three")
				if !assert.DirExists(test, path) {
					return
				}

				stat, err := os.Stat(path)
				if !assert.NoError(test, err) {
					return
				}

				assert.Equal(test, 0700|fs.ModeDir, stat.Mode())
			},
			wantErr: func(test *testing.T, err error) {
				assert.NoError(test, err)
			},
		},
		{
			name: "error/non-directory path element",
			preparation: func(test *testing.T, tempDir string) {
				dirPath := filepath.Join(tempDir, "path-one")
				err := os.Mkdir(dirPath, 0700)
				require.NoError(test, err)

				filePath := filepath.Join(dirPath, "path-two")
				err = os.WriteFile(filePath, []byte("content"), 0666)
				require.NoError(test, err)
			},
			args: args{
				path:        "path-one/path-two/path-three",
				permissions: 0700,
			},
			want: func(_ *testing.T, _ string) {},
			wantErr: func(test *testing.T, err error) {
				if !assert.IsType(test, (*fs.PathError)(nil), err) {
					return
				}

				typedErr := err.(*fs.PathError)
				assert.Equal(test, "mkdir", typedErr.Op)
				assert.Equal(test, "path-one/path-two/path-three", typedErr.Path)
				assert.ErrorIs(test, typedErr.Err, syscall.ENOTDIR)
			},
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			tempDir, err := os.MkdirTemp("", "test-*")
			require.NoError(test, err)
			defer os.RemoveAll(tempDir)

			data.preparation(test, tempDir)

			dfs := writablefs.NewDirFS(tempDir)
			err = MkdirAll(dfs, data.args.path, data.args.permissions)

			data.want(test, tempDir)
			data.wantErr(test, err)
		})
	}
}
