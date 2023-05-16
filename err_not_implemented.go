package writablefs

import (
	"errors"
	"fmt"
	"runtime"
)

var (
	ErrNotImplemented = errors.New("not implemented")
)

func MakeSpecificErrNotImplemented() error {
	callerName, ok := getCallerName(1)
	if !ok {
		callerName = "unknown"
	}

	return fmt.Errorf("function `%s()` is %w", callerName, ErrNotImplemented)
}

func getCallerName(levelsToSkip int) (name string, ok bool) {
	programCounter, _, _, ok := runtime.Caller(levelsToSkip + 1)
	if !ok {
		return "", false
	}

	function := runtime.FuncForPC(programCounter)
	if function == nil {
		return "", false
	}

	return function.Name(), true
}
