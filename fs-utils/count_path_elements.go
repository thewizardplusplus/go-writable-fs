package fsutils

import (
	"strings"
)

func countPathElements(path string) int {
	return strings.Count(path, "/") + 1
}
