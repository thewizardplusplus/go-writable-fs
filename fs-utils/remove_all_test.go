package fsutils

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	writablefs "github.com/thewizardplusplus/go-writable-fs"
	fsutilsmocks "github.com/thewizardplusplus/go-writable-fs/fs-utils/mocks"
	writablefsmocks "github.com/thewizardplusplus/go-writable-fs/mocks"
)

func TestRemoveAll(test *testing.T) {
	type args struct {
		wfs  func(test *testing.T) writablefs.WritableFS
		path string
	}

	for _, data := range []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success/interface `RemoveAllFS`",
			args: args{
				wfs: func(test *testing.T) writablefs.WritableFS {
					wfs := fsutilsmocks.NewRemoveAllFS(test)
					wfs.EXPECT().
						RemoveAll("path").
						Return(nil)

					return wfs
				},
				path: "path",
			},
			wantErr: assert.NoError,
		},
		{
			name: "success/the path is a file",
			args: args{
				wfs: func(test *testing.T) writablefs.WritableFS {
					fileInfo := writablefsmocks.NewFileInfo(test)
					fileInfo.EXPECT().
						IsDir().
						Return(false)

					file := writablefsmocks.NewWritableFile(test)
					file.EXPECT().
						Stat().
						Return(fileInfo, nil)
					file.EXPECT().
						Close().
						Return(nil)

					wfs := writablefsmocks.NewWritableFS(test)
					wfs.EXPECT().
						Open("path").
						Return(file, nil)
					wfs.EXPECT().
						Remove("path").
						Return(nil)

					return wfs
				},
				path: "path",
			},
			wantErr: assert.NoError,
		},
		{
			name: "success/the path is a directory/single directory",
			args: args{
				wfs: func(test *testing.T) writablefs.WritableFS {
					fileInfo1 := writablefsmocks.NewFileInfo(test)
					fileInfo1.EXPECT().
						IsDir().
						Return(false)
					fileInfo1.EXPECT().
						Name().
						Return("file-one")

					fileInfo2 := writablefsmocks.NewFileInfo(test)
					fileInfo2.EXPECT().
						IsDir().
						Return(false)
					fileInfo2.EXPECT().
						Name().
						Return("file-two")

					directoryFileInfo := writablefsmocks.NewFileInfo(test)
					directoryFileInfo.EXPECT().
						IsDir().
						Return(true)

					directory := fsutilsmocks.NewReadDirWritableFile(test)
					directory.EXPECT().
						Stat().
						Return(directoryFileInfo, nil)
					directory.EXPECT().
						ReadDir(-1).
						Return(
							[]fs.DirEntry{
								fs.FileInfoToDirEntry(fileInfo1),
								fs.FileInfoToDirEntry(fileInfo2),
							},
							nil,
						)
					directory.EXPECT().
						Close().
						Return(nil)

					wfs := writablefsmocks.NewWritableFS(test)
					wfs.EXPECT().
						Open("path").
						Return(directory, nil)
					wfs.EXPECT().
						Remove("path/file-one").
						Return(nil)
					wfs.EXPECT().
						Remove("path/file-two").
						Return(nil)
					wfs.EXPECT().
						Remove("path").
						Return(nil)

					return wfs
				},
				path: "path",
			},
			wantErr: assert.NoError,
		},
		{
			name: "success" +
				"/the path is a directory" +
				"/multiple directories (with a sub-directory)",
			args: args{
				wfs: func(test *testing.T) writablefs.WritableFS {
					fileInfo1 := writablefsmocks.NewFileInfo(test)
					fileInfo1.EXPECT().
						IsDir().
						Return(false)
					fileInfo1.EXPECT().
						Name().
						Return("file-one")

					fileInfo2 := writablefsmocks.NewFileInfo(test)
					fileInfo2.EXPECT().
						IsDir().
						Return(false)
					fileInfo2.EXPECT().
						Name().
						Return("file-two")

					fileInfo3 := writablefsmocks.NewFileInfo(test)
					fileInfo3.EXPECT().
						IsDir().
						Return(false)
					fileInfo3.EXPECT().
						Name().
						Return("file-three")

					fileInfo4 := writablefsmocks.NewFileInfo(test)
					fileInfo4.EXPECT().
						IsDir().
						Return(false)
					fileInfo4.EXPECT().
						Name().
						Return("file-four")

					directoryFileInfo := writablefsmocks.NewFileInfo(test)
					directoryFileInfo.EXPECT().
						IsDir().
						Return(true)

					subDirectoryFileInfo := writablefsmocks.NewFileInfo(test)
					subDirectoryFileInfo.EXPECT().
						IsDir().
						Return(true)
					subDirectoryFileInfo.EXPECT().
						Name().
						Return("sub-path")

					directory := fsutilsmocks.NewReadDirWritableFile(test)
					directory.EXPECT().
						Stat().
						Return(directoryFileInfo, nil)
					directory.EXPECT().
						ReadDir(-1).
						Return(
							[]fs.DirEntry{
								fs.FileInfoToDirEntry(fileInfo1),
								fs.FileInfoToDirEntry(fileInfo2),
								fs.FileInfoToDirEntry(subDirectoryFileInfo),
							},
							nil,
						)
					directory.EXPECT().
						Close().
						Return(nil)

					subDirectory := fsutilsmocks.NewReadDirWritableFile(test)
					subDirectory.EXPECT().
						ReadDir(-1).
						Return(
							[]fs.DirEntry{
								fs.FileInfoToDirEntry(fileInfo3),
								fs.FileInfoToDirEntry(fileInfo4),
							},
							nil,
						)
					subDirectory.EXPECT().
						Close().
						Return(nil)

					wfs := writablefsmocks.NewWritableFS(test)
					wfs.EXPECT().
						Open("path").
						Return(directory, nil)
					wfs.EXPECT().
						Open("path/sub-path").
						Return(subDirectory, nil)
					wfs.EXPECT().
						Remove("path/file-one").
						Return(nil)
					wfs.EXPECT().
						Remove("path/file-two").
						Return(nil)
					wfs.EXPECT().
						Remove("path/sub-path/file-three").
						Return(nil)
					wfs.EXPECT().
						Remove("path/sub-path/file-four").
						Return(nil)
					wfs.EXPECT().
						Remove("path/sub-path").
						Return(nil)
					wfs.EXPECT().
						Remove("path").
						Return(nil)

					return wfs
				},
				path: "path",
			},
			wantErr: assert.NoError,
		},
		{
			name: "error/interface `RemoveAllFS`",
			args: args{
				wfs: func(test *testing.T) writablefs.WritableFS {
					wfs := fsutilsmocks.NewRemoveAllFS(test)
					wfs.EXPECT().
						RemoveAll("path").
						Return(iotest.ErrTimeout)

					return wfs
				},
				path: "path",
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
				path: "/invalid-path",
			},
			wantErr: func(test assert.TestingT, err error, _ ...interface{}) bool {
				wantErr := &fs.PathError{
					Op:   "RemoveAll",
					Path: "/invalid-path",
					Err:  fs.ErrInvalid,
				}
				return assert.Equal(test, wantErr, err)
			},
		},
		{
			name: "error/on root directory reading/`fs.FS.Open()`",
			args: args{
				wfs: func(test *testing.T) writablefs.WritableFS {
					wfs := writablefsmocks.NewWritableFS(test)
					wfs.EXPECT().
						Open("path").
						Return(nil, iotest.ErrTimeout)

					return wfs
				},
				path: "path",
			},
			wantErr: func(test assert.TestingT, err error, _ ...interface{}) bool {
				return assert.ErrorIs(test, err, iotest.ErrTimeout)
			},
		},
		{
			name: "error/on root directory reading/`fs.File.Stat()`",
			args: args{
				wfs: func(test *testing.T) writablefs.WritableFS {
					directory := fsutilsmocks.NewReadDirWritableFile(test)
					directory.EXPECT().
						Stat().
						Return(nil, iotest.ErrTimeout)
					directory.EXPECT().
						Close().
						Return(nil)

					wfs := writablefsmocks.NewWritableFS(test)
					wfs.EXPECT().
						Open("path").
						Return(directory, nil)

					return wfs
				},
				path: "path",
			},
			wantErr: func(test assert.TestingT, err error, _ ...interface{}) bool {
				return assert.ErrorIs(test, err, iotest.ErrTimeout)
			},
		},
		{
			name: "error/on root directory reading/`fs.ReadDirFile.ReadDir()`",
			args: args{
				wfs: func(test *testing.T) writablefs.WritableFS {
					directoryFileInfo := writablefsmocks.NewFileInfo(test)
					directoryFileInfo.EXPECT().
						IsDir().
						Return(true)

					directory := fsutilsmocks.NewReadDirWritableFile(test)
					directory.EXPECT().
						Stat().
						Return(directoryFileInfo, nil)
					directory.EXPECT().
						ReadDir(-1).
						Return(nil, iotest.ErrTimeout)
					directory.EXPECT().
						Close().
						Return(nil)

					wfs := writablefsmocks.NewWritableFS(test)
					wfs.EXPECT().
						Open("path").
						Return(directory, nil)
					wfs.EXPECT().
						Remove("path").
						Return(nil)

					return wfs
				},
				path: "path",
			},
			wantErr: func(test assert.TestingT, err error, _ ...interface{}) bool {
				return assert.ErrorIs(test, err, iotest.ErrTimeout)
			},
		},
		{
			name: "error/on sub-directory reading/`fs.FS.Open()`",
			args: args{
				wfs: func(test *testing.T) writablefs.WritableFS {
					fileInfo1 := writablefsmocks.NewFileInfo(test)
					fileInfo1.EXPECT().
						IsDir().
						Return(false)
					fileInfo1.EXPECT().
						Name().
						Return("file-one")

					fileInfo2 := writablefsmocks.NewFileInfo(test)
					fileInfo2.EXPECT().
						IsDir().
						Return(false)
					fileInfo2.EXPECT().
						Name().
						Return("file-two")

					directoryFileInfo := writablefsmocks.NewFileInfo(test)
					directoryFileInfo.EXPECT().
						IsDir().
						Return(true)

					subDirectoryFileInfo := writablefsmocks.NewFileInfo(test)
					subDirectoryFileInfo.EXPECT().
						IsDir().
						Return(true)
					subDirectoryFileInfo.EXPECT().
						Name().
						Return("sub-path")

					directory := fsutilsmocks.NewReadDirWritableFile(test)
					directory.EXPECT().
						Stat().
						Return(directoryFileInfo, nil)
					directory.EXPECT().
						ReadDir(-1).
						Return(
							[]fs.DirEntry{
								fs.FileInfoToDirEntry(fileInfo1),
								fs.FileInfoToDirEntry(fileInfo2),
								fs.FileInfoToDirEntry(subDirectoryFileInfo),
							},
							nil,
						)
					directory.EXPECT().
						Close().
						Return(nil)

					wfs := writablefsmocks.NewWritableFS(test)
					wfs.EXPECT().
						Open("path").
						Return(directory, nil)
					wfs.EXPECT().
						Open("path/sub-path").
						Return(nil, iotest.ErrTimeout)
					wfs.EXPECT().
						Remove("path/file-one").
						Return(nil)
					wfs.EXPECT().
						Remove("path/file-two").
						Return(nil)
					wfs.EXPECT().
						Remove("path/sub-path").
						Return(nil)
					wfs.EXPECT().
						Remove("path").
						Return(nil)

					return wfs
				},
				path: "path",
			},
			wantErr: func(test assert.TestingT, err error, _ ...interface{}) bool {
				return assert.ErrorIs(test, err, iotest.ErrTimeout)
			},
		},
		{
			name: "error/on sub-directory reading/`fs.ReadDirFile.ReadDir()`",
			args: args{
				wfs: func(test *testing.T) writablefs.WritableFS {
					fileInfo1 := writablefsmocks.NewFileInfo(test)
					fileInfo1.EXPECT().
						IsDir().
						Return(false)
					fileInfo1.EXPECT().
						Name().
						Return("file-one")

					fileInfo2 := writablefsmocks.NewFileInfo(test)
					fileInfo2.EXPECT().
						IsDir().
						Return(false)
					fileInfo2.EXPECT().
						Name().
						Return("file-two")

					directoryFileInfo := writablefsmocks.NewFileInfo(test)
					directoryFileInfo.EXPECT().
						IsDir().
						Return(true)

					subDirectoryFileInfo := writablefsmocks.NewFileInfo(test)
					subDirectoryFileInfo.EXPECT().
						IsDir().
						Return(true)
					subDirectoryFileInfo.EXPECT().
						Name().
						Return("sub-path")

					directory := fsutilsmocks.NewReadDirWritableFile(test)
					directory.EXPECT().
						Stat().
						Return(directoryFileInfo, nil)
					directory.EXPECT().
						ReadDir(-1).
						Return(
							[]fs.DirEntry{
								fs.FileInfoToDirEntry(fileInfo1),
								fs.FileInfoToDirEntry(fileInfo2),
								fs.FileInfoToDirEntry(subDirectoryFileInfo),
							},
							nil,
						)
					directory.EXPECT().
						Close().
						Return(nil)

					subDirectory := fsutilsmocks.NewReadDirWritableFile(test)
					subDirectory.EXPECT().
						ReadDir(-1).
						Return(nil, iotest.ErrTimeout)
					subDirectory.EXPECT().
						Close().
						Return(nil)

					wfs := writablefsmocks.NewWritableFS(test)
					wfs.EXPECT().
						Open("path").
						Return(directory, nil)
					wfs.EXPECT().
						Open("path/sub-path").
						Return(subDirectory, nil)
					wfs.EXPECT().
						Remove("path/file-one").
						Return(nil)
					wfs.EXPECT().
						Remove("path/file-two").
						Return(nil)
					wfs.EXPECT().
						Remove("path/sub-path").
						Return(nil)
					wfs.EXPECT().
						Remove("path").
						Return(nil)

					return wfs
				},
				path: "path",
			},
			wantErr: func(test assert.TestingT, err error, _ ...interface{}) bool {
				return assert.ErrorIs(test, err, iotest.ErrTimeout)
			},
		},
		{
			name: "error/on file removing",
			args: args{
				wfs: func(test *testing.T) writablefs.WritableFS {
					fileInfo1 := writablefsmocks.NewFileInfo(test)
					fileInfo1.EXPECT().
						IsDir().
						Return(false)
					fileInfo1.EXPECT().
						Name().
						Return("file-one")

					fileInfo2 := writablefsmocks.NewFileInfo(test)
					fileInfo2.EXPECT().
						IsDir().
						Return(false)
					fileInfo2.EXPECT().
						Name().
						Return("file-two")

					fileInfo3 := writablefsmocks.NewFileInfo(test)
					fileInfo3.EXPECT().
						IsDir().
						Return(false)
					fileInfo3.EXPECT().
						Name().
						Return("file-three")

					fileInfo4 := writablefsmocks.NewFileInfo(test)
					fileInfo4.EXPECT().
						IsDir().
						Return(false)
					fileInfo4.EXPECT().
						Name().
						Return("file-four")

					directoryFileInfo := writablefsmocks.NewFileInfo(test)
					directoryFileInfo.EXPECT().
						IsDir().
						Return(true)

					subDirectoryFileInfo := writablefsmocks.NewFileInfo(test)
					subDirectoryFileInfo.EXPECT().
						IsDir().
						Return(true)
					subDirectoryFileInfo.EXPECT().
						Name().
						Return("sub-path")

					directory := fsutilsmocks.NewReadDirWritableFile(test)
					directory.EXPECT().
						Stat().
						Return(directoryFileInfo, nil)
					directory.EXPECT().
						ReadDir(-1).
						Return(
							[]fs.DirEntry{
								fs.FileInfoToDirEntry(fileInfo1),
								fs.FileInfoToDirEntry(fileInfo2),
								fs.FileInfoToDirEntry(subDirectoryFileInfo),
							},
							nil,
						)
					directory.EXPECT().
						Close().
						Return(nil)

					subDirectory := fsutilsmocks.NewReadDirWritableFile(test)
					subDirectory.EXPECT().
						ReadDir(-1).
						Return(
							[]fs.DirEntry{
								fs.FileInfoToDirEntry(fileInfo3),
								fs.FileInfoToDirEntry(fileInfo4),
							},
							nil,
						)
					subDirectory.EXPECT().
						Close().
						Return(nil)

					wfs := writablefsmocks.NewWritableFS(test)
					wfs.EXPECT().
						Open("path").
						Return(directory, nil)
					wfs.EXPECT().
						Open("path/sub-path").
						Return(subDirectory, nil)
					wfs.EXPECT().
						Remove("path/file-one").
						Return(errors.New("unable to remove the file #1"))
					wfs.EXPECT().
						Remove("path/file-two").
						Return(errors.New("unable to remove the file #2"))
					wfs.EXPECT().
						Remove("path/sub-path/file-three").
						Return(errors.New("unable to remove the file #3"))
					wfs.EXPECT().
						Remove("path/sub-path/file-four").
						Return(errors.New("unable to remove the file #4"))
					wfs.EXPECT().
						Remove("path/sub-path").
						Return(nil)
					wfs.EXPECT().
						Remove("path").
						Return(nil)

					return wfs
				},
				path: "path",
			},
			wantErr: func(test assert.TestingT, err error, _ ...interface{}) bool {
				return assert.EqualError(test, err, "unable to remove the file #1")
			},
		},
		{
			name: "error/on directory removing",
			args: args{
				wfs: func(test *testing.T) writablefs.WritableFS {
					fileInfo1 := writablefsmocks.NewFileInfo(test)
					fileInfo1.EXPECT().
						IsDir().
						Return(false)
					fileInfo1.EXPECT().
						Name().
						Return("file-one")

					fileInfo2 := writablefsmocks.NewFileInfo(test)
					fileInfo2.EXPECT().
						IsDir().
						Return(false)
					fileInfo2.EXPECT().
						Name().
						Return("file-two")

					fileInfo3 := writablefsmocks.NewFileInfo(test)
					fileInfo3.EXPECT().
						IsDir().
						Return(false)
					fileInfo3.EXPECT().
						Name().
						Return("file-three")

					fileInfo4 := writablefsmocks.NewFileInfo(test)
					fileInfo4.EXPECT().
						IsDir().
						Return(false)
					fileInfo4.EXPECT().
						Name().
						Return("file-four")

					directoryFileInfo := writablefsmocks.NewFileInfo(test)
					directoryFileInfo.EXPECT().
						IsDir().
						Return(true)

					subDirectoryFileInfo := writablefsmocks.NewFileInfo(test)
					subDirectoryFileInfo.EXPECT().
						IsDir().
						Return(true)
					subDirectoryFileInfo.EXPECT().
						Name().
						Return("sub-path")

					directory := fsutilsmocks.NewReadDirWritableFile(test)
					directory.EXPECT().
						Stat().
						Return(directoryFileInfo, nil)
					directory.EXPECT().
						ReadDir(-1).
						Return(
							[]fs.DirEntry{
								fs.FileInfoToDirEntry(fileInfo1),
								fs.FileInfoToDirEntry(fileInfo2),
								fs.FileInfoToDirEntry(subDirectoryFileInfo),
							},
							nil,
						)
					directory.EXPECT().
						Close().
						Return(nil)

					subDirectory := fsutilsmocks.NewReadDirWritableFile(test)
					subDirectory.EXPECT().
						ReadDir(-1).
						Return(
							[]fs.DirEntry{
								fs.FileInfoToDirEntry(fileInfo3),
								fs.FileInfoToDirEntry(fileInfo4),
							},
							nil,
						)
					subDirectory.EXPECT().
						Close().
						Return(nil)

					wfs := writablefsmocks.NewWritableFS(test)
					wfs.EXPECT().
						Open("path").
						Return(directory, nil)
					wfs.EXPECT().
						Open("path/sub-path").
						Return(subDirectory, nil)
					wfs.EXPECT().
						Remove("path/file-one").
						Return(nil)
					wfs.EXPECT().
						Remove("path/file-two").
						Return(nil)
					wfs.EXPECT().
						Remove("path/sub-path/file-three").
						Return(nil)
					wfs.EXPECT().
						Remove("path/sub-path/file-four").
						Return(nil)
					wfs.EXPECT().
						Remove("path/sub-path").
						Return(errors.New("unable to remove the directory #1"))
					wfs.EXPECT().
						Remove("path").
						Return(errors.New("unable to remove the directory #2"))

					return wfs
				},
				path: "path",
			},
			wantErr: func(test assert.TestingT, err error, _ ...interface{}) bool {
				return assert.EqualError(test, err, "unable to remove the directory #1")
			},
		},
		{
			name: "error/on each step",
			args: args{
				wfs: func(test *testing.T) writablefs.WritableFS {
					fileInfo1 := writablefsmocks.NewFileInfo(test)
					fileInfo1.EXPECT().
						IsDir().
						Return(false)
					fileInfo1.EXPECT().
						Name().
						Return("file-one")

					fileInfo2 := writablefsmocks.NewFileInfo(test)
					fileInfo2.EXPECT().
						IsDir().
						Return(false)
					fileInfo2.EXPECT().
						Name().
						Return("file-two")

					directoryFileInfo := writablefsmocks.NewFileInfo(test)
					directoryFileInfo.EXPECT().
						IsDir().
						Return(true)

					subDirectoryFileInfo := writablefsmocks.NewFileInfo(test)
					subDirectoryFileInfo.EXPECT().
						IsDir().
						Return(true)
					subDirectoryFileInfo.EXPECT().
						Name().
						Return("sub-path")

					directory := fsutilsmocks.NewReadDirWritableFile(test)
					directory.EXPECT().
						Stat().
						Return(directoryFileInfo, nil)
					directory.EXPECT().
						ReadDir(-1).
						Return(
							[]fs.DirEntry{
								fs.FileInfoToDirEntry(fileInfo1),
								fs.FileInfoToDirEntry(fileInfo2),
								fs.FileInfoToDirEntry(subDirectoryFileInfo),
							},
							nil,
						)
					directory.EXPECT().
						Close().
						Return(nil)

					wfs := writablefsmocks.NewWritableFS(test)
					wfs.EXPECT().
						Open("path").
						Return(directory, nil)
					wfs.EXPECT().
						Open("path/sub-path").
						Return(nil, errors.New("unable to open the sub-directory"))
					wfs.EXPECT().
						Remove("path/file-one").
						Return(errors.New("unable to remove the file #1"))
					wfs.EXPECT().
						Remove("path/file-two").
						Return(errors.New("unable to remove the file #2"))
					wfs.EXPECT().
						Remove("path/sub-path").
						Return(errors.New("unable to remove the directory #1"))
					wfs.EXPECT().
						Remove("path").
						Return(errors.New("unable to remove the directory #2"))

					return wfs
				},
				path: "path",
			},
			wantErr: func(test assert.TestingT, err error, _ ...interface{}) bool {
				return assert.EqualError(test, err, "unable to open the sub-directory")
			},
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			err := RemoveAll(data.args.wfs(test), data.args.path)

			data.wantErr(test, err)
		})
	}
}

func TestRemoveAll_withDirFS(test *testing.T) {
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
			name: "success/the path is a file",
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
			name: "success/the path is a directory/single directory",
			preparation: func(test *testing.T, tempDir string) {
				dirPath := filepath.Join(tempDir, "path")
				err := os.Mkdir(dirPath, 0700)
				require.NoError(test, err)

				fileOnePath := filepath.Join(dirPath, "file-one")
				err = os.WriteFile(fileOnePath, []byte("content-one"), 0666)
				require.NoError(test, err)

				fileTwoPath := filepath.Join(dirPath, "file-two")
				err = os.WriteFile(fileTwoPath, []byte("content-two"), 0666)
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
			name: "success" +
				"/the path is a directory" +
				"/multiple directories (with a sub-directory)",
			preparation: func(test *testing.T, tempDir string) {
				dirPath := filepath.Join(tempDir, "path")
				err := os.Mkdir(dirPath, 0700)
				require.NoError(test, err)

				fileOnePath := filepath.Join(dirPath, "file-one")
				err = os.WriteFile(fileOnePath, []byte("content-one"), 0666)
				require.NoError(test, err)

				fileTwoPath := filepath.Join(dirPath, "file-two")
				err = os.WriteFile(fileTwoPath, []byte("content-two"), 0666)
				require.NoError(test, err)

				subDirPath := filepath.Join(dirPath, "sub-path")
				err = os.Mkdir(subDirPath, 0700)
				require.NoError(test, err)

				fileThreePath := filepath.Join(subDirPath, "file-three")
				err = os.WriteFile(fileThreePath, []byte("content-three"), 0666)
				require.NoError(test, err)

				fileFourPath := filepath.Join(subDirPath, "file-four")
				err = os.WriteFile(fileFourPath, []byte("content-four"), 0666)
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
			name:        "error/non-existent root directory",
			preparation: func(_ *testing.T, _ string) {},
			args: args{
				path: "path",
			},
			want: func(_ *testing.T, _ string) {},
			wantErr: func(test *testing.T, err error) {
				assert.ErrorIs(test, err, fs.ErrNotExist)
			},
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			tempDir, err := os.MkdirTemp("", "test-*")
			require.NoError(test, err)
			defer os.RemoveAll(tempDir)

			data.preparation(test, tempDir)

			dfs := writablefs.NewDirFS(tempDir)
			err = RemoveAll(dfs, data.args.path)

			data.want(test, tempDir)
			data.wantErr(test, err)
		})
	}
}
