package writablefs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDummyFS_interface(test *testing.T) {
	assert.Implements(test, (*WritableFS)(nil), DummyFS{})
}

func TestDummyFS_Open(test *testing.T) {
	file, err := DummyFS{}.Open("path")

	assert.Nil(test, file)
	assertSpecificErrNotImplemented(test, err, "DummyFS.Open")
}

func TestDummyFS_Mkdir(test *testing.T) {
	err := DummyFS{}.Mkdir("path", 1755)

	assertSpecificErrNotImplemented(test, err, "DummyFS.Mkdir")
}

func TestDummyFS_Create(test *testing.T) {
	writableFile, err := DummyFS{}.Create("path")

	assert.Nil(test, writableFile)
	assertSpecificErrNotImplemented(test, err, "DummyFS.Create")
}

func TestDummyFS_CreateExcl(test *testing.T) {
	writableFile, err := DummyFS{}.CreateExcl("path")

	assert.Nil(test, writableFile)
	assertSpecificErrNotImplemented(test, err, "DummyFS.CreateExcl")
}

func TestDummyFS_Rename(test *testing.T) {
	err := DummyFS{}.Rename("old-path", "new-path")

	assertSpecificErrNotImplemented(test, err, "DummyFS.Rename")
}

func TestDummyFS_Remove(test *testing.T) {
	err := DummyFS{}.Remove("path")

	assertSpecificErrNotImplemented(test, err, "DummyFS.Remove")
}
