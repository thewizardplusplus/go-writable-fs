package fsutils

import (
	"bytes"
	cryptorand "crypto/rand"
	"io"
	"io/fs"
	mathrand "math/rand"
	"os"
	"path/filepath"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	writablefs "github.com/thewizardplusplus/go-writable-fs"
	fsutilsmocks "github.com/thewizardplusplus/go-writable-fs/fs-utils/mocks"
	mocks "github.com/thewizardplusplus/go-writable-fs/mocks"
	writablefsmocks "github.com/thewizardplusplus/go-writable-fs/mocks"
)

func TestCreateTemp(test *testing.T) {
	type args struct {
		wfs         func(test *testing.T) writablefs.WritableFS
		baseDir     string
		pathPattern string
	}

	for _, data := range []struct {
		name         string
		randomReader io.Reader
		args         args
		want         assert.ValueAssertionFunc
		wantErr      assert.ErrorAssertionFunc
	}{
		{
			name:         "success/interface `CreateTempFS`",
			randomReader: nil,
			args: args{
				wfs: func(test *testing.T) writablefs.WritableFS {
					file := mocks.NewWritableFile(test)

					wfs := fsutilsmocks.NewCreateTempFS(test)
					wfs.EXPECT().
						CreateTemp("base-dir", "path-pattern-*").
						Return(file, nil)

					return wfs
				},
				baseDir:     "base-dir",
				pathPattern: "path-pattern-*",
			},
			want:    assert.NotNil,
			wantErr: assert.NoError,
		},
		{
			name: "success/no errors",
			randomReader: func() io.Reader {
				randomBuffer := make([]byte, maxRandomSuffixValue.BitLen()/8)
				_, err := mathrand.New(mathrand.NewSource(1)).Read(randomBuffer)
				require.NoError(test, err)

				return bytes.NewBuffer(randomBuffer)
			}(),
			args: args{
				wfs: func(test *testing.T) writablefs.WritableFS {
					file := mocks.NewWritableFile(test)

					wfs := writablefsmocks.NewWritableFS(test)
					wfs.EXPECT().
						CreateExcl("base-dir/path-pattern-5980212987775051087").
						Return(file, nil)

					return wfs
				},
				baseDir:     "base-dir",
				pathPattern: "path-pattern-*",
			},
			want:    assert.NotNil,
			wantErr: assert.NoError,
		},
		{
			name: "success" +
				"/error `fs.ErrExist`" +
				"/the count of tries is less than the maximum",
			randomReader: func() io.Reader {
				randomBuffer := make([]byte, maxRandomSuffixValue.BitLen()/8)
				_, err := mathrand.New(mathrand.NewSource(1)).Read(randomBuffer)
				require.NoError(test, err)

				maxCountOfTempTries := 5
				repeatedRandomBuffer := bytes.Repeat(randomBuffer, maxCountOfTempTries+1)
				return bytes.NewBuffer(repeatedRandomBuffer)
			}(),
			args: args{
				wfs: func(test *testing.T) writablefs.WritableFS {
					wantPath := "base-dir/path-pattern-5980212987775051087"
					maxCountOfTempTries := 5

					file := mocks.NewWritableFile(test)

					var callCount int
					wfs := writablefsmocks.NewWritableFS(test)
					wfs.
						On("CreateExcl", wantPath).
						Return(func(path string) (writablefs.WritableFile, error) {
							if callCount++; callCount > maxCountOfTempTries {
								return file, nil
							}

							return nil, fs.ErrExist
						}).
						Times(maxCountOfTempTries + 1)

					return wfs
				},
				baseDir:     "base-dir",
				pathPattern: "path-pattern-*",
			},
			want:    assert.NotNil,
			wantErr: assert.NoError,
		},
		{
			name:         "error/interface `CreateTempFS`",
			randomReader: nil,
			args: args{
				wfs: func(test *testing.T) writablefs.WritableFS {
					wfs := fsutilsmocks.NewCreateTempFS(test)
					wfs.EXPECT().
						CreateTemp("base-dir", "path-pattern-*").
						Return(nil, iotest.ErrTimeout)

					return wfs
				},
				baseDir:     "base-dir",
				pathPattern: "path-pattern-*",
			},
			want: assert.Nil,
			wantErr: func(test assert.TestingT, err error, _ ...interface{}) bool {
				return assert.ErrorIs(test, err, iotest.ErrTimeout)
			},
		},
		{
			name:         "error/invalid base directory",
			randomReader: nil,
			args: args{
				wfs: func(test *testing.T) writablefs.WritableFS {
					return writablefsmocks.NewWritableFS(test)
				},
				baseDir:     "/base-dir",
				pathPattern: "path-pattern-*",
			},
			want: assert.Nil,
			wantErr: func(test assert.TestingT, err error, _ ...interface{}) bool {
				wantErr := &fs.PathError{
					Op:   "createtemp",
					Path: "/base-dir/path-pattern-*",
					Err:  fs.ErrInvalid,
				}
				return assert.Equal(test, wantErr, err)
			},
		},
		{
			name:         "error/random reader",
			randomReader: iotest.ErrReader(iotest.ErrTimeout),
			args: args{
				wfs: func(test *testing.T) writablefs.WritableFS {
					return writablefsmocks.NewWritableFS(test)
				},
				baseDir:     "base-dir",
				pathPattern: "path-pattern-*",
			},
			want: assert.Nil,
			wantErr: func(test assert.TestingT, err error, _ ...interface{}) bool {
				wantErr := &fs.PathError{
					Op:   "createtemp",
					Path: "base-dir/path-pattern-*",
					Err:  iotest.ErrTimeout,
				}
				return assert.Equal(test, wantErr, err)
			},
		},
		{
			name: "error/not error `fs.ErrExist`",
			randomReader: func() io.Reader {
				randomBuffer := make([]byte, maxRandomSuffixValue.BitLen()/8)
				_, err := mathrand.New(mathrand.NewSource(1)).Read(randomBuffer)
				require.NoError(test, err)

				return bytes.NewBuffer(randomBuffer)
			}(),
			args: args{
				wfs: func(test *testing.T) writablefs.WritableFS {
					wfs := writablefsmocks.NewWritableFS(test)
					wfs.EXPECT().
						CreateExcl("base-dir/path-pattern-5980212987775051087").
						Return(nil, iotest.ErrTimeout)

					return wfs
				},
				baseDir:     "base-dir",
				pathPattern: "path-pattern-*",
			},
			want: assert.Nil,
			wantErr: func(test assert.TestingT, err error, _ ...interface{}) bool {
				return assert.ErrorIs(test, err, iotest.ErrTimeout)
			},
		},
		{
			name: "error" +
				"/error `fs.ErrExist`" +
				"/the count of tries is greater than the maximum",
			randomReader: func() io.Reader {
				randomBuffer := make([]byte, maxRandomSuffixValue.BitLen()/8)
				_, err := mathrand.New(mathrand.NewSource(1)).Read(randomBuffer)
				require.NoError(test, err)

				repeatedRandomBuffer := bytes.Repeat(randomBuffer, maxCountOfTempTries+1)
				return bytes.NewBuffer(repeatedRandomBuffer)
			}(),
			args: args{
				wfs: func(test *testing.T) writablefs.WritableFS {
					wfs := writablefsmocks.NewWritableFS(test)
					wfs.EXPECT().
						CreateExcl("base-dir/path-pattern-5980212987775051087").
						Return(nil, fs.ErrExist).
						Times(maxCountOfTempTries + 1)

					return wfs
				},
				baseDir:     "base-dir",
				pathPattern: "path-pattern-*",
			},
			want: assert.Nil,
			wantErr: func(test assert.TestingT, err error, _ ...interface{}) bool {
				return assert.ErrorIs(test, err, fs.ErrExist)
			},
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			previousRandomReader := cryptorand.Reader
			cryptorand.Reader = data.randomReader
			defer func() {
				cryptorand.Reader = previousRandomReader
			}()

			got, err := CreateTemp(
				data.args.wfs(test),
				data.args.baseDir,
				data.args.pathPattern,
			)

			data.want(test, got)
			data.wantErr(test, err)
		})
	}
}

func TestCreateTemp_withDirFS(test *testing.T) {
	type args struct {
		baseDir     string
		pathPattern string
	}

	for _, data := range []struct {
		name         string
		randomReader io.Reader
		preparation  func(test *testing.T, tempDir string)
		args         args
		wantTempDir  func(test *testing.T, tempDir string)
		want         assert.ValueAssertionFunc
		wantErr      func(test *testing.T, err error)
	}{
		{
			name: "success/non-existent path",
			randomReader: func() io.Reader {
				randomBuffer := make([]byte, maxRandomSuffixValue.BitLen()/8)
				_, err := mathrand.New(mathrand.NewSource(1)).Read(randomBuffer)
				require.NoError(test, err)

				return bytes.NewBuffer(randomBuffer)
			}(),
			preparation: func(test *testing.T, tempDir string) {
				path := filepath.Join(tempDir, "base-dir")
				err := os.Mkdir(path, tempDirPermission)
				require.NoError(test, err)
			},
			args: args{
				baseDir:     "base-dir",
				pathPattern: "path-pattern-*",
			},
			wantTempDir: func(test *testing.T, tempDir string) {
				wantPath := "base-dir/path-pattern-5980212987775051087"

				path := filepath.Join(tempDir, wantPath)
				if !assert.FileExists(test, path) {
					return
				}

				stat, err := os.Stat(path)
				if !assert.NoError(test, err) {
					return
				}

				assert.Equal(test, fs.FileMode(0664), stat.Mode())
			},
			want: assert.NotNil,
			wantErr: func(test *testing.T, err error) {
				assert.NoError(test, err)
			},
		},
		{
			name: "success/existent path/the count of tries is less than the maximum",
			randomReader: func() io.Reader {
				randomReader := mathrand.New(mathrand.NewSource(1))

				randomBufferOne := make([]byte, maxRandomSuffixValue.BitLen()/8)
				_, err := randomReader.Read(randomBufferOne)
				require.NoError(test, err)

				randomBufferTwo := make([]byte, maxRandomSuffixValue.BitLen()/8)
				_, err = randomReader.Read(randomBufferTwo)
				require.NoError(test, err)

				joinedRandomBuffers :=
					bytes.Join([][]byte{randomBufferOne, randomBufferTwo}, nil)
				return bytes.NewBuffer(joinedRandomBuffers)
			}(),
			preparation: func(test *testing.T, tempDir string) {
				dirPath := filepath.Join(tempDir, "base-dir")
				err := os.Mkdir(dirPath, 0700)
				require.NoError(test, err)

				filePath := filepath.Join(dirPath, "path-pattern-5980212987775051087")
				err = os.WriteFile(filePath, []byte("content"), 0666)
				require.NoError(test, err)
			},
			args: args{
				baseDir:     "base-dir",
				pathPattern: "path-pattern-*",
			},
			wantTempDir: func(test *testing.T, tempDir string) {
				wantPath := "base-dir/path-pattern-1603104512986455410"

				path := filepath.Join(tempDir, wantPath)
				if !assert.FileExists(test, path) {
					return
				}

				stat, err := os.Stat(path)
				if !assert.NoError(test, err) {
					return
				}

				assert.Equal(test, fs.FileMode(0664), stat.Mode())
			},
			want: assert.NotNil,
			wantErr: func(test *testing.T, err error) {
				assert.NoError(test, err)
			},
		},
		{
			name: "error/non-existent base directory",
			randomReader: func() io.Reader {
				randomBuffer := make([]byte, maxRandomSuffixValue.BitLen()/8)
				_, err := mathrand.New(mathrand.NewSource(1)).Read(randomBuffer)
				require.NoError(test, err)

				return bytes.NewBuffer(randomBuffer)
			}(),
			preparation: func(_ *testing.T, _ string) {},
			args: args{
				baseDir:     "base-dir",
				pathPattern: "path-pattern-*",
			},
			wantTempDir: func(_ *testing.T, _ string) {},
			want:        assert.Nil,
			wantErr: func(test *testing.T, err error) {
				assert.ErrorIs(test, err, fs.ErrNotExist)
			},
		},
		{
			name: "error/existent path/the count of tries is greater than the maximum",
			randomReader: func() io.Reader {
				randomBuffer := make([]byte, maxRandomSuffixValue.BitLen()/8)
				_, err := mathrand.New(mathrand.NewSource(1)).Read(randomBuffer)
				require.NoError(test, err)

				repeatedRandomBuffer := bytes.Repeat(randomBuffer, maxCountOfTempTries+1)
				return bytes.NewBuffer(repeatedRandomBuffer)
			}(),
			preparation: func(test *testing.T, tempDir string) {
				dirPath := filepath.Join(tempDir, "base-dir")
				err := os.Mkdir(dirPath, 0700)
				require.NoError(test, err)

				filePath := filepath.Join(dirPath, "path-pattern-5980212987775051087")
				err = os.WriteFile(filePath, []byte("content"), 0666)
				require.NoError(test, err)
			},
			args: args{
				baseDir:     "base-dir",
				pathPattern: "path-pattern-*",
			},
			wantTempDir: func(_ *testing.T, _ string) {},
			want:        assert.Nil,
			wantErr: func(test *testing.T, err error) {
				assert.ErrorIs(test, err, fs.ErrExist)
			},
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			previousRandomReader := cryptorand.Reader
			cryptorand.Reader = data.randomReader
			defer func() {
				cryptorand.Reader = previousRandomReader
			}()

			tempDir, err := os.MkdirTemp("", "test-*")
			require.NoError(test, err)
			defer os.RemoveAll(tempDir)

			data.preparation(test, tempDir)

			dfs := writablefs.NewDirFS(tempDir)
			got, err := CreateTemp(dfs, data.args.baseDir, data.args.pathPattern)

			data.wantTempDir(test, tempDir)
			data.want(test, got)
			data.wantErr(test, err)
		})
	}
}
