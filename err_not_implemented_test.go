package writablefs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeSpecificErrNotImplemented(test *testing.T) {
	err := MakeSpecificErrNotImplemented()

	const expectedCallerName = "TestMakeSpecificErrNotImplemented"
	assertSpecificErrNotImplemented(test, err, expectedCallerName)
}

func assertSpecificErrNotImplemented(
	test *testing.T,
	err error,
	expectedCallerName string,
) {
	const expectedPackageName = "github.com/thewizardplusplus/go-writable-fs"
	assert.EqualError(test, err, fmt.Sprintf(
		"function `%s.%s()` is not implemented",
		expectedPackageName,
		expectedCallerName,
	))

	assert.ErrorIs(test, err, ErrNotImplemented)
}
