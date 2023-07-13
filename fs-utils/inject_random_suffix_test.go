package fsutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_injectRandomSuffix(test *testing.T) {
	type args struct {
		pattern      string
		randomSuffix string
	}

	for _, data := range []struct {
		name string
		args args
		want string
	}{
		{
			name: "without the placeholder",
			args: args{
				pattern:      "test",
				randomSuffix: "23",
			},
			want: "test23",
		},
		{
			name: "with the placeholder at the beginning",
			args: args{
				pattern:      "*test",
				randomSuffix: "23",
			},
			want: "23test",
		},
		{
			name: "with the placeholder in the middle",
			args: args{
				pattern:      "te*st",
				randomSuffix: "23",
			},
			want: "te23st",
		},
		{
			name: "with the placeholder at the end",
			args: args{
				pattern:      "test*",
				randomSuffix: "23",
			},
			want: "test23",
		},
		{
			name: "with multiple placeholders",
			args: args{
				pattern:      "te*****st",
				randomSuffix: "23",
			},
			want: "te****23st",
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			got := injectRandomSuffix(data.args.pattern, data.args.randomSuffix)

			assert.Equal(test, data.want, got)
		})
	}
}
