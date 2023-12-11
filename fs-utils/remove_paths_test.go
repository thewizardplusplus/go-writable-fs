package fsutils

import (
	"errors"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/assert"
	writablefs "github.com/thewizardplusplus/go-writable-fs"
	fsutilsmocks "github.com/thewizardplusplus/go-writable-fs/fs-utils/mocks"
)

func Test_removePaths(test *testing.T) {
	type args struct {
		wfs   func(test *testing.T) writablefs.WritableFS
		paths []string
	}

	for _, data := range []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success/no paths",
			args: args{
				wfs: func(test *testing.T) writablefs.WritableFS {
					return fsutilsmocks.NewMkdirTempFS(test)
				},
				paths: nil,
			},
			wantErr: assert.NoError,
		},
		{
			name: "success/no errors",
			args: args{
				wfs: func(test *testing.T) writablefs.WritableFS {
					wfs := fsutilsmocks.NewMkdirAllFS(test)
					wfs.EXPECT().
						Remove("path-one/path-two/path-three").
						Return(nil)
					wfs.EXPECT().
						Remove("path-one/path-two").
						Return(nil)
					wfs.EXPECT().
						Remove("path-one").
						Return(nil)

					return wfs
				},
				paths: []string{
					"path-one/path-two/path-three",
					"path-one/path-two",
					"path-one",
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "error/single error",
			args: args{
				wfs: func(test *testing.T) writablefs.WritableFS {
					wfs := fsutilsmocks.NewMkdirAllFS(test)
					wfs.EXPECT().
						Remove("path-one/path-two/path-three").
						Return(nil)
					wfs.EXPECT().
						Remove("path-one/path-two").
						Return(iotest.ErrTimeout)
					wfs.EXPECT().
						Remove("path-one").
						Return(nil)

					return wfs
				},
				paths: []string{
					"path-one/path-two/path-three",
					"path-one/path-two",
					"path-one",
				},
			},
			wantErr: func(test assert.TestingT, err error, _ ...interface{}) bool {
				return assert.ErrorIs(test, err, iotest.ErrTimeout)
			},
		},
		{
			name: "error/multiple errors",
			args: args{
				wfs: func(test *testing.T) writablefs.WritableFS {
					wfs := fsutilsmocks.NewMkdirAllFS(test)
					wfs.EXPECT().
						Remove("path-one/path-two/path-three").
						Return(nil)
					wfs.EXPECT().
						Remove("path-one/path-two").
						Return(errors.New("error #1"))
					wfs.EXPECT().
						Remove("path-one").
						Return(errors.New("error #2"))

					return wfs
				},
				paths: []string{
					"path-one/path-two/path-three",
					"path-one/path-two",
					"path-one",
				},
			},
			wantErr: func(test assert.TestingT, err error, _ ...interface{}) bool {
				return assert.EqualError(test, err, "error #1")
			},
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			err := removePaths(data.args.wfs(test), data.args.paths)

			data.wantErr(test, err)
		})
	}
}
