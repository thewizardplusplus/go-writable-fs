package writablefs

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDirFS(test *testing.T) {
	const baseDir = "base-dir"
	got := NewDirFS(baseDir)

	assert.Equal(test, os.DirFS(baseDir), got.innerDirFS)
	assert.Equal(test, baseDir, got.baseDir)
}

func TestDirFS_Open(test *testing.T) {
	type args struct {
		path string
	}

	for _, data := range []struct {
		name        string
		preparation func(test *testing.T, tempDir string)
		args        args
		want        func(test *testing.T, file fs.File)
		wantErr     func(test *testing.T, tempDir string, err error)
	}{
		{
			name: "success",
			preparation: func(test *testing.T, tempDir string) {
				path := filepath.Join(tempDir, "path")
				err := os.WriteFile(path, []byte("content"), 0666)
				require.NoError(test, err)
			},
			args: args{
				path: "path",
			},
			want: func(test *testing.T, file fs.File) {
				if !assert.NotNil(test, file) {
					return
				}

				content, err := io.ReadAll(file)
				if !assert.NoError(test, err) {
					return
				}

				assert.Equal(test, []byte("content"), content)
			},
			wantErr: func(test *testing.T, _ string, err error) {
				assert.NoError(test, err)
			},
		},
		{
			name:        "error/invalid path",
			preparation: func(_ *testing.T, _ string) {},
			args: args{
				path: "/invalid-path",
			},
			want: func(test *testing.T, file fs.File) {
				assert.Nil(test, file)
			},
			wantErr: func(test *testing.T, _ string, err error) {
				wantPath := "/invalid-path"
				wantErr := &fs.PathError{Op: "open", Path: wantPath, Err: fs.ErrInvalid}
				assert.Equal(test, wantErr, err)
			},
		},
		{
			name:        "error/non-existent path",
			preparation: func(_ *testing.T, _ string) {},
			args: args{
				path: "non-existent-path",
			},
			want: func(test *testing.T, file fs.File) {
				assert.Nil(test, file)
			},
			wantErr: func(test *testing.T, tempDir string, err error) {
				if !assert.IsType(test, (*fs.PathError)(nil), err) {
					return
				}

				typedErr := err.(*fs.PathError)
				assert.Equal(test, "open", typedErr.Op)
				assert.ErrorIs(test, typedErr.Err, fs.ErrNotExist)

				wantPath := filepath.Join(tempDir, "/non-existent-path")
				assert.Equal(test, wantPath, typedErr.Path)
			},
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			tempDir, err := os.MkdirTemp("", "test-*")
			require.NoError(test, err)
			defer os.RemoveAll(tempDir)

			data.preparation(test, tempDir)

			dfs := NewDirFS(tempDir)
			got, err := dfs.Open(data.args.path)
			if got != nil {
				defer got.Close()
			}

			data.want(test, got)
			data.wantErr(test, tempDir, err)
		})
	}
}

func TestDirFS_Mkdir(test *testing.T) {
	type args struct {
		path        string
		permissions fs.FileMode
	}

	for _, data := range []struct {
		name        string
		preparation func(test *testing.T, tempDir string)
		args        args
		want        func(test *testing.T, tempDir string)
		wantErr     func(test *testing.T, tempDir string, err error)
	}{
		{
			name:        "success",
			preparation: func(_ *testing.T, _ string) {},
			args: args{
				path:        "path",
				permissions: 0700,
			},
			want: func(test *testing.T, tempDir string) {
				path := filepath.Join(tempDir, "path")
				if !assert.DirExists(test, path) {
					return
				}

				stat, err := os.Stat(path)
				if !assert.NoError(test, err) {
					return
				}

				assert.Equal(test, 0700|fs.ModeDir, stat.Mode())
			},
			wantErr: func(test *testing.T, _ string, err error) {
				assert.NoError(test, err)
			},
		},
		{
			name:        "error/invalid path",
			preparation: func(_ *testing.T, _ string) {},
			args: args{
				path:        "/invalid-path",
				permissions: 0700,
			},
			want: func(_ *testing.T, _ string) {},
			wantErr: func(test *testing.T, _ string, err error) {
				wantPath := "/invalid-path"
				wantErr := &fs.PathError{Op: "mkdir", Path: wantPath, Err: fs.ErrInvalid}
				assert.Equal(test, wantErr, err)
			},
		},
		{
			name: "error/existent path",
			preparation: func(test *testing.T, tempDir string) {
				path := filepath.Join(tempDir, "existent-path")
				err := os.Mkdir(path, 0700)
				require.NoError(test, err)
			},
			args: args{
				path:        "existent-path",
				permissions: 0700,
			},
			want: func(_ *testing.T, _ string) {},
			wantErr: func(test *testing.T, _ string, err error) {
				if !assert.IsType(test, (*fs.PathError)(nil), err) {
					return
				}

				typedErr := err.(*fs.PathError)
				assert.Equal(test, "mkdir", typedErr.Op)
				assert.Equal(test, "existent-path", typedErr.Path)
				assert.ErrorIs(test, typedErr.Err, fs.ErrExist)
			},
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			tempDir, err := os.MkdirTemp("", "test-*")
			require.NoError(test, err)
			defer os.RemoveAll(tempDir)

			data.preparation(test, tempDir)

			dfs := NewDirFS(tempDir)
			err = dfs.Mkdir(data.args.path, data.args.permissions)

			data.want(test, tempDir)
			data.wantErr(test, tempDir, err)
		})
	}
}

func TestDirFS_Create(test *testing.T) {
	type args struct {
		path string
	}

	for _, data := range []struct {
		name        string
		preparation func(test *testing.T, tempDir string)
		args        args
		want        func(test *testing.T, tempDir string, file fs.File)
		wantErr     func(test *testing.T, err error)
	}{
		{
			name:        "success/non-existent path",
			preparation: func(_ *testing.T, _ string) {},
			args: args{
				path: "non-existent path",
			},
			want: func(test *testing.T, tempDir string, file fs.File) {
				path := filepath.Join(tempDir, "non-existent path")
				if assert.FileExists(test, path) {
					stat, err := os.Stat(path)
					if assert.NoError(test, err) {
						assert.Equal(test, fs.FileMode(0664), stat.Mode())
					}
				}

				if !assert.NotNil(test, file) {
					return
				}

				content, err := io.ReadAll(file)
				if !assert.NoError(test, err) {
					return
				}

				assert.Empty(test, content)
			},
			wantErr: func(test *testing.T, err error) {
				assert.NoError(test, err)
			},
		},
		{
			name: "success/existent path",
			preparation: func(test *testing.T, tempDir string) {
				path := filepath.Join(tempDir, "existent path")
				err := os.WriteFile(path, []byte("content"), 0600)
				require.NoError(test, err)
			},
			args: args{
				path: "existent path",
			},
			want: func(test *testing.T, tempDir string, file fs.File) {
				path := filepath.Join(tempDir, "existent path")
				if assert.FileExists(test, path) {
					stat, err := os.Stat(path)
					if assert.NoError(test, err) {
						assert.Equal(test, fs.FileMode(0600), stat.Mode())
					}
				}

				if !assert.NotNil(test, file) {
					return
				}

				content, err := io.ReadAll(file)
				if !assert.NoError(test, err) {
					return
				}

				assert.Empty(test, content)
			},
			wantErr: func(test *testing.T, err error) {
				assert.NoError(test, err)
			},
		},
		{
			name:        "error/invalid path",
			preparation: func(_ *testing.T, _ string) {},
			args: args{
				path: "/invalid-path",
			},
			want: func(test *testing.T, _ string, file fs.File) {
				assert.Nil(test, file)
			},
			wantErr: func(test *testing.T, err error) {
				wantPath := "/invalid-path"
				wantErr := &fs.PathError{Op: "open", Path: wantPath, Err: fs.ErrInvalid}
				assert.Equal(test, wantErr, err)
			},
		},
		// I added such a strange test case since I couldn't find any other way
		// to make function `os.Create()` return an error
		{
			name:        "error/null-terminated path",
			preparation: func(_ *testing.T, _ string) {},
			args: args{
				path: "null-terminated-path\x00",
			},
			want: func(test *testing.T, _ string, file fs.File) {
				assert.Nil(test, file)
			},
			wantErr: func(test *testing.T, err error) {
				if !assert.IsType(test, (*fs.PathError)(nil), err) {
					return
				}

				typedErr := err.(*fs.PathError)
				assert.Equal(test, "open", typedErr.Op)
				assert.Equal(test, "null-terminated-path\x00", typedErr.Path)
				assert.ErrorIs(test, typedErr.Err, syscall.EINVAL)
			},
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			tempDir, err := os.MkdirTemp("", "test-*")
			require.NoError(test, err)
			defer os.RemoveAll(tempDir)

			data.preparation(test, tempDir)

			dfs := NewDirFS(tempDir)
			got, err := dfs.Create(data.args.path)
			if got != nil {
				defer got.Close()
			}

			data.want(test, tempDir, got)
			data.wantErr(test, err)
		})
	}
}

func TestDirFS_CreateExcl(test *testing.T) {
	type args struct {
		path string
	}

	for _, data := range []struct {
		name        string
		preparation func(test *testing.T, tempDir string)
		args        args
		want        func(test *testing.T, tempDir string, file fs.File)
		wantErr     func(test *testing.T, err error)
	}{
		{
			name:        "success",
			preparation: func(_ *testing.T, _ string) {},
			args: args{
				path: "path",
			},
			want: func(test *testing.T, tempDir string, file fs.File) {
				path := filepath.Join(tempDir, "path")
				if assert.FileExists(test, path) {
					stat, err := os.Stat(path)
					if assert.NoError(test, err) {
						assert.Equal(test, fs.FileMode(0664), stat.Mode())
					}
				}

				if !assert.NotNil(test, file) {
					return
				}

				content, err := io.ReadAll(file)
				if !assert.NoError(test, err) {
					return
				}

				assert.Empty(test, content)
			},
			wantErr: func(test *testing.T, err error) {
				assert.NoError(test, err)
			},
		},
		{
			name:        "error/invalid path",
			preparation: func(_ *testing.T, _ string) {},
			args: args{
				path: "/invalid-path",
			},
			want: func(test *testing.T, _ string, file fs.File) {
				assert.Nil(test, file)
			},
			wantErr: func(test *testing.T, err error) {
				wantPath := "/invalid-path"
				wantErr := &fs.PathError{Op: "open", Path: wantPath, Err: fs.ErrInvalid}
				assert.Equal(test, wantErr, err)
			},
		},
		{
			name: "error/existent path",
			preparation: func(test *testing.T, tempDir string) {
				path := filepath.Join(tempDir, "existent path")
				err := os.WriteFile(path, []byte("content"), 0666)
				require.NoError(test, err)
			},
			args: args{
				path: "existent path",
			},
			want: func(test *testing.T, _ string, file fs.File) {
				assert.Nil(test, file)
			},
			wantErr: func(test *testing.T, err error) {
				if !assert.IsType(test, (*fs.PathError)(nil), err) {
					return
				}

				typedErr := err.(*fs.PathError)
				assert.Equal(test, "open", typedErr.Op)
				assert.Equal(test, "existent path", typedErr.Path)
				assert.ErrorIs(test, typedErr.Err, fs.ErrExist)
			},
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			tempDir, err := os.MkdirTemp("", "test-*")
			require.NoError(test, err)
			defer os.RemoveAll(tempDir)

			data.preparation(test, tempDir)

			dfs := NewDirFS(tempDir)
			got, err := dfs.CreateExcl(data.args.path)
			if got != nil {
				defer got.Close()
			}

			data.want(test, tempDir, got)
			data.wantErr(test, err)
		})
	}
}

func TestDirFS_Remove(test *testing.T) {
	type args struct {
		path string
	}

	for _, data := range []struct {
		name        string
		preparation func(test *testing.T, tempDir string)
		args        args
		want        func(test *testing.T, tempDir string)
		wantErr     func(test *testing.T, err error)
	}{
		{
			name: "success/file",
			preparation: func(test *testing.T, tempDir string) {
				path := filepath.Join(tempDir, "path")
				err := os.WriteFile(path, []byte("content"), 0666)
				require.NoError(test, err)
			},
			args: args{
				path: "path",
			},
			want: func(test *testing.T, tempDir string) {
				path := filepath.Join(tempDir, "path")
				assert.NoFileExists(test, path)
			},
			wantErr: func(test *testing.T, err error) {
				assert.NoError(test, err)
			},
		},
		{
			name: "success/empty directory",
			preparation: func(test *testing.T, tempDir string) {
				path := filepath.Join(tempDir, "path")
				err := os.Mkdir(path, 0700)
				require.NoError(test, err)
			},
			args: args{
				path: "path",
			},
			want: func(test *testing.T, tempDir string) {
				path := filepath.Join(tempDir, "path")
				assert.NoDirExists(test, path)
			},
			wantErr: func(test *testing.T, err error) {
				assert.NoError(test, err)
			},
		},
		{
			name:        "error/invalid path",
			preparation: func(_ *testing.T, _ string) {},
			args: args{
				path: "/invalid-path",
			},
			want: func(_ *testing.T, _ string) {},
			wantErr: func(test *testing.T, err error) {
				wantPath := "/invalid-path"
				wantErr := &fs.PathError{Op: "remove", Path: wantPath, Err: fs.ErrInvalid}
				assert.Equal(test, wantErr, err)
			},
		},
		{
			name:        "error/non-existent path",
			preparation: func(_ *testing.T, _ string) {},
			args: args{
				path: "non-existent path",
			},
			want: func(_ *testing.T, _ string) {},
			wantErr: func(test *testing.T, err error) {
				if !assert.IsType(test, (*fs.PathError)(nil), err) {
					return
				}

				typedErr := err.(*fs.PathError)
				assert.Equal(test, "remove", typedErr.Op)
				assert.Equal(test, "non-existent path", typedErr.Path)
				assert.ErrorIs(test, typedErr.Err, fs.ErrNotExist)
			},
		},
		{
			name: "error/non-empty directory",
			preparation: func(test *testing.T, tempDir string) {
				dirPath := filepath.Join(tempDir, "non-empty-directory")
				err := os.Mkdir(dirPath, 0700)
				require.NoError(test, err)

				filePath := filepath.Join(dirPath, "path")
				err = os.WriteFile(filePath, []byte("content"), 0666)
				require.NoError(test, err)
			},
			args: args{
				path: "non-empty-directory",
			},
			want: func(_ *testing.T, _ string) {},
			wantErr: func(test *testing.T, err error) {
				if !assert.IsType(test, (*fs.PathError)(nil), err) {
					return
				}

				typedErr := err.(*fs.PathError)
				assert.Equal(test, "remove", typedErr.Op)
				assert.Equal(test, "non-empty-directory", typedErr.Path)
				assert.ErrorIs(test, typedErr.Err, syscall.ENOTEMPTY)
			},
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			tempDir, err := os.MkdirTemp("", "test-*")
			require.NoError(test, err)
			defer os.RemoveAll(tempDir)

			data.preparation(test, tempDir)

			dfs := NewDirFS(tempDir)
			err = dfs.Remove(data.args.path)

			data.want(test, tempDir)
			data.wantErr(test, err)
		})
	}
}
