// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	fs "io/fs"

	mock "github.com/stretchr/testify/mock"

	writablefs "github.com/thewizardplusplus/go-writable-fs"
)

// RemoveAllFS is an autogenerated mock type for the RemoveAllFS type
type RemoveAllFS struct {
	mock.Mock
}

type RemoveAllFS_Expecter struct {
	mock *mock.Mock
}

func (_m *RemoveAllFS) EXPECT() *RemoveAllFS_Expecter {
	return &RemoveAllFS_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: path
func (_m *RemoveAllFS) Create(path string) (writablefs.WritableFile, error) {
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

// RemoveAllFS_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type RemoveAllFS_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - path string
func (_e *RemoveAllFS_Expecter) Create(path interface{}) *RemoveAllFS_Create_Call {
	return &RemoveAllFS_Create_Call{Call: _e.mock.On("Create", path)}
}

func (_c *RemoveAllFS_Create_Call) Run(run func(path string)) *RemoveAllFS_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *RemoveAllFS_Create_Call) Return(_a0 writablefs.WritableFile, _a1 error) *RemoveAllFS_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RemoveAllFS_Create_Call) RunAndReturn(run func(string) (writablefs.WritableFile, error)) *RemoveAllFS_Create_Call {
	_c.Call.Return(run)
	return _c
}

// CreateExcl provides a mock function with given fields: path
func (_m *RemoveAllFS) CreateExcl(path string) (writablefs.WritableFile, error) {
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

// RemoveAllFS_CreateExcl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateExcl'
type RemoveAllFS_CreateExcl_Call struct {
	*mock.Call
}

// CreateExcl is a helper method to define mock.On call
//   - path string
func (_e *RemoveAllFS_Expecter) CreateExcl(path interface{}) *RemoveAllFS_CreateExcl_Call {
	return &RemoveAllFS_CreateExcl_Call{Call: _e.mock.On("CreateExcl", path)}
}

func (_c *RemoveAllFS_CreateExcl_Call) Run(run func(path string)) *RemoveAllFS_CreateExcl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *RemoveAllFS_CreateExcl_Call) Return(_a0 writablefs.WritableFile, _a1 error) *RemoveAllFS_CreateExcl_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RemoveAllFS_CreateExcl_Call) RunAndReturn(run func(string) (writablefs.WritableFile, error)) *RemoveAllFS_CreateExcl_Call {
	_c.Call.Return(run)
	return _c
}

// Mkdir provides a mock function with given fields: path, permissions
func (_m *RemoveAllFS) Mkdir(path string, permissions fs.FileMode) error {
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

// RemoveAllFS_Mkdir_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Mkdir'
type RemoveAllFS_Mkdir_Call struct {
	*mock.Call
}

// Mkdir is a helper method to define mock.On call
//   - path string
//   - permissions fs.FileMode
func (_e *RemoveAllFS_Expecter) Mkdir(path interface{}, permissions interface{}) *RemoveAllFS_Mkdir_Call {
	return &RemoveAllFS_Mkdir_Call{Call: _e.mock.On("Mkdir", path, permissions)}
}

func (_c *RemoveAllFS_Mkdir_Call) Run(run func(path string, permissions fs.FileMode)) *RemoveAllFS_Mkdir_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(fs.FileMode))
	})
	return _c
}

func (_c *RemoveAllFS_Mkdir_Call) Return(_a0 error) *RemoveAllFS_Mkdir_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RemoveAllFS_Mkdir_Call) RunAndReturn(run func(string, fs.FileMode) error) *RemoveAllFS_Mkdir_Call {
	_c.Call.Return(run)
	return _c
}

// Open provides a mock function with given fields: name
func (_m *RemoveAllFS) Open(name string) (fs.File, error) {
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

// RemoveAllFS_Open_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Open'
type RemoveAllFS_Open_Call struct {
	*mock.Call
}

// Open is a helper method to define mock.On call
//   - name string
func (_e *RemoveAllFS_Expecter) Open(name interface{}) *RemoveAllFS_Open_Call {
	return &RemoveAllFS_Open_Call{Call: _e.mock.On("Open", name)}
}

func (_c *RemoveAllFS_Open_Call) Run(run func(name string)) *RemoveAllFS_Open_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *RemoveAllFS_Open_Call) Return(_a0 fs.File, _a1 error) *RemoveAllFS_Open_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RemoveAllFS_Open_Call) RunAndReturn(run func(string) (fs.File, error)) *RemoveAllFS_Open_Call {
	_c.Call.Return(run)
	return _c
}

// Remove provides a mock function with given fields: path
func (_m *RemoveAllFS) Remove(path string) error {
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

// RemoveAllFS_Remove_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Remove'
type RemoveAllFS_Remove_Call struct {
	*mock.Call
}

// Remove is a helper method to define mock.On call
//   - path string
func (_e *RemoveAllFS_Expecter) Remove(path interface{}) *RemoveAllFS_Remove_Call {
	return &RemoveAllFS_Remove_Call{Call: _e.mock.On("Remove", path)}
}

func (_c *RemoveAllFS_Remove_Call) Run(run func(path string)) *RemoveAllFS_Remove_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *RemoveAllFS_Remove_Call) Return(_a0 error) *RemoveAllFS_Remove_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RemoveAllFS_Remove_Call) RunAndReturn(run func(string) error) *RemoveAllFS_Remove_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveAll provides a mock function with given fields: path
func (_m *RemoveAllFS) RemoveAll(path string) error {
	ret := _m.Called(path)

	if len(ret) == 0 {
		panic("no return value specified for RemoveAll")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveAllFS_RemoveAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveAll'
type RemoveAllFS_RemoveAll_Call struct {
	*mock.Call
}

// RemoveAll is a helper method to define mock.On call
//   - path string
func (_e *RemoveAllFS_Expecter) RemoveAll(path interface{}) *RemoveAllFS_RemoveAll_Call {
	return &RemoveAllFS_RemoveAll_Call{Call: _e.mock.On("RemoveAll", path)}
}

func (_c *RemoveAllFS_RemoveAll_Call) Run(run func(path string)) *RemoveAllFS_RemoveAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *RemoveAllFS_RemoveAll_Call) Return(_a0 error) *RemoveAllFS_RemoveAll_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RemoveAllFS_RemoveAll_Call) RunAndReturn(run func(string) error) *RemoveAllFS_RemoveAll_Call {
	_c.Call.Return(run)
	return _c
}

// Rename provides a mock function with given fields: oldPath, newPath
func (_m *RemoveAllFS) Rename(oldPath string, newPath string) error {
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

// RemoveAllFS_Rename_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Rename'
type RemoveAllFS_Rename_Call struct {
	*mock.Call
}

// Rename is a helper method to define mock.On call
//   - oldPath string
//   - newPath string
func (_e *RemoveAllFS_Expecter) Rename(oldPath interface{}, newPath interface{}) *RemoveAllFS_Rename_Call {
	return &RemoveAllFS_Rename_Call{Call: _e.mock.On("Rename", oldPath, newPath)}
}

func (_c *RemoveAllFS_Rename_Call) Run(run func(oldPath string, newPath string)) *RemoveAllFS_Rename_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *RemoveAllFS_Rename_Call) Return(_a0 error) *RemoveAllFS_Rename_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RemoveAllFS_Rename_Call) RunAndReturn(run func(string, string) error) *RemoveAllFS_Rename_Call {
	_c.Call.Return(run)
	return _c
}

// NewRemoveAllFS creates a new instance of RemoveAllFS. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRemoveAllFS(t interface {
	mock.TestingT
	Cleanup(func())
}) *RemoveAllFS {
	mock := &RemoveAllFS{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
