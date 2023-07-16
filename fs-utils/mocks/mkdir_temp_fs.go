// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	fs "io/fs"

	mock "github.com/stretchr/testify/mock"

	writablefs "github.com/thewizardplusplus/go-writable-fs"
)

// MkdirTempFS is an autogenerated mock type for the MkdirTempFS type
type MkdirTempFS struct {
	mock.Mock
}

type MkdirTempFS_Expecter struct {
	mock *mock.Mock
}

func (_m *MkdirTempFS) EXPECT() *MkdirTempFS_Expecter {
	return &MkdirTempFS_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: path
func (_m *MkdirTempFS) Create(path string) (writablefs.WritableFile, error) {
	ret := _m.Called(path)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 writablefs.WritableFile
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (writablefs.WritableFile, error)); ok {
		return rf(path)
	}
	if rf, ok := ret.Get(0).(func(string) writablefs.WritableFile); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(writablefs.WritableFile)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MkdirTempFS_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MkdirTempFS_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - path string
func (_e *MkdirTempFS_Expecter) Create(path interface{}) *MkdirTempFS_Create_Call {
	return &MkdirTempFS_Create_Call{Call: _e.mock.On("Create", path)}
}

func (_c *MkdirTempFS_Create_Call) Run(run func(path string)) *MkdirTempFS_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MkdirTempFS_Create_Call) Return(_a0 writablefs.WritableFile, _a1 error) *MkdirTempFS_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MkdirTempFS_Create_Call) RunAndReturn(run func(string) (writablefs.WritableFile, error)) *MkdirTempFS_Create_Call {
	_c.Call.Return(run)
	return _c
}

// CreateExcl provides a mock function with given fields: path
func (_m *MkdirTempFS) CreateExcl(path string) (writablefs.WritableFile, error) {
	ret := _m.Called(path)

	if len(ret) == 0 {
		panic("no return value specified for CreateExcl")
	}

	var r0 writablefs.WritableFile
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (writablefs.WritableFile, error)); ok {
		return rf(path)
	}
	if rf, ok := ret.Get(0).(func(string) writablefs.WritableFile); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(writablefs.WritableFile)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MkdirTempFS_CreateExcl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateExcl'
type MkdirTempFS_CreateExcl_Call struct {
	*mock.Call
}

// CreateExcl is a helper method to define mock.On call
//   - path string
func (_e *MkdirTempFS_Expecter) CreateExcl(path interface{}) *MkdirTempFS_CreateExcl_Call {
	return &MkdirTempFS_CreateExcl_Call{Call: _e.mock.On("CreateExcl", path)}
}

func (_c *MkdirTempFS_CreateExcl_Call) Run(run func(path string)) *MkdirTempFS_CreateExcl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MkdirTempFS_CreateExcl_Call) Return(_a0 writablefs.WritableFile, _a1 error) *MkdirTempFS_CreateExcl_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MkdirTempFS_CreateExcl_Call) RunAndReturn(run func(string) (writablefs.WritableFile, error)) *MkdirTempFS_CreateExcl_Call {
	_c.Call.Return(run)
	return _c
}

// Mkdir provides a mock function with given fields: path, permissions
func (_m *MkdirTempFS) Mkdir(path string, permissions fs.FileMode) error {
	ret := _m.Called(path, permissions)

	if len(ret) == 0 {
		panic("no return value specified for Mkdir")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, fs.FileMode) error); ok {
		r0 = rf(path, permissions)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MkdirTempFS_Mkdir_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Mkdir'
type MkdirTempFS_Mkdir_Call struct {
	*mock.Call
}

// Mkdir is a helper method to define mock.On call
//   - path string
//   - permissions fs.FileMode
func (_e *MkdirTempFS_Expecter) Mkdir(path interface{}, permissions interface{}) *MkdirTempFS_Mkdir_Call {
	return &MkdirTempFS_Mkdir_Call{Call: _e.mock.On("Mkdir", path, permissions)}
}

func (_c *MkdirTempFS_Mkdir_Call) Run(run func(path string, permissions fs.FileMode)) *MkdirTempFS_Mkdir_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(fs.FileMode))
	})
	return _c
}

func (_c *MkdirTempFS_Mkdir_Call) Return(_a0 error) *MkdirTempFS_Mkdir_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MkdirTempFS_Mkdir_Call) RunAndReturn(run func(string, fs.FileMode) error) *MkdirTempFS_Mkdir_Call {
	_c.Call.Return(run)
	return _c
}

// MkdirTemp provides a mock function with given fields: baseDir, pathPattern
func (_m *MkdirTempFS) MkdirTemp(baseDir string, pathPattern string) (string, error) {
	ret := _m.Called(baseDir, pathPattern)

	if len(ret) == 0 {
		panic("no return value specified for MkdirTemp")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (string, error)); ok {
		return rf(baseDir, pathPattern)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(baseDir, pathPattern)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(baseDir, pathPattern)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MkdirTempFS_MkdirTemp_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MkdirTemp'
type MkdirTempFS_MkdirTemp_Call struct {
	*mock.Call
}

// MkdirTemp is a helper method to define mock.On call
//   - baseDir string
//   - pathPattern string
func (_e *MkdirTempFS_Expecter) MkdirTemp(baseDir interface{}, pathPattern interface{}) *MkdirTempFS_MkdirTemp_Call {
	return &MkdirTempFS_MkdirTemp_Call{Call: _e.mock.On("MkdirTemp", baseDir, pathPattern)}
}

func (_c *MkdirTempFS_MkdirTemp_Call) Run(run func(baseDir string, pathPattern string)) *MkdirTempFS_MkdirTemp_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MkdirTempFS_MkdirTemp_Call) Return(_a0 string, _a1 error) *MkdirTempFS_MkdirTemp_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MkdirTempFS_MkdirTemp_Call) RunAndReturn(run func(string, string) (string, error)) *MkdirTempFS_MkdirTemp_Call {
	_c.Call.Return(run)
	return _c
}

// Open provides a mock function with given fields: name
func (_m *MkdirTempFS) Open(name string) (fs.File, error) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for Open")
	}

	var r0 fs.File
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (fs.File, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) fs.File); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(fs.File)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MkdirTempFS_Open_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Open'
type MkdirTempFS_Open_Call struct {
	*mock.Call
}

// Open is a helper method to define mock.On call
//   - name string
func (_e *MkdirTempFS_Expecter) Open(name interface{}) *MkdirTempFS_Open_Call {
	return &MkdirTempFS_Open_Call{Call: _e.mock.On("Open", name)}
}

func (_c *MkdirTempFS_Open_Call) Run(run func(name string)) *MkdirTempFS_Open_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MkdirTempFS_Open_Call) Return(_a0 fs.File, _a1 error) *MkdirTempFS_Open_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MkdirTempFS_Open_Call) RunAndReturn(run func(string) (fs.File, error)) *MkdirTempFS_Open_Call {
	_c.Call.Return(run)
	return _c
}

// Remove provides a mock function with given fields: path
func (_m *MkdirTempFS) Remove(path string) error {
	ret := _m.Called(path)

	if len(ret) == 0 {
		panic("no return value specified for Remove")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MkdirTempFS_Remove_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Remove'
type MkdirTempFS_Remove_Call struct {
	*mock.Call
}

// Remove is a helper method to define mock.On call
//   - path string
func (_e *MkdirTempFS_Expecter) Remove(path interface{}) *MkdirTempFS_Remove_Call {
	return &MkdirTempFS_Remove_Call{Call: _e.mock.On("Remove", path)}
}

func (_c *MkdirTempFS_Remove_Call) Run(run func(path string)) *MkdirTempFS_Remove_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MkdirTempFS_Remove_Call) Return(_a0 error) *MkdirTempFS_Remove_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MkdirTempFS_Remove_Call) RunAndReturn(run func(string) error) *MkdirTempFS_Remove_Call {
	_c.Call.Return(run)
	return _c
}

// Rename provides a mock function with given fields: oldPath, newPath
func (_m *MkdirTempFS) Rename(oldPath string, newPath string) error {
	ret := _m.Called(oldPath, newPath)

	if len(ret) == 0 {
		panic("no return value specified for Rename")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(oldPath, newPath)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MkdirTempFS_Rename_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Rename'
type MkdirTempFS_Rename_Call struct {
	*mock.Call
}

// Rename is a helper method to define mock.On call
//   - oldPath string
//   - newPath string
func (_e *MkdirTempFS_Expecter) Rename(oldPath interface{}, newPath interface{}) *MkdirTempFS_Rename_Call {
	return &MkdirTempFS_Rename_Call{Call: _e.mock.On("Rename", oldPath, newPath)}
}

func (_c *MkdirTempFS_Rename_Call) Run(run func(oldPath string, newPath string)) *MkdirTempFS_Rename_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MkdirTempFS_Rename_Call) Return(_a0 error) *MkdirTempFS_Rename_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MkdirTempFS_Rename_Call) RunAndReturn(run func(string, string) error) *MkdirTempFS_Rename_Call {
	_c.Call.Return(run)
	return _c
}

// NewMkdirTempFS creates a new instance of MkdirTempFS. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMkdirTempFS(t interface {
	mock.TestingT
	Cleanup(func())
}) *MkdirTempFS {
	mock := &MkdirTempFS{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
