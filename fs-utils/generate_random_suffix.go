package fsutils

import (
	"crypto/rand"
	"math"
	"math/big"
)

var (
	maxRandomSuffixValue = big.NewInt(0).SetUint64(math.MaxUint64)
)

func generateRandomSuffix() (string, error) {
	randomSuffixValue, err := rand.Int(rand.Reader, maxRandomSuffixValue)
	if err != nil {
		return "", err
	}

	return randomSuffixValue.String(), nil
}
