package fsutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countPathElements(test *testing.T) {
	type args struct {
		path string
	}

	for _, data := range []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty path",
			args: args{
				path: "",
			},
			want: 1,
		},
		{
			name: "single path element",
			args: args{
				path: "path",
			},
			want: 1,
		},
		{
			name: "multiple path elements",
			args: args{
				path: "path-one/path-two/path-three",
			},
			want: 3,
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			got := countPathElements(data.args.path)

			assert.Equal(test, data.want, got)
		})
	}
}
