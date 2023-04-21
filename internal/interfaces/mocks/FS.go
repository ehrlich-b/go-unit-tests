// Code generated by mockery v2.26.0. DO NOT EDIT.

package mocks

import (
	fs "io/fs"

	interfaces "github.com/ehrlich-b/go-unit-tests/internal/interfaces"
	mock "github.com/stretchr/testify/mock"
)

// FS is an autogenerated mock type for the FS type
type FS struct {
	mock.Mock
}

// OpenFile provides a mock function with given fields: name, flag, perm
func (_m *FS) OpenFile(name string, flag int, perm fs.FileMode) (interfaces.WriteCloser, error) {
	ret := _m.Called(name, flag, perm)

	var r0 interfaces.WriteCloser
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int, fs.FileMode) (interfaces.WriteCloser, error)); ok {
		return rf(name, flag, perm)
	}
	if rf, ok := ret.Get(0).(func(string, int, fs.FileMode) interfaces.WriteCloser); ok {
		r0 = rf(name, flag, perm)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.WriteCloser)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int, fs.FileMode) error); ok {
		r1 = rf(name, flag, perm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewFS interface {
	mock.TestingT
	Cleanup(func())
}

// NewFS creates a new instance of FS. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFS(t mockConstructorTestingTNewFS) *FS {
	mock := &FS{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}