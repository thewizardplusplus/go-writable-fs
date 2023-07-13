package fsutils

import (
	"strings"
)

const (
	randomSuffixPlaceholder = '*'
)

func injectRandomSuffix(pattern string, randomSuffix string) string {
	randomSuffixPosition := strings.LastIndexByte(pattern, randomSuffixPlaceholder)
	if randomSuffixPosition == -1 {
		return pattern + randomSuffix
	}

	return pattern[:randomSuffixPosition] +
		randomSuffix +
		pattern[randomSuffixPosition+1:]
}
