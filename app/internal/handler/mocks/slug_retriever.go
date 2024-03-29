// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// SlugRetriever is an autogenerated mock type for the slugRetriever type
type SlugRetriever struct {
	mock.Mock
}

// RetrieveURL provides a mock function with given fields: ctx, slug
func (_m *SlugRetriever) RetrieveURL(ctx context.Context, slug string) (string, error) {
	ret := _m.Called(ctx, slug)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, slug)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, slug)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, slug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSlugRetriever creates a new instance of SlugRetriever. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSlugRetriever(t interface {
	mock.TestingT
	Cleanup(func())
}) *SlugRetriever {
	mock := &SlugRetriever{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
