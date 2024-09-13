// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	context "context"
	user "golang-training/logic/user"

	mock "github.com/stretchr/testify/mock"
)

// Users is an autogenerated mock type for the Users type
type Users struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, request
func (_m *Users) Create(ctx context.Context, request *user.CreateUserRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *user.CreateUserRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, uuid
func (_m *Users) Get(ctx context.Context, uuid string) (*user.User, error) {
	ret := _m.Called(ctx, uuid)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*user.User, error)); ok {
		return rf(ctx, uuid)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *user.User); ok {
		r0 = rf(ctx, uuid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUsers creates a new instance of Users. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUsers(t interface {
	mock.TestingT
	Cleanup(func())
}) *Users {
	mock := &Users{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
