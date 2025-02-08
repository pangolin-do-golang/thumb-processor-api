// Code generated by mockery v2.51.1. DO NOT EDIT.

package adaptermocks

import (
	entity "github.com/pangolin-do-golang/thumb-processor-api/internal/core/domain/entity"
	mock "github.com/stretchr/testify/mock"
)

// IThumbRepositoryAdapter is an autogenerated mock type for the IThumbRepositoryAdapter type
type IThumbRepositoryAdapter struct {
	mock.Mock
}

// Create provides a mock function with given fields: process
func (_m *IThumbRepositoryAdapter) Create(process *entity.ThumbProcess) error {
	ret := _m.Called(process)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.ThumbProcess) error); ok {
		r0 = rf(process)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List provides a mock function with no fields
func (_m *IThumbRepositoryAdapter) List() *[]entity.ThumbProcess {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 *[]entity.ThumbProcess
	if rf, ok := ret.Get(0).(func() *[]entity.ThumbProcess); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entity.ThumbProcess)
		}
	}

	return r0
}

// Update provides a mock function with given fields: process
func (_m *IThumbRepositoryAdapter) Update(process *entity.ThumbProcess) (*entity.ThumbProcess, error) {
	ret := _m.Called(process)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *entity.ThumbProcess
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.ThumbProcess) (*entity.ThumbProcess, error)); ok {
		return rf(process)
	}
	if rf, ok := ret.Get(0).(func(*entity.ThumbProcess) *entity.ThumbProcess); ok {
		r0 = rf(process)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.ThumbProcess)
		}
	}

	if rf, ok := ret.Get(1).(func(*entity.ThumbProcess) error); ok {
		r1 = rf(process)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIThumbRepositoryAdapter creates a new instance of IThumbRepositoryAdapter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIThumbRepositoryAdapter(t interface {
	mock.TestingT
	Cleanup(func())
}) *IThumbRepositoryAdapter {
	mock := &IThumbRepositoryAdapter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
