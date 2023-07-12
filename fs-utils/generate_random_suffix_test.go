package fsutils

import (
	"bytes"
	cryptorand "crypto/rand"
	"io"
	mathrand "math/rand"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_generateRandomSuffix(test *testing.T) {
	for _, data := range []struct {
		name         string
		randomReader io.Reader
		want         string
		wantErr      assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			randomReader: func() io.Reader {
				randomBuffer := make([]byte, maxRandomSuffixValue.BitLen()/8)
				_, err := mathrand.New(mathrand.NewSource(1)).Read(randomBuffer)
				require.NoError(test, err)

				return bytes.NewBuffer(randomBuffer)
			}(),
			want:    "5980212987775051087",
			wantErr: assert.NoError,
		},
		{
			name:         "error",
			randomReader: iotest.ErrReader(iotest.ErrTimeout),
			want:         "",
			wantErr: func(
				test assert.TestingT,
				err error,
				msgAndArgs ...interface{},
			) bool {
				return assert.ErrorIs(test, err, iotest.ErrTimeout)
			},
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			previousRandomReader := cryptorand.Reader
			cryptorand.Reader = data.randomReader
			defer func() {
				cryptorand.Reader = previousRandomReader
			}()

			got, err := generateRandomSuffix()

			assert.Equal(test, data.want, got)
			data.wantErr(test, err)
		})
	}
}
