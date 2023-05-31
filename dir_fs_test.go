package writablefs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDirFS(test *testing.T) {
	const baseDir = "base-dir"
	got := NewDirFS(baseDir)

	assert.Equal(test, os.DirFS(baseDir), got.innerDirFS)
	assert.Equal(test, baseDir, got.baseDir)
}
