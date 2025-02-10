// Code generated by mockery v2.51.1. DO NOT EDIT.

package adaptermocks

import (
	context "context"

	entity "github.com/pangolin-do-golang/thumb-processor-api/internal/core/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

// IThumbQueueAdapter is an autogenerated mock type for the IThumbQueueAdapter type
type IThumbQueueAdapter struct {
	mock.Mock
}

// SendEvent provides a mock function with given fields: ctx, process
func (_m *IThumbQueueAdapter) SendEvent(ctx context.Context, process *entity.ThumbProcess) error {
	ret := _m.Called(ctx, process)

	if len(ret) == 0 {
		panic("no return value specified for SendEvent")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.ThumbProcess) error); ok {
		r0 = rf(ctx, process)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIThumbQueueAdapter creates a new instance of IThumbQueueAdapter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIThumbQueueAdapter(t interface {
	mock.TestingT
	Cleanup(func())
}) *IThumbQueueAdapter {
	mock := &IThumbQueueAdapter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
