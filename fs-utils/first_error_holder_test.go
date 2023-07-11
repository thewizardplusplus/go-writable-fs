package fsutils

import (
	"errors"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/assert"
)

func Test_firstErrorHolder_firstErr(test *testing.T) {
	firstErrorHolder := firstErrorHolder{
		innerErr: iotest.ErrTimeout,
	}
	err := firstErrorHolder.firstErr()

	assert.ErrorIs(test, err, iotest.ErrTimeout)
}

func Test_firstErrorHolder_updateErr(test *testing.T) {
	type fields struct {
		innerErr error
	}
	type args struct {
		anotherErr error
	}

	for _, data := range []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "first call",
			fields: fields{
				innerErr: nil,
			},
			args: args{
				anotherErr: iotest.ErrTimeout,
			},
			wantErr: func(
				test assert.TestingT,
				err error,
				msgAndArgs ...interface{},
			) bool {
				return assert.ErrorIs(test, err, iotest.ErrTimeout)
			},
		},
		{
			name: "second call",
			fields: fields{
				innerErr: iotest.ErrTimeout,
			},
			args: args{
				anotherErr: errors.New("dummy"),
			},
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
			firstErrorHolder := &firstErrorHolder{
				innerErr: data.fields.innerErr,
			}
			firstErrorHolder.updateErr(data.args.anotherErr)

			data.wantErr(test, firstErrorHolder.innerErr)
		})
	}
}
